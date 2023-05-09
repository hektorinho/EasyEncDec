package jsonencrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"os"
)

func DecryptToFile(src string, dst string, key string) error {
	body, err := DecryptFile(src, key)
	if err != nil {
		return err
	}

	if err := os.WriteFile(dst, body, 0700); err != nil {
		return err
	}
	return nil
}

func DecryptFile(src string, key string) ([]byte, error) {
	ciph, err := ioutil.ReadFile(src)
	if err != nil {
		return nil, err
	}

	kd, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	mac, err := GetMacAddress()
	if err != nil {
		return nil, err
	}

	keyData := append(kd, mac...)

	c, err := aes.NewCipher(keyData)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciph) < nonceSize {
		return nil, err
	}

	nonce, ciphertext := ciph[:nonceSize], ciph[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
