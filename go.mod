module github.com/UCCNetworkingSociety/Windlass

go 1.12

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.4.0

//replace github.com/UCCNetworkingSociety/Windlass-worker => ../Windlass-worker

require (
	github.com/Strum355/log v1.0.1
	github.com/UCCNetworkingSociety/Windlass-worker v0.0.0-20190723225326-f33e2aafc1e8
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/vault/api v1.0.2
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0 // indirect
)
