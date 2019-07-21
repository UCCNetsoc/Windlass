package config

import (
	"github.com/spf13/viper"
)

func initDefaults() {
	// LDAP settings
	viper.SetDefault("ldap.user", "admin")
	viper.SetDefault("ldap.dn", "dc=netsoc,dc=co")
	viper.SetDefault("ldap.pass", "pass")
	viper.SetDefault("ldap.host", "localhost")

	// MySQL Settings
	viper.SetDefault("db.host", "db")
	viper.SetDefault("db.port", 3306)
	viper.SetDefault("db.user", "netsoc")
	viper.SetDefault("db.pass", "netsoc")
	viper.SetDefault("db.name", "netsoc_admin")

	// Container Host settings
	viper.SetDefault("container.host", "lxd")

	// LXD settings
	viper.SetDefault("lxd.socket", "/var/lib/lxd/unix.socket")

	// Auth settings
	viper.SetDefault("auth.provider", "")

	// Consul settings
	viper.SetDefault("consul.host", "consul:8500")
	viper.SetDefault("consul.token", "") // ACL token
	viper.SetDefault("consul.path", "windlass")

	// Vault settings
	viper.SetDefault("vault.enabled", false)
	viper.SetDefault("vault.token", "")

	// Misc settings
	viper.SetDefault("misc.debug", true)
}
