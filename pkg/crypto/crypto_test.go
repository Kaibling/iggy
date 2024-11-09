package crypto_test

import (
	"testing"

	"github.com/kaibling/iggy/pkg/crypto"
)

func TestHashPassword(t *testing.T) {

	pwdHash, _ := crypto.HashPassword("password", 11)

	ok, err := crypto.CheckPasswordHash("password", pwdHash)
	if err != nil {
		t.Fail()
	}
	if !ok {
		t.Fail()
	}

}
