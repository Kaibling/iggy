package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type RunRepo struct {
	ctx       context.Context
	q         *sqlcrepo.Queries
	db        *pgxpool.Pool
	username  string
	requestID string
}

func NewRunRepo(ctx context.Context, username, requestID string, dbPool *pgxpool.Pool) *RunRepo {
	return &RunRepo{
		ctx:       ctx,
		q:         sqlcrepo.New(dbPool),
		db:        dbPool,
		username:  username,
		requestID: requestID,
	}
}

func (r *RunRepo) SaveRun(newModel entity.NewRun) (*entity.Run, error) {
	var pgError pgtype.Text

	if newModel.Error != nil {
		pgError.String = *newModel.Error
		pgError.Valid = true
	} else {
		pgError.Valid = false
	}

	newRunID, err := r.q.SaveRun(r.ctx, sqlcrepo.SaveRunParams{
		RequestID: pgtype.Text{
			String: r.requestID,
			Valid:  true,
		},
		ID:         newModel.ID,
		WorkflowID: newModel.WorkflowID,
		UserID:     newModel.UserID,
		Error:      pgError,
		StartTime: pgtype.Timestamp{
			Time:  newModel.StartTime,
			Valid: true,
		},
		FinishTime: pgtype.Timestamp{
			Time:  newModel.FinishTime,
			Valid: true,
		},
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

	rr, err := r.FetchRuns([]string{newRunID})
	if err != nil {
		return nil, err
	}
	// todo if len is not 0
	return rr[0], nil
}

func (r *RunRepo) FetchRuns(ids []string) ([]*entity.Run, error) {
	ru, err := r.q.FetchRuns(r.ctx, ids)
	if err != nil {
		return nil, err
	}

	runs := []*entity.Run{}

	for _, rt := range ru {
		runs = append(runs, &entity.Run{
			ID:         rt.ID,
			Workflow:   entity.Identifier{ID: rt.WorkflowID, Name: rt.WorkflowName},
			User:       entity.Identifier{ID: rt.UserID, Name: rt.UserName},
			Error:      &rt.Error.String,
			StartTime:  rt.StartTime.Time,
			FinishTime: rt.FinishTime.Time,
			Meta: entity.MetaData{
				CreatedAt:  rt.CreatedAt.Time,
				CreatedBy:  rt.CreatedBy,
				ModifiedAt: rt.ModifiedAt.Time,
				ModifiedBy: rt.ModifiedBy,
			},
		})
	}

	return runs, nil
}

func (r *RunRepo) FetchRunByWorkflow(workflowID string) ([]*entity.Run, error) {
	t, err := r.q.FetchRunByWorkflow(r.ctx, workflowID)
	if err != nil {
		return nil, err
	}

	runs := []*entity.Run{}

	for _, rt := range t {
		pErr := rt.Error
		runs = append(runs, &entity.Run{
			ID:         rt.ID,
			Workflow:   entity.Identifier{ID: rt.WorkflowID, Name: rt.WorkflowName},
			User:       entity.Identifier{ID: rt.UserID, Name: rt.UserName},
			Error:      &pErr.String,
			StartTime:  rt.StartTime.Time,
			FinishTime: rt.FinishTime.Time,
			Meta: entity.MetaData{
				CreatedAt:  rt.CreatedAt.Time,
				CreatedBy:  rt.CreatedBy,
				ModifiedAt: rt.ModifiedAt.Time,
				ModifiedBy: rt.ModifiedBy,
			},
		})
	}

	return runs, nil
}

func (r *RunRepo) FetchRunByRequestID(requestID string) (*entity.Run, error) {
	rt, err := r.q.FetchRunByRequestID(r.ctx, pgtype.Text{String: requestID, Valid: true})
	if err != nil {
		return nil, err
	}

	pErr := rt.Error

	return &entity.Run{
		ID:         rt.ID,
		Workflow:   entity.Identifier{ID: rt.WorkflowID, Name: rt.WorkflowName},
		User:       entity.Identifier{ID: rt.UserID, Name: rt.UserName},
		Error:      &pErr.String,
		StartTime:  rt.StartTime.Time,
		FinishTime: rt.FinishTime.Time,
		Meta: entity.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *RunRepo) IDQuery(idQuery string) ([]string, error) {
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
