package rsvp

import (
	"crypto/aes"
	"encoding/hex"
)

func EncryptAES(text string) (string, error) {
	// initialize the aes with the SECRET_KEY
	c, err := aes.NewCipher([]byte(SETTINGS.SECRET_KEY))
	if err != nil {
		return "", err
	}

	// allocate memory for the ciphertext
	cipher := make([]byte, len(text))
	c.Encrypt(cipher, []byte(text))

	// convert to HEX string and return
	return hex.EncodeToString(cipher), nil
}

func DecryptAES(cipher string) (string, error) {
	// convert the HEX String to []byte
	ciphertext, err := hex.DecodeString(cipher)
	if err != nil {
		return "", err
	}

	// initialize the aes with SECRET_KEY
	c, err2 := aes.NewCipher([]byte(SETTINGS.SECRET_KEY))
	if err2 != nil {
		return "", err2
	}

	// allocate space for plaintext
	plaintext := make([]byte, len(ciphertext))
	// decrypt the cipher into plaintext
	c.Decrypt(plaintext, ciphertext)

	// convert the plaintext []bytes into string
	plaintextStr := string(plaintext)

	return plaintextStr, nil
}
