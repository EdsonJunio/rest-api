package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewConflictError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "conflict",
		Code:    http.StatusConflict,
	}
}

func NewPayloadTooLargeError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "payload_too_large",
		Code:    http.StatusRequestEntityTooLarge,
	}
}

func NewUnsupportedMediaTypeError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "unsupported_media_type",
		Code:    http.StatusUnsupportedMediaType,
	}
}

func NewNotImplementedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not_implemented",
		Code:    http.StatusNotImplemented,
	}
}

func NewServiceUnavailableError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "service_unavailable",
		Code:    http.StatusServiceUnavailable,
	}
}
