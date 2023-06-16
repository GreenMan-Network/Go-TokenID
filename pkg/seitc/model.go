package seitc

import "errors"

var (
	ErrHttpStatusNot2xx     = errors.New("http status ")
	ErrMissingMetadata      = errors.New("missing metadata")
	ErrAppendingCertificate = errors.New("error appending cert")
)

type RequestData interface {
	ToJson() ([]byte, error)
}

type ErrorResponse struct {
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

type ResponseStatus struct {
	Status   *int    `json:"status,omitempty"`
	Code     *string `json:"code,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Message  *string `json:"message,omitempty"`
	Info     *string `json:"info,omitempty"`
}
