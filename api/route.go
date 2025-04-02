package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/api/auth"
	dynTab "github.com/kaibling/iggy/api/dyntab"
	"github.com/kaibling/iggy/api/export"
	"github.com/kaibling/iggy/api/run"
	"github.com/kaibling/iggy/api/token"
	"github.com/kaibling/iggy/api/uiconfig"
	"github.com/kaibling/iggy/api/user"
	"github.com/kaibling/iggy/api/workflow"
)

func Route() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Mount("/auth", auth.Route())
	r.Mount("/users", user.Route())
	r.Mount("/tokens", token.Route())
	r.Mount("/workflows", workflow.Route())
	r.Mount("/runs", run.Route())
	r.Mount("/dynamic-tables", dynTab.Route())
	r.Mount("/ui-config", uiconfig.Route())
	r.Mount("/backup", export.Route())

	return r
}
