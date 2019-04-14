package auth

import (
	"errors"
	"fmt"
	"github.com/Strum355/viper"
	"github.com/UCCNetworkingSociety/Windlass/auth/ldap"
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
)

func GetProvider() (provider.AuthProvider, error) {
	switch viper.GetString("AUTH_PROVIDER") {
	case "ldap":
		return ldap.Init()
	}
	return nil, errors.New(fmt.Sprintf("auth provider not recognized: %s", viper.GetString("AUTH_PROVIDER")))
}
