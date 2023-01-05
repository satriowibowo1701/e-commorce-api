package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/satriowibowo1701/e-commorce-api/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateToken(id int, role string) (string, error) {
	timer := time.Now().Add(config.JWT_EXPIRATION_DURATION).Unix()
	secret := config.JWT_SECRET
	claims := jwt.MapClaims{}
	claims["role"] = role
	claims["id"] = id
	claims["exp"] = int(timer)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, err
}
