package auth

import "errors"

// Authenticate autentica o usuario
func Authenticate(u User) (bool, error) {
	if u.Name != "joao" {
		return false, errors.New("Usuário Inválido")
	}
	return true, nil
}
