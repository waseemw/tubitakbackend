package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//6 digits code is generated!
func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

// TokenGenerator  is generated here
func TokenGenerator() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	currentTime := time.Now().Format("Mon Jan _2 15:04:05 2006")
	b = append(b, []byte(currentTime)...)

	return hex.EncodeToString(b)
}

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}

func Checkpassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
