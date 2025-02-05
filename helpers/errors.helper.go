package helpers

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeHeader = "Content-Type"
	ApplicationJSON   = "application/json"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendHandlerErrResponse(w http.ResponseWriter, msg string, status int) {
	response := ErrorResponse{Error: msg}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}

	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(status)
	w.Write(responseJSON)
}

type CustomError struct {
	Error   error  `json:"error"`
	ErrCode string `json:"errCode"`
}

type CustomErrorResponse struct {
	Error   string `json:"error"`
	ErrCode string `json:"errCode"`
}

func NewCustomError(err error, errCode string) *CustomError {
	return &CustomError{
		Error:   err,
		ErrCode: errCode,
	}
}

func SendHandlerCustomErrResponse(w http.ResponseWriter, customErr *CustomError, status int) {
	responseJSON, err := json.Marshal(
		CustomErrorResponse{Error: customErr.Error.Error(), ErrCode: customErr.ErrCode},
	)
	if err != nil {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`))
		return
	}

	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(status)
	w.Write(responseJSON)
}

func (e *CustomError) Is(target *CustomError) bool {
	return e.Error == target.Error && e.ErrCode == target.ErrCode
}
