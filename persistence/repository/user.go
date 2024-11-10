package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type UserRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	db       *pgxpool.Pool
	Username string
}

func NewUserRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		db:       dbPool,
		Username: username,
	}
}

func (r *UserRepo) SaveUser(t entity.NewUser) (*entity.User, error) {
	ret, err := r.q.SaveUser(r.ctx, sqlcrepo.SaveUserParams{
		ID:       t.ID,
		Username: t.Username,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		CreatedBy: r.Username,
		ModifiedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ModifiedBy: r.Username,
		Active:     bool2Int(t.Active),
		Password:   pgtype.Text{String: t.Password, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return marshalUser(ret), nil
}

func (r *UserRepo) FetchUser(id string) (*entity.User, error) {
	rt, err := r.q.FetchUser(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return marshalUser(rt), nil
}

func (r *UserRepo) FetchByIDs(ids []string) ([]*entity.User, error) {
	rt, err := r.q.FetchByIDs(r.ctx, ids)
	if err != nil {
		return nil, err
	}
	return marshalUsers(rt), nil
}

func (r *UserRepo) FetchUserByName(name string) (*entity.User, error) {
	rt, err := r.q.FetchUserByName(r.ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return marshalUser(rt), nil
}

func (r *UserRepo) DeleteUser(id string) error {
	err := r.q.DeleteUser(r.ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}

func (r *UserRepo) UpdatePassword(passwordHash string, userID string) (*entity.User, error) {
	user, err := r.q.UpdatePassword(r.ctx, sqlcrepo.UpdatePasswordParams{
		ID: userID,
		Password: pgtype.Text{
			String: passwordHash,
			Valid:  true,
		},
		ModifiedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ModifiedBy: r.Username,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return marshalUser(user), nil
}

func (r *UserRepo) IDQuery(idQuery string) ([]string, error) {
	rows, err := r.db.Query(context.Background(), idQuery)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	var ids []string
	// Process the results
	for rows.Next() {
		var id string
		if err := rows.Scan(
			&id,
		); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("row iteration error: %v", err)
	}

	defer rows.Close()
	return ids, nil
}

func marshalUser(t sqlcrepo.User) *entity.User {
	u := &entity.User{
		ID:       t.ID,
		Username: t.Username,
		Active:   int2Bool(t.Active),
		Password: t.Password.String,
		Meta: entity.MetaData{
			CreatedAt:  t.CreatedAt.Time,
			CreatedBy:  t.CreatedBy,
			ModifiedAt: t.ModifiedAt.Time,
			ModifiedBy: t.ModifiedBy,
		},
	}
	return u
}

func marshalUsers(repoUsers []sqlcrepo.User) []*entity.User {
	users := []*entity.User{}
	for _, t := range repoUsers {
		users = append(users, &entity.User{
			ID:       t.ID,
			Username: t.Username,
			Active:   int2Bool(t.Active),
			Password: t.Password.String,
			Meta: entity.MetaData{
				CreatedAt:  t.CreatedAt.Time,
				CreatedBy:  t.CreatedBy,
				ModifiedAt: t.ModifiedAt.Time,
				ModifiedBy: t.ModifiedBy,
			},
		})
	}
	return users
}

func int2Bool(i int32) bool {
	return i == 1
}

func bool2Int(b bool) int32 {
	if b {
		return 1
	}
	return 0
}
