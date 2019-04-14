package provider

type AuthProvider interface {
	Authenticate(user, pass string) (bool, error)
	Close() error
}
