package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WorkflowRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	username string
}

func NewWorkflowRepo(ctx context.Context, username string) *WorkflowRepo {
	return &WorkflowRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		// username: ctx.Value(apictx.String("user_name")).(string),
		username: username,
	}
}

func (r *WorkflowRepo) SaveWorkflow(newModel entity.NewWorkflow) (*entity.Workflow, error) {
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
		//DeletedAt:   nil,
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
	return r.FetchWorkflow(newWorkflowID)
}

func (r *WorkflowRepo) FetchWorkflow(id string) (*entity.Workflow, error) {
	rt, err := r.q.FetchWorkflow(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return &entity.Workflow{
		ID:          rt.ID,
		Name:        rt.Name,
		Code:        rt.Code.String,
		FailOnError: rt.FailOnError,
		Meta: entity.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *WorkflowRepo) FetchAll() ([]*entity.Workflow, error) {
	t, err := r.q.FetchAllWorkflows(r.ctx)
	if err != nil {
		return nil, err
	}
	workflows := []*entity.Workflow{}
	for _, rt := range t {
		workflows = append(workflows, &entity.Workflow{
			ID:          rt.ID,
			Name:        rt.Name,
			Code:        rt.Code.String,
			FailOnError: rt.FailOnError,
			Meta: entity.MetaData{
				CreatedAt:  rt.CreatedAt.Time,
				CreatedBy:  rt.CreatedBy,
				ModifiedAt: rt.ModifiedAt.Time,
				ModifiedBy: rt.ModifiedBy,
			},
		})
	}
	return workflows, nil
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
		if err == pgx.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}
