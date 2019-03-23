package types

import (
	"fmt"
	"os"
	"time"

	ldap "github.com/UCCNetworkingSociety/netsoc-go-ldap"
	docker "github.com/fsouza/go-dockerclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/sqlite"
)

type ServerGroup struct {
	Database *gorm.DB
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

func NewServerGroup() (*ServerGroup, error) {
	cli, err := docker.NewClient("unix:///var/run/docker.sock")
	/* if err != nil {
		return nil, ServerGroupError{"Docker", err}
	} */
	cli.SetTimeout(time.Second * 3)

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"))
	dbConn, err := gorm.Open("mysql", dbURI)
	/* if err != nil {
		return nil, ServerGroupError{"Database", err}
	} */

	ldapConn, err := ldap.New(ldap.Config{
		BaseDN:   os.Getenv("LDAP_DN"),
		BindUser: os.Getenv("LDAP_USER"),
		BindPass: os.Getenv("LDAP_PASS"),
		Host:     os.Getenv("LDAP_HOST"),
	})
	/* if err != nil {
		return nil, ServerGroupError{"LDAP", err}
	} */

	sqliteConn, err := sqlite.Open(sqlite.ConnectionURL{
		Database: "./tokens.db",
	})
	/* if err != nil {
		return nil, ServerGroupError{"SQLite", err}
	} */
	_ = err
	return &ServerGroup{
		Docker:   cli,
		Database: dbConn,
		LDAP:     ldapConn,
		SQLite:   sqliteConn,
	}, nil
}

func (s *ServerGroup) Close() {
	s.SQLite.Close()
	s.Database.Close()
	s.LDAP.Close()
}
