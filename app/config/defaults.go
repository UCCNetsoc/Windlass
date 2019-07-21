package config

import (
	"github.com/spf13/viper"
)

func initDefaults() {
	viper.SetDefault("windlass.secret", "sample_text")
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
	viper.SetDefault("vault.enabled", false) // If enabled, gets dynamic secret to access Consul from Vault
	viper.SetDefault("vault.token", "")

	// Misc settings
	viper.SetDefault("misc.debug", true)
}
