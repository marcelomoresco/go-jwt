package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/marcelomoresco/go-jwt/initializers"
	model "github.com/marcelomoresco/go-jwt/models"
)

func RequiredAuth(ctx *gin.Context){
	tokenString := ctx.GetHeader("Authorization")


	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		exp:=claims["exp"]
		sub:=claims["sub"]
		time := time.Now().Unix()
		if float64(time) > exp.(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user model.User
		initializers.DB.First(&user,sub)

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		ctx.Set("user",user)
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}