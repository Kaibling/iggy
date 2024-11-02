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

type RunLogRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	username string
}

func NewRunLogRepo(ctx context.Context, username string) *RunLogRepo {
	return &RunLogRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		// username: ctx.Value(apictx.String("user_name")).(string),
		username: username,
	}
}

func (r *RunLogRepo) SaveRunLog(newModel model.NewRunLog) (*model.RunLog, error) {
	newRunLogID, err := r.q.SaveRunLog(r.ctx, sqlcrepo.SaveRunLogParams{
		ID:      newModel.ID,
		RunID:   newModel.RunID,
		Message: newModel.Message,
		Timestamp: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return r.FetchRunLog(newRunLogID)
}

func (r *RunLogRepo) FetchRunLogsByRun(runID string) ([]model.RunLog, error) {
	rt, err := r.q.FetchRunLogsByRun(r.ctx, runID)
	if err != nil {
		return nil, err
	}
	runLogs := []model.RunLog{}
	for _, run := range rt {
		runLogs = append(runLogs,
			model.RunLog{
				ID:        run.ID,
				RunID:     run.RunID,
				Message:   run.Message,
				Timestamp: run.Timestamp.Time,
			})
	}

	return runLogs, nil
}

func (r *RunLogRepo) FetchRunLog(id string) (*model.RunLog, error) {
	rt, err := r.q.FetchRunLog(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.RunLog{
		ID:        rt.ID,
		RunID:     rt.RunID,
		Message:   rt.Message,
		Timestamp: rt.Timestamp.Time,
	}, nil
}
