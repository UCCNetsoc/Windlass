package types

import (
	"fmt"
)

type APIError struct {
	ErrorID uint8 `json:"error"`
	APIResponse
}

func (e APIError) Error() string {
	return fmt.Sprintf("%d - %d %v", e.Status, e.ErrorID, e.Content)
}
