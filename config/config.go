package config

import (
	"github.com/spf13/viper"
)

func InitDefaults(v *viper.Viper) {
	v.SetDefault("LDAP_USER", "admin")
	v.SetDefault("LDAP_DN", "dc=netsoc,dc=co")
	v.SetDefault("LDAP_PASS", "pass")
	v.SetDefault("LDAP_HOST", "localhost")
}

func Load(v *viper.Viper) {
	InitDefaults(v)
	v.AutomaticEnv()
}
