package rest

import (
	"net/http"

	"github.com/mridulrai/verloop/pkg/payment"
)

// AppContext holds core components for running this web application i.e. router and payment service.
type AppContext struct {
	router         http.Handler
	paymentService payment.Servicer
}

// NewAppContext is a factory func to create a new AppContext.
func NewAppContext(pr payment.Servicer) *AppContext {
	app := &AppContext{paymentService: pr}
	app.SetRoutes()
	return app
}

// GetRouter returns the app's router.
func (app *AppContext) GetRouter() http.Handler {
	return app.router
}
