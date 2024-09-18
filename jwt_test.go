package ipusher

import "testing"

func Test_TwtToken(t *testing.T) {
	token, err := generateJWT(100)

	if err != nil {
		t.Error("Error during generating jwt token")
	}

	if token == "" {
		t.Error("Error: Empty token found")
	}
}
