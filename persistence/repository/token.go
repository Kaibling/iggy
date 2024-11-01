package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/lib/utils"
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

func (r *TokenRepo) CreateToken(t model.Token) (*model.Token, error) {

	ret, err := r.q.CreateToken(r.ctx, sqlcrepo.CreateTokenParams{
		Value:  t.Value,
		ID:     utils.NewULID().String(),
		UserID: t.UserID,
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
		Expires: pgtype.Timestamp{
			Time:  t.Expires,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return marshalToken(ret), nil
}

func (r *TokenRepo) ReadToken(id string) (*model.Token, error) {
	rt, err := r.q.GetToken(r.ctx, id)
	if err != nil {
		return nil, err
	}
	return marshalToken(rt), nil
}

func (r *TokenRepo) ReadTokenByValue(t string) (*model.Token, error) {
	rt, err := r.q.GetTokenByValue(r.ctx, t)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return marshalToken(rt), nil
}

func (r *TokenRepo) ListTokens() ([]*model.Token, error) {
	rt, err := r.q.ListTokens(r.ctx)
	if err != nil {
		return nil, err
	}
	return marshalTokens(rt), nil
}

func (r *TokenRepo) ListUserToken(username string) ([]*model.Token, error) {
	rt, err := r.q.ListUserTokensByName(r.ctx, username)
	if err != nil {
		return nil, err
	}
	return marshalTokens(rt), nil
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

func marshalToken(t sqlcrepo.Token) *model.Token {
	u := &model.Token{
		ID:        t.ID,
		Value:     t.Value,
		CreatedAt: t.CreatedAt.Time,
		CreatedBy: t.CreatedBy,
		UpdatedAt: t.UpdatedAt.Time,
		UpdatedBy: t.UpdatedBy,
		Active:    int2Bool(t.Active),
		Expires:   t.Expires.Time,
		UserID:    t.UserID,
	}
	return u
}

func marshalTokens(repoTokens []sqlcrepo.Token) []*model.Token {
	users := []*model.Token{}
	for _, t := range repoTokens {
		users = append(users, &model.Token{
			ID:        t.ID,
			Value:     t.Value,
			CreatedAt: t.CreatedAt.Time,
			CreatedBy: t.CreatedBy,
			UpdatedAt: t.UpdatedAt.Time,
			UpdatedBy: t.UpdatedBy,
			Active:    int2Bool(t.Active),
			Expires:   t.Expires.Time,
			UserID:    t.UserID,
		})
	}
	return users
}
