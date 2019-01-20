package types

import (
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/UCCNetworkingSociety/netsoc-go-ldap"
	"time"
	"net/http"
	"github.com/docker/docker/client"
)

type ServerGroup struct {
	Database *gorm.DB
	SQLite sqlbuilder.Database
	Docker *client.Client
	LDAP *ldap.Conn
}

func NewServerGroup() (*ServerGroup, error) {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "", &http.Client{Timeout: time.Second * 5}, nil)
	if err != nil {
		return nil, err
	}

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASS"))
	dbConn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	ldapConn, err := ldap.New(ldap.Config{
		BaseDN: os.Getenv("LDAP_DN"),
		BindUser: os.Getenv("LDAP_USER"),
		BindPass: os.Getenv("LDAP_PASS"),
		Host: os.Getenv("LDAP_HOST"),
	})
	if err != nil {
		return nil, err	
	}

	sqliteConn, err := sqlite.Open(sqlite.ConnectionURL{
		Database: "./tokens.db",
	})
	if err != nil {
		return nil, err
	}

	return &ServerGroup{
		Docker: cli,
		Database: dbConn,
		LDAP: ldapConn,
		SQLite: sqliteConn,
	}, nil
}

func (s *ServerGroup) Close() {
	s.SQLite.Close()
	s.Database.Close()
	s.Docker.Close()
	s.LDAP.Close()
}