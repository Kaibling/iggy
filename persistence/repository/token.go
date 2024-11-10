package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/persistence/sqlcrepo"
)

type TokenRepo struct {
	ctx      context.Context
	q        *sqlcrepo.Queries
	Username string
}

func NewTokenRepo(ctx context.Context, username string, dbPool *pgxpool.Pool) *TokenRepo {
	return &TokenRepo{
		ctx:      ctx,
		q:        sqlcrepo.New(dbPool),
		Username: username,
	}
}

func (r *TokenRepo) CreateToken(t entity.NewToken) (*entity.Token, error) {
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

func (r *TokenRepo) ReadToken(id string) (*entity.Token, error) {
	rt, err := r.q.GetToken(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.Token{
		ID:      rt.ID,
		Value:   rt.Value,
		Active:  int2Bool(rt.Active),
		Expires: rt.Expires.Time,
		User:    entity.Identifier{ID: rt.UserID, Name: rt.Username},
		Meta: entity.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *TokenRepo) ReadTokenByValue(t string) (*entity.Token, error) {
	rt, err := r.q.GetTokenByValue(r.ctx, t)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sql.ErrNoRows
		}

		return nil, err
	}

	return &entity.Token{
		ID:      rt.ID,
		Value:   rt.Value,
		Active:  int2Bool(rt.Active),
		Expires: rt.Expires.Time,
		User:    entity.Identifier{ID: rt.UserID, Name: rt.Username},
		Meta: entity.MetaData{
			CreatedAt:  rt.CreatedAt.Time,
			CreatedBy:  rt.CreatedBy,
			ModifiedAt: rt.ModifiedAt.Time,
			ModifiedBy: rt.ModifiedBy,
		},
	}, nil
}

func (r *TokenRepo) ListTokens() ([]*entity.Token, error) {
	rt, err := r.q.ListTokens(r.ctx)
	if err != nil {
		return nil, err
	}

	users := []*entity.Token{}

	for _, t := range rt {
		users = append(users, &entity.Token{
			ID:      t.ID,
			Value:   t.Value,
			Active:  int2Bool(t.Active),
			Expires: t.Expires.Time,
			User:    entity.Identifier{ID: t.UserID, Name: t.Username},
			Meta: entity.MetaData{
				CreatedAt:  t.CreatedAt.Time,
				CreatedBy:  t.CreatedBy,
				ModifiedAt: t.ModifiedAt.Time,
				ModifiedBy: t.ModifiedBy,
			},
		})
	}

	return users, nil
}

func (r *TokenRepo) ListUserToken(username string) ([]*entity.Token, error) {
	rt, err := r.q.ListUserTokensByName(r.ctx, username)
	if err != nil {
		return nil, err
	}

	users := []*entity.Token{}

	for _, t := range rt {
		users = append(users, &entity.Token{
			ID:      t.ID,
			Value:   t.Value,
			Active:  int2Bool(t.Active),
			Expires: t.Expires.Time,
			User:    entity.Identifier{ID: t.UserID, Name: t.Username},
			Meta: entity.MetaData{
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
		if errors.Is(err, pgx.ErrNoRows) {
			return sql.ErrNoRows
		}

		return err
	}

	return nil
}
