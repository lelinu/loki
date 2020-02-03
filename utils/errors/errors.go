package errors

import "net/http"

type ApiErrorInterface interface{
	Status() int
	Message() string
	Error() error
}

type apiError struct {
	AStatus int  `json:"status"`
	AMessage string  `json:"message"`
	AnError error `json:"error,omitempty"`
}

func (e *apiError) Status() int{
	return e.AStatus
}

func (e *apiError) Message() string{
	return e.AMessage
}

func (e *apiError) Error() error{
	return e.AnError
}

func NewApiError(statusCode int, message string) ApiErrorInterface{
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

func NewInternalServerError(message string) ApiErrorInterface{
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewNotFoundError(message string) ApiErrorInterface{
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

func NewBadRequestError(message string) ApiErrorInterface{
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}