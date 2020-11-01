package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/isongjosiah/lernen-api/common"
	"github.com/isongjosiah/lernen-api/tracing"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	requestIDKey     = "request_id"
	requestSourceKey = "request_source"
)

// ServerResponse struct ...
type ServerResponse struct {
	Err         error
	Message     string
	StatusCode  int
	Context     context.Context
	ContentType common.ContentType
	Payload     interface{}
}

// ErrorResponse struct ...
type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

// Error returns the response error as a string
func (r *ServerResponse) Error() string {
	return r.Err.Error()
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, errString string) {
	r := respondWithError(errors.New(errString), statusCode, nil)
	errorResponse, _ := json.Marshal(r)
	WriteJSONResponse(w, r.StatusCode, errorResponse)
}

// WriteJSONResponse writes data and statuus code to the ResponseWriter
func WriteJSONResponse(r http.ResponseWriter, statusCode int, content []byte) {
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)
	r.Write(content)
}

// WithContext
func WithContext(tracingContext *tracing.Context) *logrus.Entry {
	fields := logrus.Fields{}

	if tracingContext != nil {
		fields[requestIDKey] = tracingContext.RequestID
		fields[requestSourceKey] = tracingContext.RequestSource
	}

	return logrus.WithFields(fields)
}

// RespondWithJSONPayload ...
func WriteJSONPayload(w http.ResponseWriter, data *ServerResponse) {
	result, err := json.Marshal(data)
	if err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Error creating json response")
		return
	}
	WriteJSONResponse(w, data.StatusCode, result)
}

// respondWithError
func respondWithError(err error, httpStatusCode int, tracingContext *tracing.Context) *ServerResponse {
	message := err.Error()
	if err == nil {
		err = fmt.Errorf(message)
	}

	WithContext(tracingContext).WithFields(
		logrus.Fields{
			"err": err,
		}).Error(message)

	return &ServerResponse{
		Err:        fmt.Errorf("%v, request-id: %v", message, tracingContext.RequestID),
		StatusCode: httpStatusCode,
		Message:    message,
	}

}
