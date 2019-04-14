package connections

import (
	"fmt"
	"time"

	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/auth"
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
	docker "github.com/fsouza/go-dockerclient"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
	"upper.io/db.v3/sqlite"
)

type ServerGroup struct {
	Database sqlbuilder.Database
	SQLite   sqlbuilder.Database
	Docker   *docker.Client
	Auth     provider.AuthProvider
}

type ServerGroupError struct {
	component string
	err       error
}

func (e ServerGroupError) Error() string {
	return fmt.Sprintf("%s: %v", e.component, e.err)
}

func (e ServerGroupError) Component() string {
	return e.component
}

var Group ServerGroup

func EstablishConnections() error {
	cli, err := docker.NewClient(viper.GetString("DOCKER_SOCKET"))
	if err != nil {
		return ServerGroupError{"Docker", err}
	}

	cli.SetTimeout(time.Second * 3)

	mysqlConn, err := mysql.Open(mysql.ConnectionURL{
		Host:     viper.GetString("DB_HOST"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASS"),
		Database: viper.GetString("DB_NAME"),
	})
	if err != nil {
		return ServerGroupError{"MySQL", err}
	}

	authProvider, err := auth.GetProvider()
	if err != nil {
		return ServerGroupError{"Auth", err}
	}

	_, err = sqlite.Open(sqlite.ConnectionURL{
		Database: "./sqlite.db",
	})
	if err != nil {
		return ServerGroupError{"SQLite", err}
	}

	Group = ServerGroup{
		Auth: authProvider,
		//SQLite:   sqliteConn,
		Docker:   cli,
		Database: mysqlConn,
	}

	return nil
}

func (s *ServerGroup) Close() {
	//s.SQLite.Close()
	s.Database.Close()
	s.Auth.Close()
}
