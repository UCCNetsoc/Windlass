package helpers

import (
	"context"

	lxd "github.com/lxc/lxd/client"
)

func OperationChannel(op lxd.Operation) <-chan error {
	channel := make(chan error)
	go func() {
		channel <- op.Wait()
	}()
	return channel
}

func OperationTimeout(ctx context.Context, op lxd.Operation) error {
	select {
	case err := <-OperationChannel(op):
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
