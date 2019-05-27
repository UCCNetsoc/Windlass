package auth

import (
	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/ldap"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/provider"
	"github.com/UCCNetworkingSociety/Windlass/app/auth/unrestricted"
)

func GetProvider() (provider.AuthProvider, error) {
	switch viper.GetString("auth.provider") {
	case "ldap":
		return ldap.Init()
	default:
		return unrestricted.Init()
	}
}
