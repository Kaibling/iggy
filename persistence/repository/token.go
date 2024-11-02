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

type TokenRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	Username string
}

func NewTokenRepo(ctx context.Context, username string) *TokenRepo {
	return &TokenRepo{
		ctx: ctx,
		q:   sqlcrepo.New(ctx.Value(ctxkeys.DBConnKey).(*pgxpool.Pool)),
		//username: ctx.Value(apictx.String("user_name")).(string),
		Username: username,
	}
}

func (r *TokenRepo) CreateToken(t model.NewToken) (*model.Token, error) {

	newTokenID, err := r.q.CreateToken(r.ctx, sqlcrepo.CreateTokenParams{
		Value:  t.Value,
		ID:     t.ID,
		UserID: t.UserID,
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
		Expires: pgtype.Timestamp{
			Time:  t.Expires,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return r.ReadToken(newTokenID)
}

func (r *TokenRepo) ReadToken(id string) (*model.Token, error) {
	rt, err := r.q.GetToken(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Token{
		ID:      rt.ID,
		Value:   rt.Value,
		Active:  int2Bool(rt.Active),
		Expires: rt.Expires.Time,
		User:    model.Identifier{ID: rt.UserID, Name: rt.Username},
		Meta: model.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *TokenRepo) ReadTokenByValue(t string) (*model.Token, error) {
	rt, err := r.q.GetTokenByValue(r.ctx, t)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &model.Token{
		ID:      rt.ID,
		Value:   rt.Value,
		Active:  int2Bool(rt.Active),
		Expires: rt.Expires.Time,
		User:    model.Identifier{ID: rt.UserID, Name: rt.Username},
		Meta: model.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *TokenRepo) ListTokens() ([]*model.Token, error) {
	rt, err := r.q.ListTokens(r.ctx)
	if err != nil {
		return nil, err
	}
	users := []*model.Token{}
	for _, t := range rt {
		users = append(users, &model.Token{
			ID:      t.ID,
			Value:   t.Value,
			Active:  int2Bool(t.Active),
			Expires: t.Expires.Time,
			User:    model.Identifier{ID: t.UserID, Name: t.Username},
			Meta: model.MetaData{
				CreatedAt:  t.CreatedAt.Time,
				CreatedBy:  t.CreatedBy,
				ModifiedAt: t.ModifiedAt.Time,
				ModifiedBy: t.ModifiedBy,
			},
		})
	}
	return users, nil
}

func (r *TokenRepo) ListUserToken(username string) ([]*model.Token, error) {
	rt, err := r.q.ListUserTokensByName(r.ctx, username)
	if err != nil {
		return nil, err
	}
	users := []*model.Token{}
	for _, t := range rt {
		users = append(users, &model.Token{
			ID:      t.ID,
			Value:   t.Value,
			Active:  int2Bool(t.Active),
			Expires: t.Expires.Time,
			User:    model.Identifier{ID: t.UserID, Name: t.Username},
			Meta: model.MetaData{
				CreatedAt:  t.CreatedAt.Time,
				CreatedBy:  t.CreatedBy,
				ModifiedAt: t.ModifiedAt.Time,
				ModifiedBy: t.ModifiedBy,
			},
		})
	}
	return users, nil
}

func (r *TokenRepo) DeleteTokenByValue(t string) error {
	err := r.q.DeleteTokenByValue(r.ctx, t)
	if err != nil {
		if err == pgx.ErrNoRows {
			return sql.ErrNoRows
		}
		return err
	}
	return nil
}
