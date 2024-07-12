package utils

import (
	"fmt"
	"os"
	"restful-api-gofiber/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JWT_SECRET string
var JWT_EXPIRES_IN uint

func GenerateJWT(user models.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	// JWT_EXPIRES_IN = os.Getenv("JWT_EXPIRES_IN")
	fmt.Print(JWT_SECRET)

	payload := jwt.MapClaims{}

	payload["id"] = user.Id
	payload["email"] = user.Email
	if payload["email"] == "admin030@gmail.com" {
		payload["role"] = "admin"
	} else {
		payload["role"] = "user"
	}
	payload["exp"] = time.Now().Add(time.Second * time.Duration(172000)).Unix()

	fmt.Println(payload)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	webtoken, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	fmt.Println(webtoken)
	return webtoken, nil
}

func VerifyJWT(encodeToken string) (*jwt.Token, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	JWT_SECRET = os.Getenv("JWT_SECRET")
	fmt.Println(JWT_SECRET)
	token, err := jwt.Parse(encodeToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyJWT(tokenString)
	if err != nil {
		return nil, err
	}
	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
