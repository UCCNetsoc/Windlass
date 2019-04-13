module github.com/UCCNetworkingSociety/Windlass

go 1.12

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.4.0

replace github.com/docker/docker v1.13.1 => github.com/docker/docker v0.7.3-0.20190309235953-33c3200e0d16

require (
	github.com/Sirupsen/logrus v1.4.0 // indirect
	github.com/UCCNetworkingSociety/netsoc-go-ldap v0.0.0-20190120010400-3994c7708032
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/flosch/pongo2 v0.0.0-20181225140029-79872a7b2769 // indirect
	github.com/fsouza/go-dockerclient v1.3.6
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/juju/webbrowser v0.0.0-20180907093207-efb9432b2bcb // indirect
	github.com/lxc/lxd v0.0.0-20190412215808-4f5329769215 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/rogpeppe/fastuuid v1.0.0 // indirect
	github.com/spf13/viper v1.3.2
	golang.org/x/net v0.0.0-20190313220215-9f648a60d977 // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/errgo.v1 v1.0.1 // indirect
	gopkg.in/httprequest.v1 v1.2.0 // indirect
	gopkg.in/ldap.v2 v2.5.1 // indirect
	gopkg.in/macaroon-bakery.v2 v2.1.0 // indirect
	gopkg.in/macaroon.v2 v2.1.0 // indirect
	gopkg.in/robfig/cron.v2 v2.0.0-20150107220207-be2e0b0deed5 // indirect
	upper.io/db.v3 v3.5.7+incompatible
)
