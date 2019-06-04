package lxd

import (
	"context"
	"fmt"

	"github.com/UCCNetworkingSociety/Windlass/app/connections"
	"github.com/UCCNetworkingSociety/Windlass/app/helpers"
	host "github.com/UCCNetworkingSociety/Windlass/app/repositories/containerHost"
	lxdclient "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared/api"
)

type LXDHost struct {
	ctx  context.Context
	conn lxdclient.ContainerServer
}

// TODO context tiemouts
func NewLXDRepository() *LXDHost {
	lxdHost, err := connections.GetLXD()
	if err != nil {
		panic(fmt.Sprintf("error getting LXD host: %v", err))
	}

	return &LXDHost{
		ctx:  context.Background(),
		conn: lxdHost,
	}
}

func (lxd *LXDHost) WithContext(ctx context.Context) host.ContainerHostRepository {
	lxd.ctx = ctx
	return lxd
}

func (lxd *LXDHost) Ping() error {
	return nil
}

func (lxd *LXDHost) CreateContainerHost(opts host.ContainerHostCreateOptions) error {
	op, err := lxd.conn.CreateContainer(api.ContainersPost{
		Name: opts.Name,
		Source: api.ContainerSource{
			Type:        "image",
			Fingerprint: "76180b5eb160", // alpine:3.8 image
		},
	})
	if err != nil {
		return err
	}

	return helpers.OperationTimeout(lxd.ctx, op)
}

func (lxd *LXDHost) StartContainerHost(opts host.ContainerHostCreateOptions) error {
	op, err := lxd.conn.UpdateContainerState(opts.Name, api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}, "")
	if err != nil {
		return err
	}

	return helpers.OperationTimeout(lxd.ctx, op)
}
