package auth

// Authenticate autentica o usuario
func Authenticate(u User) bool {
	if u.Name != "joao" {
		panic("usuario invalido")
	}
	return true
}
