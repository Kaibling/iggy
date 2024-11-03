package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/crypto"
)

type tokenRepo interface {
	ReadToken(id string) (*entity.Token, error)
	ReadTokenByValue(t string) (*entity.Token, error)
	DeleteTokenByValue(t string) error
	ListTokens() ([]*entity.Token, error)
	ListUserToken(username string) ([]*entity.Token, error)
	CreateToken(t entity.NewToken) (*entity.Token, error)
}

type TokenService struct {
	ctx  context.Context
	repo tokenRepo
	cfg  config.Configuration
}

func NewTokenService(ctx context.Context, u tokenRepo, cfg config.Configuration) *TokenService {
	return &TokenService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *TokenService) ReadToken(id string) (*entity.Token, error) {
	return ts.repo.ReadToken(id)
}

func (ts *TokenService) ReadTokenByValue(v string) (*entity.Token, error) {
	return ts.repo.ReadTokenByValue(v)
}

func (ts *TokenService) CreateToken(u entity.NewToken) (*entity.Token, error) {
	tokenVal, err := crypto.GenerateAPIKey(ts.cfg.TokenKeyLength)
	if err != nil {
		return nil, err
	}
	u.Value = tokenVal
	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}
	return ts.repo.CreateToken(u)
}

func (ts *TokenService) ListTokens() ([]*entity.Token, error) {
	return ts.repo.ListTokens()
}
func (ts *TokenService) ListUserToken(username string) ([]*entity.Token, error) {
	return ts.repo.ListUserToken(username)
}
func (ts *TokenService) DeleteTokenByValue(v string) error {
	return ts.repo.DeleteTokenByValue(v)
}
