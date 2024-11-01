package initservice

import (
	"context"

	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/iggy/config"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/service"
)

func NewUserService(ctx context.Context, username string) *service.UserService {
	userRepo := repo.NewUserRepo(ctx, username)
	return service.NewUserService(ctx, userRepo, apictx.GetValue(ctx, "cfg").(config.Configuration))
}

func NewTokenService(ctx context.Context, username string) *service.TokenService {
	tokenRepo := repo.NewTokenRepo(ctx, username)
	return service.NewTokenService(ctx, tokenRepo, apictx.GetValue(ctx, "cfg").(config.Configuration))
}
