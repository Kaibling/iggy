package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type WorkflowRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	username string
}

func NewWorkflowRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *WorkflowRepo {
	return &WorkflowRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		username: username,
	}
}

func (r *WorkflowRepo) CreateWorkflows(newModels []*entity.NewWorkflow) ([]entity.Workflow, error) {
	maxDepth := 1
	newWorkflowsIDs := []string{}

	for _, newModel := range newModels {
		newWorkflowID, err := r.q.SaveWorkflow(r.ctx, sqlcrepo.SaveWorkflowParams{
			ID:   newModel.ID,
			Name: newModel.Name,
			Code: pgtype.Text{
				String: newModel.Code,
				Valid:  true,
			},
			ObjectType:  string(newModel.ObjectType),
			BuildIn:     newModel.BuildIn,
			FailOnError: newModel.FailOnError,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			CreatedBy: r.username,
			ModifiedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			ModifiedBy: r.username,
		})
		if err != nil {
			return nil, err
		}

		newWorkflowsIDs = append(newWorkflowsIDs, newWorkflowID)
	}

	return r.FetchWorkflows(newWorkflowsIDs, maxDepth)
}

func (r *WorkflowRepo) FetchToBackup() ([]entity.Workflow, error) {
	maxDepth := 1

	backupIDs, err := r.q.FetchToBackup(r.ctx)
	if err != nil {
		return nil, err
	}

	return r.FetchWorkflows(backupIDs, maxDepth)
}

func (r *WorkflowRepo) FetchBackupAll() ([]entity.Workflow, error) {
	maxDepth := 1

	backupIDs, err := r.q.FetchBackupAll(r.ctx)
	if err != nil {
		return nil, err
	}

	return r.FetchWorkflows(backupIDs, maxDepth)
}

func (r *WorkflowRepo) Upserts(workflows []entity.Workflow) error {
	for _, wf := range workflows {
		if err := r.q.UpsertWorkflow(r.ctx, sqlcrepo.UpsertWorkflowParams{
			ID:   wf.ID,
			Name: wf.Name,
			Code: pgtype.Text{
				String: wf.Code,
				Valid:  true,
			},
			ObjectType:  string(wf.ObjectType),
			BuildIn:     wf.BuildIn,
			FailOnError: wf.FailOnError,
			ModifiedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			ModifiedBy: r.username,
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			CreatedBy: r.username,
		}); err != nil {
			if !strings.Contains(err.Error(), "duplicate key value") {
				return err
			}
			fmt.Printf("%s : %s\n", wf.Name, err.Error())

		}

	}

	return nil
}

func (r *WorkflowRepo) UpdateWorkflow(workflowID string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error) {
	maxDepth := 1

	if err := r.q.UpdateWorkflow(r.ctx, ConvertToUpdateWorkflowParams(updateEntity, workflowID, r.username)); err != nil {
		return nil, err
	}

	if updateEntity.Children != nil {
		// delete old children
		if err := r.q.DeleteWorkflowChildren(r.ctx, workflowID); err != nil {
			return nil, err
		}

		for _, child := range updateEntity.Children {
			// create new children
			if err := r.q.SaveWorkflowChildren(r.ctx, sqlcrepo.SaveWorkflowChildrenParams{
				WorkflowID: workflowID,
				ChildrenID: child.ID,
				StepOrder:  int32(child.StepOrder),
			}); err != nil {
				return nil, err
			}
		}
	}

	if wfs, err := r.FetchWorkflows([]string{workflowID}, maxDepth); err != nil {
		return nil, err
	} else { //nolint: revive
		return &wfs[0], nil
	}
}

