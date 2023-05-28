package lib

import (
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

// func ValidateJWT(token string)
