package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isongjosiah/lernen-api/tracing"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	requestIDKey     = "request_id"
	requestSourceKey = "request_source"
)

// ServerResponse struct ...
type ServerResponse struct {
	Err        string          `json:"error"`
	Message    string          `json:"message"`
	StatusCode int             `json:"status_code"`
	Context    context.Context `json:"context"`
	Payload    interface{}     `json:"payload"`
}

// ErrorResponse struct ...
type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

// Error returns the response error as a string
/*func (r *ServerResponse) Error() string {
	return r.Err.Error()
}*/

func WriteErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	r := respondWithError(err, statusCode, nil)
	errorResponse, _ := json.Marshal(r)
	WriteJSONResponse(w, r.StatusCode, errorResponse)
}

// WriteJSONResponse writes data and status code to the ResponseWriter
func WriteJSONResponse(r http.ResponseWriter, statusCode int, content []byte) {
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(statusCode)
	_, _ = r.Write(content)
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
		WriteErrorResponse(w, http.StatusInternalServerError, errors.New("Error creating json response"))
		return
	}
	WriteJSONResponse(w, data.StatusCode, result)
}

// respondWithError
func respondWithError(err error, httpStatusCode int, tracingContext *tracing.Context) *ServerResponse {
	message := err.Error()
	if err != nil {
		err = fmt.Errorf(message)
	}

	WithContext(tracingContext).WithFields(
		logrus.Fields{
			"err": err,
		}).Error(message)

	return &ServerResponse{
		Err:        message,
		Message:    "failed",
		StatusCode: httpStatusCode,
		Context:    nil,
		Payload:    nil,
	}

}
