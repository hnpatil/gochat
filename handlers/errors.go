package handlers

// ErrorResponse represents the standard error response format
type ErrorResponse struct {
	Error ErrorDetail `json:"error"` // Error details
}

// ErrorDetail contains the error message
type ErrorDetail struct {
	Message string `json:"message" example:"Invalid request"` // Error message description
}
