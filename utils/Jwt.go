package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = os.Getenv("JWT_KEY")

func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("claims: %+v\n",claims)
    // Check if the "email" claim exists in the token
    if emailClaim, ok := claims["email"]; ok {
        // Check the type of the "email" claim
        switch email := emailClaim.(type) {
        case string:
            // If it's a string, convert it to uint if needed
            // You might need to handle the conversion appropriately
            // For example, you can use strconv.Atoi if the uint representation is in string format
            // Here's a simple example assuming the email is a number
            if emailUint, err := strconv.ParseUint(email, 10, 64); err == nil {
                return uint(emailUint), nil
            }
        case float64:
            // If it's already a float64, convert it to uint if needed
            // You might need to handle the conversion appropriately
            // For example, you can use a type assertion to convert it to uint
            emailUint := uint(email)
            return emailUint, nil
        default:
            // Handle other types if needed
            return 0, errors.New("unexpected type for email claim")
        }
    }

    return 0, errors.New("token doesn't contain valid user email")
}

	return 0, errors.New("invalid token")
}
