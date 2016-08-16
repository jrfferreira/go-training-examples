package auth

// Authenticate autentica o usuario
func Authenticate(u User) bool {
	return u.Name == "joao"
}
