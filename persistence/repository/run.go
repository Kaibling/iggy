package repo

import (
	"context"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/persistence/sqlcrepo"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	username string
}

func NewRunRepo(ctx context.Context, username string) *RunRepo {
	return &RunRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		// username: ctx.Value(apictx.String("user_name")).(string),
		username: username,
	}
}

func (r *RunRepo) SaveRun(newModel model.NewRun) (*model.Run, error) {
	newRunID, err := r.q.SaveRun(r.ctx, sqlcrepo.SaveRunParams{
		ID:         newModel.ID,
		WorkflowID: newModel.WorkflowID,
		Error: pgtype.Text{
			String: *newModel.Error,
			Valid:  true,
		},
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

func (r *RunRepo) FetchRun(id string) (*model.Run, error) {
	rt, err := r.q.FetchRun(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Run{
		ID:         rt.ID,
		WorkflowID: rt.WorkflowID,
		Error:      &rt.Error.String,
		StartTime:  rt.StartTime.Time,
		FinishTime: rt.FinishTime.Time,
		Meta: model.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *RunRepo) FetchRunByWorkflow(workflowID string) ([]*model.Run, error) {
	t, err := r.q.FetchRunByWorkflow(r.ctx, workflowID)
	if err != nil {
		return nil, err
	}
	runs := []*model.Run{}
	for _, rt := range t {
		runs = append(runs, &model.Run{
			ID:         rt.ID,
			WorkflowID: rt.WorkflowID,
			Error:      &rt.Error.String,
			StartTime:  rt.StartTime.Time,
			FinishTime: rt.FinishTime.Time,
			Meta: model.MetaData{
				CreatedAt:  rt.CreatedAt.Time,
				CreatedBy:  rt.CreatedBy,
				ModifiedAt: rt.ModifiedAt.Time,
				ModifiedBy: rt.ModifiedBy,
			},
		})
	}
	return runs, nil
}
