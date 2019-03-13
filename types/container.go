package types

import (
	"github.com/fsouza/go-dockerclient"
)

type Container struct {
	docker.Container
}

