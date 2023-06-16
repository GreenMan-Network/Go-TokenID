package jwe

import (
	"log"
	"testing"
)

func TestCreateJWE(t *testing.T) {
	t.Log("TestCreateJWE")

	jweString, err := CreteJWE([]byte("test"))
	if err != nil {
		t.Errorf("Error creating JWE: %v", err)
		return
	}

	log.Println("JWE: ", jweString)
}
