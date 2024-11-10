package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/pkg/config"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read envelope
		e := envelope.ReadEnvelope(r)

		// read token
		if _, ok := r.Header["Authorization"]; !ok {
			e.SetError(apierror.Forbidden).Finish(w, r)

			return
		}

		if len(r.Header["Authorization"]) != 1 {
			e.SetError(apierror.Forbidden).Finish(w, r)

			return
		}

		authSlice := strings.Split(r.Header["Authorization"][0], " ")

		position := 2
		if len(authSlice) != position {
			e.SetError(apierror.Forbidden).Finish(w, r)

			return
		}

		token := authSlice[1]

		// validate token and get username
		ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
		if err != nil {
			e.SetError(apierror.NewGeneric(err)).Finish(w, r)

			return
		}

		us, err := bootstrap.NewUserServiceAnonym(r.Context(), config.SystemUser)
		if err != nil {
			e.SetError(apierror.NewGeneric(err)).Finish(w, r)

			return
		}
		// todo set token last used
		// TODO use not found sql error
		user, err := us.ValidateToken(token, ts)
		if err != nil {
			e.SetError(apierror.New(apperror.NewStringGeneric("invalid token"), http.StatusUnauthorized)).Finish(w, r)

			return
		}

		ctx := context.WithValue(r.Context(), ctxkeys.UserNameKey, user.Username)
		ctx = context.WithValue(ctx, ctxkeys.UserIDKey, user.ID)
		ctx = context.WithValue(ctx, ctxkeys.TokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
