package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dev-hyunsang/daily-todo/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ExtractToken(r *fiber.Ctx) (string, error) {
	bearToken := r.GetReqHeaders()

	jwtString := strings.Split(bearToken["Authorization"], "Bearer ")[1]
	if len(jwtString) == 0 { // 만약 토큰 없는 경우 오류 반환함.
		return "", errors.New("Failed to Reading JSON Web Token.")
	}

	return jwtString, nil
}

func VerifyToken(r *fiber.Ctx) (*jwt.Token, error) {
	tokenString, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetEnv("ACCESS_SECRET")), nil
	})

	return token, err
}

func TokenVailed(r *fiber.Ctx) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
