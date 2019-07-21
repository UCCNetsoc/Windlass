module github.com/UCCNetworkingSociety/Windlass

go 1.12

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.4.0

replace github.com/docker/docker v1.13.1 => github.com/docker/docker v0.7.3-0.20190309235953-33c3200e0d16

require (
	github.com/bwmarrin/lit v0.0.0-20190510005413-9c5ce4f3cafc
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/vault/api v1.0.2
	github.com/pelletier/go-toml v1.3.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sys v0.0.0-20190712062909-fae7ac547cb7 // indirect
)
