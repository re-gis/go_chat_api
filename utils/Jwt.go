package utils

import (
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = os.Getenv("JWT_KEY")

func ParseToken(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(jwtKey), nil
    })

    if err != nil {
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check if the "email" claim exists in the token
        if emailClaim, ok := claims["email"].(string); ok {
            // Successfully extracted the email claim as a string
            // Convert it to uint if needed
            return emailClaim, nil
        }

        return "", errors.New("token doesn't contain valid user email")
    }

    return "", errors.New("invalid token")
}
