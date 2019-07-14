package config

import (
	"encoding/json"
	"strings"

	"github.com/UCCNetworkingSociety/Windlass/app/connections"

	log "github.com/UCCNetworkingSociety/Windlass/utils/logging"
	"github.com/spf13/viper"
)

// Loads settings used throughout the program. Loads from environment variables before
// optionally loading from other sources
func Load() error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	initDefaults()
	viper.AutomaticEnv()

	if viper.GetBool("consul.enabled") {
		loadFromConsul()
	}

	printSettings()
	return nil
}

func loadFromConsul() error {
	client, err := connections.GetConsul()
	if err != nil {
		return err
	}

	client.KV()

	return nil
}

func printSettings() {
	// Print settings with secrets redacted
	settings := viper.AllSettings()
	settings["ldap"].(map[string]interface{})["pass"] = "[redacted]"
	settings["db"].(map[string]interface{})["pass"] = "[redacted]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config:\n%s", string(out))
}
