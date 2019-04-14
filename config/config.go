package config

import (
	"github.com/Strum355/viper"
)

func initDefaults(v *viper.Viper) {
	// LDAP settings
	v.SetDefault("LDAP_USER", "admin")
	v.SetDefault("LDAP_DN", "dc=netsoc,dc=co")
	v.SetDefault("LDAP_PASS", "pass")
	v.SetDefault("LDAP_HOST", "localhost")

	// LXD settings
	v.SetDefault("LXD_SOCKET", "/var/lib/lxd/unix.socket")
}

func Load(v *viper.Viper) {
	initDefaults(v)
	v.AutomaticEnv()
}
