package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// bcrypt.DefaultCost: The complexity of the hashing algorithm
	// Can put a number between 4 and 31
	// The password can't be more than 72 characters
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
