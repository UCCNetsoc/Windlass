package auth

import (
	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/auth/ldap"
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
	"github.com/UCCNetworkingSociety/Windlass/auth/unrestricted"
)

func GetProvider() (provider.AuthProvider, error) {
	switch viper.GetString("AUTH_PROVIDER") {
	case "ldap":
		return ldap.Init()
	default:
		return unrestricted.Init()
	}
}
