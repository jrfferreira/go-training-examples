package auth

import "testing"

func TestAuthenticate(t *testing.T) {
	cases := map[string]bool{
		"joao":   true,
		"lorena": false,
		"babi":   false,
	}

	for username, expected := range cases {
		u := User{Name: username}
		result, _ := Authenticate(u)
		if result != expected {
			t.Errorf("Login to user %s should return %t, but got %t", username, expected, result)
		}
	}
}
