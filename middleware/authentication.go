package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/iggy/config"
	"github.com/kaibling/iggy/initservice"
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
		if len(authSlice) != 2 {
			e.SetError(apierror.Forbidden).Finish(w, r)
			return
		}
		token := authSlice[1]

		// validate token and get username
		ts := initservice.NewTokenService(r.Context(), config.SYSTEM_USER)
		us := initservice.NewUserService(r.Context(), config.SYSTEM_USER)
		// todo set token last used
		// TODO use not found sql error
		user, err := us.ValidateToken(token, ts)
		if err != nil {
			e.SetError(apierror.New(fmt.Errorf("invalid token"), http.StatusUnauthorized)).Finish(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), apictx.String("user_name"), user.Username)
		ctx = context.WithValue(ctx, apictx.String("user_id"), user.ID)
		ctx = context.WithValue(ctx, apictx.String("user_token"), token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
