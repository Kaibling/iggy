package bootstrap

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/pkg/config"
)

func ContextParams(ctx context.Context) (params.Pagination, error) {
	p, ok := ctxkeys.GetValue(ctx, ctxkeys.PaginationKey).(params.Pagination)
	if !ok {
		return params.Pagination{}, apperror.ErrNewMissingContext("pagination")
	}

	return p, nil
}

func ContextToken(ctx context.Context) (string, error) {
	t, ok := ctxkeys.GetValue(ctx, ctxkeys.TokenKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("token")
	}

	return t, nil
}

func ContextLogger(ctx context.Context) (logging.Writer, error) { //nolint: ireturn
	l, ok := ctxkeys.GetValue(ctx, ctxkeys.LoggerKey).(logging.Writer)
	if !ok {
		return nil, apperror.ErrNewMissingContext("logger")
	}

	return l, nil
}

func ContextRequestID(ctx context.Context) (string, error) {
	requestID, ok := ctxkeys.GetValue(ctx, ctxkeys.RequestIDKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("requestID")
	}

	return requestID, nil
}

func ContextUserID(ctx context.Context) (string, error) {
	userID, ok := ctxkeys.GetValue(ctx, ctxkeys.UserIDKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("userID")
	}

	return userID, nil
}

func ContextDefaultData(ctx context.Context) (config.Configuration, *pgxpool.Pool, logging.Writer, string, error) {
	cfg, dbPool, log, err := contextBasisData(ctx)
	if err != nil {
		return cfg, dbPool, log, "", err
	}

	username, ok := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	if !ok {
		return cfg, dbPool, log, "", apperror.ErrNewMissingContext("username")
	}

	return cfg, dbPool, log, username, nil
}

func contextBasisData(ctx context.Context) (config.Configuration, *pgxpool.Pool, logging.Writer, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return config.Configuration{}, nil, nil, apperror.ErrNewMissingContext("config")
	}

	dbPool, ok := ctxkeys.GetValue(ctx, ctxkeys.DBConnKey).(*pgxpool.Pool)
	if !ok {
		return config.Configuration{}, nil, nil, apperror.ErrNewMissingContext("db pool")
	}

	log, ok := ctxkeys.GetValue(ctx, ctxkeys.LoggerKey).(logging.Writer)
	if !ok {
		return config.Configuration{}, nil, nil, apperror.ErrNewMissingContext("logger")
	}

	return cfg, dbPool, log, nil
}
