package config

import (
	"encoding/json"
	"strings"

	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/spf13/viper"
)

// Loads settings used throughout the program. Loads from environment variables before
// optionally loading from other sources
func Load() error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	initDefaults()
	viper.AutomaticEnv()

	if err := setSharedSecret(); err != nil {
		return err
	}

	printSettings()
	return nil
}

func printSettings() {
	// Print settings with secrets redacted
	settings := viper.AllSettings()
	settings["ldap"].(map[string]interface{})["pass"] = "[redacted]"
	settings["db"].(map[string]interface{})["pass"] = "[redacted]"
	settings["windlass"].(map[string]interface{})["secret"] = "[redacted]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config:\n%s", string(out))
}
