package tiny_aes

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	value, err := Encrypt("xxx", "TEST")
	if err != nil {
		t.Error("Encrypt Faild", err)
	}
	fmt.Println("Encrypt Value", value)
}

func Test_EncryptDecrypt(t *testing.T) {
	key := "abapollof97e41279b0538cef5716413"
	value := "TEST"
	encryptValue, err := Encrypt(key, value)
	if err != nil {
		t.Error("encrypt faild")
	}
	fmt.Println("encryptValue", encryptValue)

	decryptValue, err := Decrypt(key, encryptValue)
	if err != nil {
		t.Error("decrypt faild", err)
	}

	if decryptValue != value {
		t.Error("decrypt invalid", decryptValue, value)
	}
}

func Test_Decrypt_CPP(t *testing.T) {
	key := "abapollof97e41279b0538cef5716413"
	value := "TEST"
	encryptValue := "YXNlX2l2X2hlYWQ6X19fX96Nh/XiWw7bFTUXaMYZALUAAAAE1mtXOOGFZ9tCS4m+CvYFdQ=="

	decryptValue, err := Decrypt(key, encryptValue)
	if err != nil {
		t.Error("decrypt faild", err)
	}

	if decryptValue != value {
		t.Error("decrypt invalid", decryptValue, value)
	}
}
