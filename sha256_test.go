package minutil

import (
	"log"
	"testing"
)

func TestEncryptAndVerify(t *testing.T) {
	input := "mysecretpassword"
	log.Println("Encrypting input:", input)

	hashed := Encrypt(input)
	log.Println("Hashed value:", hashed)

	// 验证哈希值是否正确
	if !Verify(input, hashed) {
		t.Errorf("Expected hash to match, but it did not")
	} else {
		log.Println("Verification successful for correct input")
	}

	// 验证错误的输入
	wrongInput := "mysecretpassword"
	log.Println("Verifying wrong input:", wrongInput)

	if Verify(wrongInput, hashed) {
		log.Println("Verifying successfull!")
	} else {
		log.Println("Verification failed for wrong input, as expected")
	}
}
