package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
}

type apiResponse struct {
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
	Time    time.Time   `json:"time"`
}

func (resp APIResponse) MarshalJSON() ([]byte, error) {
	timed := apiResponse{
		Status:  resp.Status,
		Content: resp.Content,
		Time:    time.Now(),
	}
	return json.Marshal(timed)
}

func (resp APIResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(resp.Status)
	return nil
}
