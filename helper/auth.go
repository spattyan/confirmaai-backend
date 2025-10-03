package helper

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) *Auth {
	return &Auth{
		Secret: secret,
	}
}

func (auth Auth) GenerateToken(id uuid.UUID, email string) (string, error) {

	if id == uuid.Nil || email == "" {
		return "", errors.New("some values are missing to generate token")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(auth.Secret))

	if err != nil {
		return "", errors.New("error signing token")
	}

	return tokenString, nil
}
func (auth Auth) VerifyToken(token string) (domain.User, error) {

	tokenArray := strings.Split(token, " ")

	if len(tokenArray) != 2 {
		return domain.User{}, nil
	}

	tokenString := tokenArray[1]

	if tokenArray[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	finalToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}
		return []byte(auth.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := finalToken.Claims.(jwt.MapClaims); ok && finalToken.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token expired")
		}

		user := domain.User{}
		user.ID = uuid.MustParse(claims["user_id"].(string))
		user.Email = claims["email"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}
func (auth Auth) Authorize(ctx fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	user, err := auth.VerifyToken(authHeader)

	if err == nil && user.ID != uuid.Nil {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}

}
func (auth Auth) GetCurrentUser(ctx fiber.Ctx) domain.User {

	user := ctx.Locals("user").(domain.User)

	return user
}
