package v1

// ErrorResponse represents error response data
type ErrorResponse struct {
	Error AppError `json:"errors"`
}

// AppError represents an application error
type AppError struct {
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
