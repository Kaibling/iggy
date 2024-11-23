package middleware

import (
	"context"
	"net/http"
	"strings"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read envelope
		e, l, multiErr := envelope.GetEnvelopeAndLogger(r)
		if multiErr != nil {
			e.SetError(apierror.NewGeneric(multiErr)).Finish(w, r, l)

			return
		}

		l.Debug("trying authentication")

		token, err := extractToken(r.Header)
		if err != nil {
			e.SetError(apierror.ErrForbidden).Finish(w, r, l)

			return
		}

		user, err := validateTokenAndUser(r.Context(), token)
		if err != nil {
			e.SetError(apierror.ErrForbidden).Finish(w, r, l)

			return
		}

		l.AddStringField("username", user.Username)
		l.Debug("client authenticated")

		ctx := context.WithValue(r.Context(), ctxkeys.UserNameKey, user.Username)
		ctx = context.WithValue(ctx, ctxkeys.UserIDKey, user.ID)
		ctx = context.WithValue(ctx, ctxkeys.TokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractToken(header http.Header) (string, apierror.HTTPError) {
	if _, ok := header["Authorization"]; !ok {
		return "", apierror.ErrForbidden
	}

	if len(header["Authorization"]) != 1 {
		return "", apierror.ErrForbidden
	}

	authSlice := strings.Split(header["Authorization"][0], " ")

	position := 2
	if len(authSlice) != position {
		return "", apierror.ErrForbidden
	}

	return authSlice[1], nil
}

func validateTokenAndUser(ctx context.Context, token string) (*entity.User, apierror.HTTPError) {
	errs := apierror.NewMultiError()
	// validate token and get username

	ts, err := bootstrap.NewTokenServiceAnonym(ctx, config.SystemUser)
	errs.Add(err)

	us, err := bootstrap.NewUserServiceAnonym(ctx, config.SystemUser)
	errs.Add(err)

	if errs.HasError() {
		return nil, apierror.NewMulti(apierror.ErrContextMissing, errs.GetStrErrors())
	}

	// todo set token last used
	// TODO use not found sql error
	user, err := us.ValidateToken(token, ts)
	if err != nil {
		return nil, apierror.New(apperror.NewStringGeneric("invalid token"), http.StatusUnauthorized)
	}

	return user, nil
}
