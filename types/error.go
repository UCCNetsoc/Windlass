package types

import (
	"fmt"
	"io"
	"time"
)

type APIError struct {
	ErrorID uint8     `json:"error"`
	Status  uint16    `json:"status"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("[%v] - %d - %d %s", e.Time, e.Status, e.ErrorID, e.Message)
}

func (e APIError) Encode(w io.Writer) {

}
