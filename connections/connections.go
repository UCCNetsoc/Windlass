package connections

import (
	"fmt"
	"os"
	"time"

	"upper.io/db.v3/mysql"

	ldap "github.com/UCCNetworkingSociety/netsoc-go-ldap"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/spf13/viper"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type ServerGroup struct {
	Database sqlbuilder.Database
	SQLite   sqlbuilder.Database
	Docker   *docker.Client
	LDAP     *ldap.Conn
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
	cli, err := docker.NewClient("unix:///var/run/docker.sock")
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

	ldapConn, err := ldap.New(ldap.Config{
		BaseDN:   os.Getenv("LDAP_DN"),
		BindUser: os.Getenv("LDAP_USER"),
		BindPass: os.Getenv("LDAP_PASS"),
		Host:     os.Getenv("LDAP_HOST"),
	})
	if err != nil {
		return ServerGroupError{"LDAP", err}
	}

	sqliteConn, err := sqlite.Open(sqlite.ConnectionURL{
		Database: "./sqlite.db",
	})
	if err != nil {
		return ServerGroupError{"SQLite", err}
	}

	Group = ServerGroup{
		SQLite:   sqliteConn,
		LDAP:     ldapConn,
		Docker:   cli,
		Database: mysqlConn,
	}

	return nil
}

func (s *ServerGroup) Close() {
	s.SQLite.Close()
	s.Database.Close()
	s.LDAP.Close()
}
