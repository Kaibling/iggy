package apperror

import "net/http"

type AppError struct {
	msg        string
	StatusCode int
	errors     []string
}

// type AppMultiError struct {
// 	appError AppError
// 	errors   []error
// }

func New(err error, status int) AppError {
	return AppError{msg: err.Error(), StatusCode: status}
}

func NewGeneric(err error) AppError {
	return AppError{msg: err.Error(), StatusCode: http.StatusInternalServerError}
}

func NewStringGeneric(errMsg string) AppError {
	return AppError{msg: errMsg, StatusCode: http.StatusInternalServerError}
}

func (e AppError) Error() string {
	return e.msg
}

func (e AppError) HTTPStatus() int {
	return e.StatusCode
}

func (e AppError) Errors() []string {
	return e.errors
}

// func (e *AppMultiError) Error() string {
// 	return e.appError.msg
// }

// func (e *AppMultiError) HTTPStatus() int {
// 	return e.appError.HTTPStatus()
// }

// func (e *AppMultiError) AddError(err error) {
// 	e.errors = append(e.errors, err)
// }

// func (e *AppMultiError) HasErrors() bool {
// 	return len(e.errors) > 0
// }

var ErrForbidden = AppError{
	msg:        "permission denied",
	StatusCode: http.StatusForbidden,
}

var ErrMissing = AppError{
	msg:        "not found",
	StatusCode: http.StatusNotFound,
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

func ErrNewMissingContext(keyName string) *AppError {
	return &AppError{
		msg:        "context data missing: " + keyName,
		StatusCode: http.StatusInternalServerError,
	}
}

// func NewMultiError(err AppError) *AppMultiError {
// 	return &AppMultiError{
// 		appError: err,
// 		errors:   []error{},
// 	}
// }

// var ErrMissingContext = AppError{
// 	msg:        "context data missing",
// 	StatusCode: http.StatusInternalServerError,
// }

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
