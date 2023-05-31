package lib

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/drdofx/talk-parmad/internal/api/models"
	"github.com/spf13/viper"
)

type JWT struct {
	jwt.StandardClaims
	Data UserData `json:"data"`
}

type UserData struct {
	UserID uint
	NIM    string
	Role   string
}

func GenerateJWT(user *models.User) string {
	claims := JWT{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(viper.GetInt("DURATION_TOKEN_JWT"))).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(int(user.ID)),
		},
		Data: UserData{
			UserID: user.ID,
			NIM:    *user.NIM,
			Role:   user.Role,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))

	if err != nil {
		return ""
	}

	return signedToken

}

func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})
}
