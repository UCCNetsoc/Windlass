package config

import (
	"encoding/json"
	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/utils/logging"
)

func initDefaults() {
	// LDAP settings
	viper.SetDefault("LDAP_USER", "admin")
	viper.SetDefault("LDAP_DN", "dc=netsoc,dc=co")
	viper.SetDefault("LDAP_PASS", "pass")
	viper.SetDefault("LDAP_HOST", "localhost")

	// MySQL Settings
	viper.SetDefault("DB_HOST", "db")
	viper.SetDefault("DB_PORT", 3306)
	viper.SetDefault("DB_USER", "netsoc")
	viper.SetDefault("DB_PASS", "netsoc")
	viper.SetDefault("DB_NAME", "netsoc_admin")

	// LXD settings
	viper.SetDefault("LXD_SOCKET", "/var/lib/lxd/unix.socket")
}

func Load() {
	initDefaults()
	viper.AutomaticEnv()

	// Print settings with secrets redacted
	// Lowercase because viper lowercases everything it takes in
	settings := viper.AllSettings()
	settings["ldap_pass"] = "[redacted]"
	settings["db_pass"] = "[redacted]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config: %s", string(out))
}
