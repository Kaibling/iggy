package repo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type RunLogRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	username string
}

func NewRunLogRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *RunLogRepo {
	return &RunLogRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		username: username,
	}
}

func (r *RunLogRepo) CreateRunLog(newModel entity.NewRunLog) (*entity.RunLog, error) {
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

func (r *RunLogRepo) FetchRunLogsByRun(runID string) ([]entity.RunLog, error) {
	rt, err := r.q.FetchRunLogsByRun(r.ctx, runID)
	if err != nil {
		return nil, err
	}

	runLogs := []entity.RunLog{}

	for _, run := range rt {
		runLogs = append(runLogs,
			entity.RunLog{
				ID:        run.ID,
				RunID:     run.RunID,
				Message:   run.Message,
				Timestamp: run.Timestamp.Time,
			})
	}

	return runLogs, nil
}

func (r *RunLogRepo) FetchRunLog(id string) (*entity.RunLog, error) {
	rt, err := r.q.FetchRunLog(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.RunLog{
		ID:        rt.ID,
		RunID:     rt.RunID,
		Message:   rt.Message,
		Timestamp: rt.Timestamp.Time,
	}, nil
}
