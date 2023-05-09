package easyencdec

import (
	"fmt"
	"os"
	"testing"
)

const (
	n              = 15
	testKeyFile    = "test/private.key"
	testInputFile  = "test/postgresql.json"
	testOutputFile = "test/database.crd"
)

func TestGenerateKey(t *testing.T) {
	bytes := GenerateKey(n)
	if len(bytes) != n {
		t.Errorf("len(bytes) = %d; want %d", len(bytes), n)
	}
}

func TestGenerateKeyFile(t *testing.T) {
	if err := GenerateKeyFile(testKeyFile); err != nil {
		t.Errorf("wasn't able to write %s", testKeyFile)
	}
}

func TestGetMacAddress(t *testing.T) {
	b, err := GetMacAddress()
	if err != nil {
		t.Errorf("wasn't able to get MAC address of device %s", string(b))
	}
}

func TestCombinedKeyAndMac(t *testing.T) {
	c, err := CombineKeyAndMac(testKeyFile)
	if err != nil {
		t.Errorf("wasn't able get combined key and mac >>>>> %s", err)
	}
	if len(c) != 32 {
		t.Errorf("len of key not correct got >> %d expected >> %d", len(c), 32)
	}
}

func TestEncryptFile(t *testing.T) {
	if err := EncryptFile(testInputFile, testOutputFile, testKeyFile); err != nil {
		t.Errorf("wasn't able to encrypt file >>>>> success=%t", err)
	}
	if fi, err := os.Stat(testOutputFile); !os.IsNotExist(err) {
		fmt.Printf("EncryptFile: successfully encrypted file %s >>> filename=%s, filesize=%d, filemode=%s\n", testOutputFile, fi.Name(), fi.Size(), fi.Mode())
	}
}
