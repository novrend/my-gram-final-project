package helpers

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getSecretKey() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	return os.Getenv("SECRET_KEY")
}

func GenerateToken(id uint, email string, username string) string {
	claims := jwt.MapClaims{
		"id":       id,
		"email":    email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(getSecretKey()))

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errReponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errReponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errReponse
		}
		return []byte(getSecretKey()), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errReponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
