package httperrors

type ErrorResponse struct {
	Message string `json:"message"`
} // @name ErrorResponse

// NewErrorResponse creates a new Error Response
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}
