package server

import (
	"github.com/UCCNetworkingSociety/netsoc-go-ldap"
	"log"
	"time"
	"net/http"
	"github.com/docker/docker/client"
)

type ServerGroup struct {
	// <some database here>
	Docker *client.Client
	LDAP *ldap.Conn
}

var (
	Server *ServerGroup
)

func init() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "", &http.Client{Timeout: time.Second * 5}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	Server = &ServerGroup{
		Docker: cli,
	}
}