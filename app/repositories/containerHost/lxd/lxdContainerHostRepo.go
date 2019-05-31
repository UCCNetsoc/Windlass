package lxd

import (
	"context"
	"fmt"

	"github.com/UCCNetworkingSociety/Windlass/app/connections"

	lxd "github.com/lxc/lxd/client"

	host "github.com/UCCNetworkingSociety/Windlass/app/repositories/containerHost"
)

type LXDHost struct {
	ctx  context.Context
	conn lxd.ContainerServer
}

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

	//lxd.conn.Crea

	return nil
}
