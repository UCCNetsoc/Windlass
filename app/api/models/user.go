package common

type User struct {
	Username string
}

func (u User) HomeDir() string {
	return "/home/users/" + u.Username
}
