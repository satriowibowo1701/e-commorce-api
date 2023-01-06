package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/satriowibowo1701/e-commorce-api/config"
	"github.com/satriowibowo1701/e-commorce-api/model"

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

func Authorization(role string, id int, r *http.Request) error {
	cookieid, err := r.Cookie("id")
	if err != nil {
		return err
	}
	cookierole, err1 := r.Cookie("role")
	if err != nil {
		return err1
	}
	newcookiid, _ := strconv.Atoi(cookieid.Value)
	if newcookiid != id && cookierole.Value != role {
		return errors.New("Unauthorized")
	}
	return nil

}
func Authentication(r *http.Request) error {
	token, err := VerifyToken(r)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := TakeToken(r)

	if len(tokenString) == 0 {
		return nil, errors.New("Token not found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TakeToken(r *http.Request) string {
	keys := r.Header.Get("Authorization")
	barearkeys := strings.Split(keys, " ")
	if len(barearkeys) > 1 {
		return barearkeys[1]
	} else {
		return ""
	}

}

func ClaimsAuthToken(r *http.Request) (*model.Authorization, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		id, ok := claims["id"].(float64)
		if !ok {
			return nil, errors.New("Invalid token ")
		}
		role, ok := claims["role"].(string)
		if !ok {
			return nil, errors.New("Invalid token")
		}
		authDetail := &model.Authorization{
			Id:   int(id),
			Role: role,
		}

		return authDetail, nil
	}

	return nil, errors.New("UnAuthorized")
}
