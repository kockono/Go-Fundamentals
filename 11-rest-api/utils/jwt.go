package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {

	// Encrypt and sign the token using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	// Parse the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		panic(err)
	}

	if !parsedToken.Valid {
		panic("Invalid token")
	}

	// Here we can extract the claims from the token
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	panic("Invalid claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
