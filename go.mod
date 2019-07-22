module github.com/UCCNetworkingSociety/Windlass

go 1.12

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.4.0

replace github.com/docker/docker v1.13.1 => github.com/docker/docker v0.7.3-0.20190309235953-33c3200e0d16

require (
	github.com/UCCNetworkingSociety/Windlass-worker v0.0.0-20190721205052-aa796cdacfff
	github.com/bwmarrin/lit v0.0.0-20190510005413-9c5ce4f3cafc
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1 // indirect
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/vault/api v1.0.2
	github.com/spf13/viper v1.4.0
)
