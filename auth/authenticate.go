package auth

type InvalidUser struct{}

func (InvalidUser) Error() string {
	return "Usuário Inválido"
}

// Authenticate autentica o usuario
func Authenticate(u User) (bool, error) {
	if u.Name != "joao" {
		return false, &InvalidUser{}
	}
	return true, nil
}
