package types

import (
	"encoding/json"
	"time"
)

type APIResponse struct {
	Status  uint16      `json:"status"`
	Content interface{} `json:"content"`
}

type apiResponse struct {
	Status  uint16      `json:"status"`
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
