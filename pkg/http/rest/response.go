package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	v1 "github.com/mridulrai/verloop/pkg/http/rest/v1"
)

// SendHTTPSuccessResponse sends a HTTP response with a single error.
func SendHTTPSuccessResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// SendHTTPErrorResponse sends a HTTP response with a single error.
func SendHTTPErrorResponse(w http.ResponseWriter, status int, title string, description string) {
	e := v1.ErrorResponse{}
	e.Error.Status = strconv.Itoa(status)
	e.Error.Title = title
	e.Error.Description = description

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(e)
}
