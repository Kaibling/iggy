package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/kaibling/apiforge/apictx"
// 	"github.com/kaibling/apiforge/envelope"
// 	apierror "github.com/kaibling/apiforge/apierror"
// 	"github.com/kaibling/iggy/config"
// 	"github.com/kaibling/iggy/bootstrap"
// )

// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// read envelope
// 		e, l, err := envelope.GetEnvelopeAndLogger(r)
// if err != nil {
// 	e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

// 	return
// }
// 		username := apictx.GetValue(r.Context(), "user_name").(string)
// 		if username != "admin" {

// 			// remove api_prefix for permission calculation
// 			base_url := strings.TrimPrefix(r.URL.String(), "/"+config.API_ENDPOINT_PREFIX)
// 			permission, err := ps.CollectionPermission(username, base_url)

// 			if err != nil {
// 				e.SetError(apierror.ServerError).Finish(w, r,l)
// 				return
// 			}

// 			decision, err := ps.Validate(r.Method, permission)
// 			if err != nil {
// 				e.SetError(apierror.ServerError).Finish(w, r,l)
// 				return
// 			}

// 			if !decision {
// 				e.SetError(apierror.Forbidden).Finish(w, r,l)
// 				return
// 			}
// 		}

// 		// // read token
// 		// if _, ok := r.Header["Authorization"]; !ok {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r,l)
// 		// 	return
// 		// }
// 		// if len(r.Header["Authorization"]) != 1 {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r,l)
// 		// 	return
// 		// }
// 		// authSlice := strings.Split(r.Header["Authorization"][0], " ")
// 		// if len(authSlice) != 2 {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r,l)
// 		// 	return
// 		// }
// 		// token := authSlice[1]
// 		// // validate token and get username
// 		// ts := initService.NewTokenService(r.Context(), config.SYSTEM_USER)
// 		// us := initService.NewUserService(r.Context(), config.SYSTEM_USER)
// 		// // todo set token last used
// 		// user, err := us.ValidateToken(token, ts)
// 		// if err != nil {
// 		// 	e.SetError(apierror.New(fmt.Errorf("unvalid token"), http.StatusBadRequest)).Finish(w, r,l)
// 		// 	return
// 		// }
// 		// ctx := context.WithValue(r.Context(), apictx.String("user_name"), user.Username)
// 		// ctx = context.WithValue(ctx, apictx.String("user_id"), user.ID)
// 		// ctx = context.WithValue(ctx, apictx.String("user_token"), token)
// 		next.ServeHTTP(w, r)
// 	})
// }
