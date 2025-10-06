package httpserver

import (
	"encoding/json"
	"net/http"
)

type BaseResponse struct {
	Message  string `json:"message,omitempty"`
	Data     any    `json:"data,omitempty"`
	HttpCode int    `json:"http_code"`
}

func (h *Handler) writeError(w http.ResponseWriter, err error) {
	res := BaseResponse{
		Message:  err.Error(),
		HttpCode: http.StatusBadRequest,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)

}

func (h *Handler) writeSuccess(w http.ResponseWriter, data any, httpCode int) {
	res := BaseResponse{
		Data:     data,
		HttpCode: httpCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	responseBody, _ := json.Marshal(res)
	_, _ = w.Write(responseBody)
}
