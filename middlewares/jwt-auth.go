package middlewares

import (
	"PR_gin_g/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER):]

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["name"])
			log.Println("Claims[Admin]: ", claims["admin"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])
		} else {
			log.Println("Token is not valid: ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
