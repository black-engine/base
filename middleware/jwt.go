package middleware

import (
	"github.com/black-engine/base/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"os"
)
var TokenCookieName = "t"
var JwtKey = []byte( os.Getenv("JWT_SECRET") )

type Claims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func JwtMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		validToken := false

		if authorization := context.GetHeader( headers.Authorization ); len(authorization) > 0 {
			claims := &Claims{}
			if token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (i interface{}, e error) {
				return JwtKey, nil
			}); err == nil {
				validToken = true
				context.Set(TokenCookieName, token.Claims)
			}
		}

		if authorization, er := context.Cookie( TokenCookieName ); er == nil {
			claims := &Claims{}
			if token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (i interface{}, e error) {
				return JwtKey, nil
			}); err == nil {
				validToken = true
				context.Set(TokenCookieName, token.Claims)
			}
		}

		if !validToken {
			context.JSON(403, constants.InvalidTokenError )
			context.Abort()
			return
		}

		context.Next()
	}
}
