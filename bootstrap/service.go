package bootstrap

import (
	"context"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/config"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/service"
)

func NewUserService(ctx context.Context) *service.UserService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)

	userRepo := repo.NewUserRepo(ctx, username)
	return service.NewUserService(ctx, userRepo, cfg)
}

func NewTokenService(ctx context.Context) *service.TokenService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	tokenRepo := repo.NewTokenRepo(ctx, username)
	return service.NewTokenService(ctx, tokenRepo, cfg)
}

func NewUserServiceAnonym(ctx context.Context, username string) *service.UserService {
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	userRepo := repo.NewUserRepo(ctx, username)
	return service.NewUserService(ctx, userRepo, cfg)
}

func NewTokenServiceAnonym(ctx context.Context, username string) *service.TokenService {
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	tokenRepo := repo.NewTokenRepo(ctx, username)
	return service.NewTokenService(ctx, tokenRepo, cfg)
}
