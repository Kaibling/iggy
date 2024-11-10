package repo

import (
	"context"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunRepo struct {
	ctx       context.Context
	q         *sqlcrepo.Queries
	username  string
	requestID string
}

func NewRunRepo(ctx context.Context, username, requestID string) *RunRepo {
	return &RunRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		// username: ctx.Value(apictx.String("user_name")).(string),
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

	return r.FetchRun(newRunID)
}

func (r *RunRepo) FetchRun(id string) (*entity.Run, error) {
	rt, err := r.q.FetchRun(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Run{
		ID:         rt.ID,
		WorkflowID: rt.WorkflowID,
		Error:      &rt.Error.String,
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

func (r *RunRepo) FetchRunByWorkflow(workflowID string) ([]*entity.Run, error) {
	t, err := r.q.FetchRunByWorkflow(r.ctx, workflowID)
	if err != nil {
		return nil, err
	}

	runs := []*entity.Run{}
	for _, rt := range t {
		runs = append(runs, &entity.Run{
			ID:         rt.ID,
			WorkflowID: rt.WorkflowID,
			Error:      &rt.Error.String, // TODO fix pointer
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
