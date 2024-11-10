package apperror

import "net/http"

type AppError struct {
	msg        string
	StatusCode int
}

func New(err error, status int) AppError {
	return AppError{msg: err.Error(), StatusCode: status}
}

func NewGeneric(err error) AppError {
	return AppError{msg: err.Error(), StatusCode: http.StatusInternalServerError}
}

func (e AppError) Error() string {
	return e.msg
}

func (e AppError) HTTPStatus() int {
	return e.StatusCode
}

var ErrForbidden = AppError{
	msg:        "permission denied",
	StatusCode: http.StatusForbidden,
}

var ErrMissingJSAdapter = AppError{
	msg:        "Adapter missing",
	StatusCode: http.StatusInternalServerError,
}

var ErrMalformedRequest = AppError{
	msg:        "client request malformed",
	StatusCode: http.StatusUnprocessableEntity,
}

var ErrMissingContext = AppError{
	msg:        "context data missing",
	StatusCode: http.StatusInternalServerError,
}

// var ServerError = AppError{
// 	msg:        "internal server error",
// 	StatusCode: http.StatusInternalServerError,
// }

// var NotFound = AppError{
// 	msg:        "path not found",
// 	StatusCode: http.StatusNotFound,
// }

// var DataNotFound = AppError{
// 	msg:        "not found",
// 	StatusCode: http.StatusNotFound,
// }
