package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/persistence/sqlcrepo"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	Username string
}

func NewUserRepo(ctx context.Context, username string) *UserRepo {
	return &UserRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		// username: ctx.Value(apictx.String("user_name")).(string),
		Username: username,
	}
}

func (r *UserRepo) SaveUser(t model.NewUser) (*model.User, error) {
	ret, err := r.q.SaveUser(r.ctx, sqlcrepo.SaveUserParams{
		ID:       t.ID,
		Username: t.Username,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		CreatedBy: r.Username,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedBy: r.Username,
		Active:    bool2Int(t.Active),
		Password:  pgtype.Text{String: t.Password, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return marshalUser(ret), nil
}

func (r *UserRepo) FetchUser(id string) (*model.User, error) {
	rt, err := r.q.FetchUser(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return marshalUser(rt), nil
}

func (r *UserRepo) FetchAll() ([]*model.User, error) {
	rt, err := r.q.FetchAll(r.ctx)
	if err != nil {
		return nil, err
	}
	return marshalUsers(rt), nil
}

func (r *UserRepo) FetchUserByName(name string) (*model.User, error) {
	rt, err := r.q.FetchUserByName(r.ctx, name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return marshalUser(rt), nil
}

func (r *UserRepo) DeleteUser(id string) error {
	err := r.q.DeleteUser(r.ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}

func (r *UserRepo) UpdatePassword(passwordHash string, userID string) (*model.User, error) {
	user, err := r.q.UpdatePassword(r.ctx, sqlcrepo.UpdatePasswordParams{
		ID: userID,
		Password: pgtype.Text{
			String: passwordHash,
			Valid:  true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedBy: r.Username,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return marshalUser(user), nil
}

func marshalUser(t sqlcrepo.User) *model.User {
	u := &model.User{
		ID:       t.ID,
		Username: t.Username,
		Active:   int2Bool(t.Active),
		Password: t.Password.String,
		Meta: model.MetaData{
			CreatedAt: t.CreatedAt.Time,
			CreatedBy: t.CreatedBy,
			UpdatedAt: t.UpdatedAt.Time,
			UpdatedBy: t.UpdatedBy,
		},
	}
	return u
}

func marshalUsers(repoUsers []sqlcrepo.User) []*model.User {
	users := []*model.User{}
	for _, t := range repoUsers {
		users = append(users, &model.User{
			ID:       t.ID,
			Username: t.Username,
			Active:   int2Bool(t.Active),
			Meta: model.MetaData{
				CreatedAt: t.CreatedAt.Time,
				CreatedBy: t.CreatedBy,
				UpdatedAt: t.UpdatedAt.Time,
				UpdatedBy: t.UpdatedBy,
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
	} else {
		return 0
	}
}
