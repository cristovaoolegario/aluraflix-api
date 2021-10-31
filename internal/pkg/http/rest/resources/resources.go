package resources

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ErrorMessage represents a error model
type ErrorMessage struct {
	Error string `json:"error" example:"example error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		response, _ := json.Marshal(payload)
		w.Write(response)
	}
}

func GetQueryParams(queryParams url.Values) (filter string, page int64, pageSize int64) {
	filter = queryParams.Get("search")
	page = 1
	if n, err := strconv.Atoi(queryParams.Get("page")); err == nil {
		page = int64(n)
	}
	pageSize = 5
	if n, err := strconv.Atoi(queryParams.Get("pageSize")); err == nil {
		pageSize = int64(n)
	}
	return filter, page, pageSize
}