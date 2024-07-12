package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	bcrypt, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return "", err
	}
	return string(bcrypt), nil

}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
