package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPassword(hashPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}
