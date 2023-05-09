package jsonencrypter

import (
	"fmt"
	"testing"
)

const (
	testDecryptedFile   = "test/database.crd"
	testUndecryptedFile = "test/test_postgresql.json"
)

func TestDecryptFile(t *testing.T) {
	d, err := DecryptFile(testDecryptedFile, testKeyFile)
	if err != nil {
		t.Errorf("failed to read the decrypted file %s, err >>>> %s", testDecryptedFile, err)
	} else {
		fmt.Println("undecrypted data >>>", string(d))
	}
}

func TestDecryptToFile(t *testing.T) {
	if err := DecryptToFile(testDecryptedFile, testUndecryptedFile, testKeyFile); err != nil {
		t.Errorf("failed to write the undecrypted file %s, err >>>> %s", testUndecryptedFile, err)
	}
}
