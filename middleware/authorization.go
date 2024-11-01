package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/kaibling/apiforge/apictx"
// 	"github.com/kaibling/apiforge/envelope"
// 	apierror "github.com/kaibling/apiforge/error"
// 	"github.com/kaibling/iggy/config"
// 	"github.com/kaibling/iggy/initservice"
// )

// func Authorization(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// read envelope
// 		e := envelope.ReadEnvelope(r)
// 		username := apictx.GetValue(r.Context(), "user_name").(string)
// 		if username != "admin" {

// 			// remove api_prefix for permission calculation
// 			base_url := strings.TrimPrefix(r.URL.String(), "/"+config.API_ENDPOINT_PREFIX)
// 			permission, err := ps.CollectionPermission(username, base_url)

// 			if err != nil {
// 				e.SetError(apierror.ServerError).Finish(w, r)
// 				return
// 			}

// 			decision, err := ps.Validate(r.Method, permission)
// 			if err != nil {
// 				e.SetError(apierror.ServerError).Finish(w, r)
// 				return
// 			}

// 			if !decision {
// 				e.SetError(apierror.Forbidden).Finish(w, r)
// 				return
// 			}
// 		}

// 		// // read token
// 		// if _, ok := r.Header["Authorization"]; !ok {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r)
// 		// 	return
// 		// }
// 		// if len(r.Header["Authorization"]) != 1 {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r)
// 		// 	return
// 		// }
// 		// authSlice := strings.Split(r.Header["Authorization"][0], " ")
// 		// if len(authSlice) != 2 {
// 		// 	e.SetError(apierror.Forbidden).Finish(w, r)
// 		// 	return
// 		// }
// 		// token := authSlice[1]
// 		// // validate token and get username
// 		// ts := initService.NewTokenService(r.Context(), config.SYSTEM_USER)
// 		// us := initService.NewUserService(r.Context(), config.SYSTEM_USER)
// 		// // todo set token last used
// 		// user, err := us.ValidateToken(token, ts)
// 		// if err != nil {
// 		// 	e.SetError(apierror.New(fmt.Errorf("unvalid token"), http.StatusBadRequest)).Finish(w, r)
// 		// 	return
// 		// }
// 		// ctx := context.WithValue(r.Context(), apictx.String("user_name"), user.Username)
// 		// ctx = context.WithValue(ctx, apictx.String("user_id"), user.ID)
// 		// ctx = context.WithValue(ctx, apictx.String("user_token"), token)
// 		next.ServeHTTP(w, r)
// 	})
// }
