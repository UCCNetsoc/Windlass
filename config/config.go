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

	// Docker settings
	v.SetDefault("DOCKER_SOCKET", "unix:///var/run/docker.sock")
}

func Load(v *viper.Viper) {
	initDefaults(v)
	v.AutomaticEnv()
}
