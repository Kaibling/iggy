package crypto

import "testing"

func TestHashPassword(t *testing.T) {

	pwdHash, _ := HashPassword("password", 11)

	ok, err := CheckPasswordHash("password", pwdHash)
	if err != nil {
		t.Fail()
	}
	if !ok {
		t.Fail()
	}

}
