package unrestricted

import (
	"github.com/UCCNetworkingSociety/Windlass/auth/provider"
)

type UnrestrictedAuthProvider struct{}

func Init() (provider.AuthProvider, error) {
	return UnrestrictedAuthProvider{}, nil
}

func (u UnrestrictedAuthProvider) Authenticate(user, pass string) (bool, error) {
	return true, nil
}

func (u UnrestrictedAuthProvider) Close() error {
	return nil
}
