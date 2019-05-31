package connections

import (
	"fmt"
)

type ConnectionsError struct {
	component string
	err       error
}

func NewConnectionError(err error, component string) error {
	return ConnectionsError{component, err}
}

func (e ConnectionsError) Error() string {
	return fmt.Sprintf("error connecting to %s: %v", e.component, e.err)
}

func (e ConnectionsError) Component() string {
	return e.component
}
