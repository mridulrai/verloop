package main

import (
	"net/http"

	"github.com/mridulrai/verloop/pkg/config"
	"github.com/mridulrai/verloop/pkg/http/rest"
	"github.com/mridulrai/verloop/pkg/logger"
	"github.com/mridulrai/verloop/pkg/payment"
	"github.com/mridulrai/verloop/pkg/storage"
	//"github.com/mridulrai/verloop/pkg/stripe"
)

const appVersion string = "0.0.0"

func main() {
	logger := logger.InitSugarLogger()
	defer logger.Sync() // Flushes buffer, if any
	logger.Info("Starting My-Sales-Champion Payment Service")
	logger.Info("Initializing application for connecting to database and setting up API routes.")

	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	s, err := storage.NewStorage(config.GetDatabaseURI(), config.GetDatabaseName())
	if err != nil {
		panic(err)
	}

	st := stripe.NewStripeClient()

	ps := payment.NewPaymentService(s, st)
	app := rest.NewAppContext(ps)
	logger.Info("My-Sales-Champion Payment Service Ready", "version", appVersion)
	logger.Fatal(http.ListenAndServe(":80", app.GetRouter()))
}