func (r *WorkflowRepo) FetchWorkflows(ids []string, depth int) ([]entity.Workflow, error) { //nolint: funlen
	rawQuery := `
	WITH RECURSIVE workflow_hierarchy AS (
		SELECT 
			w.id,
			w.name,
			w.code,
			w.object_type,
			w.fail_on_error,
			w.build_in,
			w.created_by,
			w.modified_by,
			w.created_at,
			w.modified_at,
			NULL::text AS parent,
			0 AS depth
		FROM 
			workflows w
		WHERE 
			w.deleted_at IS NULL AND w.id IN(%s)
		
		UNION ALL
		
		SELECT 
			w.id,
			w.name,
			w.code,
			w.object_type,
			w.fail_on_error,
			w.build_in,
			w.created_by,
			w.modified_by,
			w.created_at,
			w.modified_at,
			wc.workflow_id AS parent,
			wh.depth + 1 AS depth
		FROM 
			workflows w
		JOIN 
			workflows_children wc ON w.id = wc.children_id
		JOIN 
			workflow_hierarchy wh ON wc.workflow_id = wh.id
		WHERE 
			wh.depth < $1
	)
	SELECT * FROM workflow_hierarchy;
`
	placeholders := "'" + strings.Join(ids, "','") + "'"
	query := fmt.Sprintf(rawQuery, placeholders)

	rows, err := r.db.Query(context.Background(), query, depth)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var workflows []RawWorkflow

	// Process the results
	for rows.Next() {
		var wf RawWorkflow
		if err := rows.Scan(
			&wf.ID,
			&wf.Name,
			&wf.Code,
			&wf.ObjectType,
			&wf.FailOnError,
			&wf.BuildIn,
			&wf.Meta.CreatedBy,
			&wf.Meta.ModifiedBy,
			&wf.Meta.CreatedAt,
			&wf.Meta.ModifiedAt,
			&wf.Parent,
			&wf.Depth,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		workflows = append(workflows, wf)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	workflowMap := map[string][]entity.Workflow{}
	root := doa(workflows, depth, workflowMap)

	v, ok := root["root"]
	if !ok {
		return []entity.Workflow{}, nil
	}

	return v, nil
}

func (r *WorkflowRepo) IDQuery(idQuery string) ([]string, error) {
	rows, err := r.db.Query(context.Background(), idQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	var ids []string
	// Process the results
	for rows.Next() {
		var id string
		if err := rows.Scan(
			&id,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	defer rows.Close()

	return ids, nil
}

func (r *WorkflowRepo) DeleteWorkflow(id string) error {
	err := r.q.DeleteWorkflow(r.ctx, sqlcrepo.DeleteWorkflowParams{
		DeletedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ID: id,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sql.ErrNoRows
		}

		return err
	}

	return nil
}

type RawWorkflow struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Code        string              `json:"code"`
	BuildIn     bool                `json:"build_in"`
	FailOnError bool                `json:"fail_on_error"`
	ObjectType  entity.WorkflowType `json:"object_type"`
	Meta        entity.MetaData     `json:"meta"`
	Parent      *string
	Depth       int
}

func doa(r []RawWorkflow, depth int, b map[string][]entity.Workflow) map[string][]entity.Workflow {
	for _, rr := range r {
		if rr.Depth == depth { //nolint: nestif
			newModel := entity.Workflow{ //nolint: exhaustruct
				ID:          rr.ID,
				Name:        rr.Name,
				Code:        rr.Code,
				BuildIn:     rr.BuildIn,
				FailOnError: rr.FailOnError,
				ObjectType:  rr.ObjectType,
				Meta:        rr.Meta,
			} // Create a new Workflow from RawResult

			// Determine parent key; default to "root" if nil
			parentKey := "root"
			if rr.Parent != nil {
				parentKey = *rr.Parent
			}

			// If children exist for this workflow, merge them
			if children, exists := b[newModel.ID]; exists {
				if newModel.Children == nil {
					newModel.Children = []entity.Workflow{}
				}

				newModel.Children = append(newModel.Children, children...)
			}

			// Save the workflow into the map under the parent key
			if parentVec, exists := b[parentKey]; exists {
				b[parentKey] = append(parentVec, newModel)
			} else {
				b[parentKey] = []entity.Workflow{newModel}
			}
		}
	}

	if depth > 0 {
		// Recursive call to process deeper levels
		lower := doa(r, depth-1, b)
		for k, v := range lower {
			b[k] = v // Extend the current map with the results from lower depth
		}
	}

	return b
}

func ConvertToUpdateWorkflowParams(
	input entity.UpdateWorkflow,
	workflowID, modifiedBy string,
) sqlcrepo.UpdateWorkflowParams {
	var params sqlcrepo.UpdateWorkflowParams

	params.ID = workflowID
	params.ModifiedBy = modifiedBy

	// Set ModifiedAt to current timestamp
	params.ModifiedAt.Time = time.Now()
	params.ModifiedAt.Valid = true

	// Set each field conditionally based on whether it was provided
	if input.Name != nil {
		params.Name.String = *input.Name
		params.Name.Valid = true
	} else {
		params.Name.Valid = false
	}

	if input.Code != nil {
		params.Code.String = *input.Code
		params.Code.Valid = true
	} else {
		params.Code.Valid = false
	}

	if input.ObjectType != nil {
		params.ObjectType.String = string(*input.ObjectType)
		params.ObjectType.Valid = true
	} else {
		params.ObjectType.Valid = false
	}

	if input.FailOnError != nil {
		params.FailOnError.Bool = *input.FailOnError
		params.FailOnError.Valid = true
	} else {
		params.FailOnError.Valid = false
	}

	return params
}
