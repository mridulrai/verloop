package rest

import (
	"fmt"
	"net/http"

	"github.com/mridulrai/verloop/pkg/logger"
)

// indexHandler is a handler for the home page.
func (app *AppContext) indexHandler(w http.ResponseWriter, r *http.Request) {
	slog := logger.InitSugarLogger()
	defer slog.Sync()
	fmt.Fprintf(w, "My-Sales-Champion")
	slog.Debug("Index Handler")
}
