module github.com/UCCNetworkingSociety/Windlass

go 1.12

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.4.0

replace github.com/docker/docker v1.13.1 => github.com/docker/docker v0.7.3-0.20190309235953-33c3200e0d16

require (
	code.cloudfoundry.org/systemcerts v0.0.0-20180917154049-ca00b2f806f2 // indirect
	github.com/UCCNetworkingSociety/netsoc-go-ldap v0.0.0-20190120010400-3994c7708032
	github.com/bwmarrin/lit v0.0.0-20190510005413-9c5ce4f3cafc
	github.com/flosch/pongo2 v0.0.0-20181225140029-79872a7b2769 // indirect
	github.com/fsouza/go-dockerclient v1.4.1
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/vault/api v1.0.2
	github.com/juju/clock v0.0.0-20190205081909-9c5c9712527c // indirect
	github.com/juju/retry v0.0.0-20180821225755-9058e192b216 // indirect
	github.com/juju/utils v0.0.0-20180820210520-bf9cc5bdd62d // indirect
	github.com/juju/version v0.0.0-20180108022336-b64dbd566305 // indirect
	github.com/juju/webbrowser v0.0.0-20180907093207-efb9432b2bcb // indirect
	github.com/lxc/lxd v0.0.0-20190412215808-4f5329769215
	github.com/pelletier/go-toml v1.3.0 // indirect
	github.com/rogpeppe/fastuuid v1.0.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	golang.org/x/crypto v0.0.0-20190411191339-88737f569e3a // indirect
	golang.org/x/sys v0.0.0-20190412213103-97732733099d // indirect
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/errgo.v1 v1.0.1 // indirect
	gopkg.in/httprequest.v1 v1.2.0 // indirect
	gopkg.in/ldap.v2 v2.5.1 // indirect
	gopkg.in/macaroon-bakery.v2 v2.1.0 // indirect
	gopkg.in/macaroon.v2 v2.1.0 // indirect
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5 // indirect
	upper.io/db.v3 v3.5.7+incompatible
)
