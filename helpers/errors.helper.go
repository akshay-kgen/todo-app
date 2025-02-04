package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ContentTypeHeader = "Content-Type"
	ApplicationJSON   = "application/json"
)

func ParsePaginationParams(r *http.Request) (limit int, offset interface{}, err error) {
	limit = 15 // Default page size
	if limitQParam := r.URL.Query().Get("pageLimit"); limitQParam != "" {
		limit, err = strconv.Atoi(limitQParam)
		if err != nil {
			return limit, offset, fmt.Errorf("failed in parsing limit: %w", err)
		}
	}

	var pageOffsetMap map[string]interface{}
	pageOffset := r.URL.Query().Get("pageOffset")
	if pageOffset != "" {
		decodedPageOffset, err := url.QueryUnescape(pageOffset)
		if err != nil {
			return limit, offset, fmt.Errorf("failed to decode pageOffset: %w", err)
		}
		err = json.Unmarshal([]byte(decodedPageOffset), &pageOffsetMap)
		if err != nil {
			return limit, offset, fmt.Errorf("failed to unmarshal pageOffset: %w", err)
		}
	}

	return limit, pageOffsetMap, nil
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func SendHandlerErrResponse(w http.ResponseWriter, msg string, status int) {
	response := ErrorResponse{Error: msg}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.Header().Set(ContentTypeHeader, ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Internal Server Error"}`)) // Simple JSON response
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
		w.Write([]byte(`{"error":"Internal Server Error"}`)) // Simple JSON response
		return
	}

	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(status)
	w.Write(responseJSON)
}

func (e *CustomError) Is(target *CustomError) bool {
	return e.Error == target.Error && e.ErrCode == target.ErrCode
}
