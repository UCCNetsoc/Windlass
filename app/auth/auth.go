package auth

import (
	"github.com/UCCNetworkingSociety/Windlass/app/auth/ldap"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/provider"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/unrestricted"
	"github.com/spf13/viper"
)

func GetProvider() (provider.AuthProvider, error) {
	switch viper.GetString("auth.provider") {
	case "ldap":
		return ldap.Init()
	default:
		return unrestricted.Init()
	}
}
