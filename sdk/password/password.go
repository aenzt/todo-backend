package password

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func Hash(password string) (string, error) {
	saltRound, err := strconv.Atoi(os.Getenv("BCRYPT_SALT_ROUND"))
	if err != nil {
		return "", err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), saltRound)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Compare(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
