package auth

import "testing"

func TestAuthenticate(t *testing.T) {
	u := User{Name: "joao"}
	if a, _ := Authenticate(u); a == false {
		t.Error("Authenticate should return true")
	}

	u2 := User{Name: "Lorena"}
	if a, _ := Authenticate(u2); a == true {
		t.Error("Authenticate should return false")
	}
}
