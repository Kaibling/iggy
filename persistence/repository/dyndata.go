package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type DynDataRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	username string
}

func NewDynDataRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *DynDataRepo {
	return &DynDataRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		username: username,
	}
}

func (r *DynDataRepo) Query(idQuery string) error {
	rows, err := r.db.Query(context.Background(), idQuery)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	return nil
}
