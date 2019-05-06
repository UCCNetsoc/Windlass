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

	// MySQL Settings
	v.SetDefault("DB_HOST", "db")
	v.SetDefault("DB_PORT", 3306)
	v.SetDefault("DB_USER", "netsoc")
	v.SetDefault("DB_PASS", "netsoc")
	v.SetDefault("DB_NAME", "netsoc_admin")

	// LXD settings
	v.SetDefault("LXD_SOCKET", "/var/lib/lxd/unix.socket")
}

func Load(v *viper.Viper) {
	initDefaults(v)
	v.AutomaticEnv()
}
