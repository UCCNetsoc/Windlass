package config

import (
	"encoding/json"
	"strings"

	"github.com/Strum355/viper"
	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
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
}

func Load() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	initDefaults()
	viper.AutomaticEnv()

	// Print settings with secrets redacted
	settings := viper.AllSettings()
	settings["ldap"].(map[string]interface{})["pass"] = "[redacted]"
	settings["db"].(map[string]interface{})["pass"] = "[redacted]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config: %s", string(out))
}
