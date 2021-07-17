package rest

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// SetRoutes sets up routes
func (app *AppContext) SetRoutes() {
	indexHandlers := alice.New(context.ClearHandler)
	// noAuthHandlers := alice.New(context.ClearHandler, identifierHandler, loggingHandler, recoverHandler)
	// AuthHandlers := alice.New(context.ClearHandler, identifierHandler, loggingHandler, recoverHandler, app.authHandler)

	r := mux.NewRouter()
	r.Handle("/", indexHandlers.ThenFunc(app.indexHandler))

	app.router = r
}
