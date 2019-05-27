package ldap

import (
	"github.com/Strum355/viper"

	"github.com/UCCNetworkingSociety/Windlass/app/auth/provider"
	ldap "github.com/UCCNetworkingSociety/netsoc-go-ldap"
)

type LDAPAuthProvider struct {
	conn *ldap.Conn
}

func Init() (provider.AuthProvider, error) {
	ldapConn, err := ldap.New(ldap.Config{
		BaseDN:   viper.GetString("ldap.dn"),
		BindUser: viper.GetString("ldap.user"),
		BindPass: viper.GetString("ldap.pass"),
		Host:     viper.GetString("ldap.host"),
	})
	if err != nil {
		return nil, err
	}
	return LDAPAuthProvider{ldapConn}, nil
}

func (l LDAPAuthProvider) Authenticate(user, pass string) (bool, error) {
	ldapUser, err := l.conn.GetUser(user)
	if err != nil {
		return false, err
	}

	return l.conn.VerifyPassword(pass, ldapUser)
}

func (l LDAPAuthProvider) Close() error {
	l.conn.Close()
	return nil
}
