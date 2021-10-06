package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/golang-jwt/jwt"
	"github.com/vavilen84/gocommerce/env"
	"github.com/vavilen84/gocommerce/logger"
	"github.com/vavilen84/gocommerce/models"
	"time"
)

const (
	tokenName = "AccessToken"
)

type CustomClaims struct {
	IsLoggedIn     bool   `json:"isLoggedIn"`
	User           string `json:"user"`
	StandardClaims jwt.StandardClaims
}

func (c CustomClaims) Valid() error {
	return nil
}

func LoginHandler(u models.User, Ctx *context.Context) {
	userBytes, err := json.Marshal(u)
	if err != nil {
		logger.LogError(err)
	}
	now := time.Now()
	claims := CustomClaims{
		IsLoggedIn: true,
		User:       string(userBytes),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(24 * 30 * 12 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(env.GetSecret()))
	Ctx.SetCookie(tokenName, tokenString)
}

func ParseToken(Ctx *context.Context) (jwt.MapClaims, error) {
	tokenString := Ctx.GetCookie(tokenName)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			e := errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
			logger.LogError(e)
			return nil, e
		}
		return []byte(env.GetSecret()), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, nil
	} else {
		logger.LogError(err)
		return nil, err
	}
}

func Logout(Ctx *context.Context) {
	Ctx.SetCookie(tokenName, string(""))
}
