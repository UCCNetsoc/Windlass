package config

import (
	"encoding/json"
	"strings"

	"github.com/Strum355/log"
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

	return nil
}

func PrintSettings() {
	// Print settings with secrets redacted
	settings := viper.AllSettings()
	settings["windlass"].(map[string]interface{})["secret"] = "[redacted]"

	out, _ := json.MarshalIndent(settings, "", "\t")
	log.Debug("config:\n" + string(out))
}
