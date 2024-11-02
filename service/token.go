package service

import (
	"context"

	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/crypto"
)

type tokenRepo interface {
	ReadToken(id string) (*model.Token, error)
	ReadTokenByValue(t string) (*model.Token, error)
	DeleteTokenByValue(t string) error
	ListTokens() ([]*model.Token, error)
	ListUserToken(username string) ([]*model.Token, error)
	CreateToken(t model.NewToken) (*model.Token, error)
}

type TokenService struct {
	ctx  context.Context
	repo tokenRepo
	cfg  config.Configuration
}

func NewTokenService(ctx context.Context, u tokenRepo, cfg config.Configuration) *TokenService {
	return &TokenService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *TokenService) ReadToken(id string) (*model.Token, error) {
	return ts.repo.ReadToken(id)
}

func (ts *TokenService) ReadTokenByValue(v string) (*model.Token, error) {
	return ts.repo.ReadTokenByValue(v)
}

func (ts *TokenService) CreateToken(u model.NewToken) (*model.Token, error) {
	tokenVal, err := crypto.GenerateAPIKey(ts.cfg.TokenKeyLength)
	if err != nil {
		return nil, err
	}
	u.Value = tokenVal
	return ts.repo.CreateToken(u)
}

func (ts *TokenService) ListTokens() ([]*model.Token, error) {
	return ts.repo.ListTokens()
}
func (ts *TokenService) ListUserToken(username string) ([]*model.Token, error) {
	return ts.repo.ListUserToken(username)
}
func (ts *TokenService) DeleteTokenByValue(v string) error {
	return ts.repo.DeleteTokenByValue(v)
}
