package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/marcelomoresco/go-jwt/initializers"
	model "github.com/marcelomoresco/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context){
	var body struct {
		Name string
		Email string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read body",
		})
		return 
	}

	hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err !=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to hash password",
		})
		return 
	}

	user := model.User{Name: body.Name,Email: body.Email,Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway,gin.H{
					"error": "Failed create user",
		})
		return 
	}

	ctx.JSON(http.StatusOK,gin.H{})
}

func Login(ctx *gin.Context){
	var body struct {
		Email string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": "Failed to read body",
		})
		return 
	}
	var user model.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID==0{
		ctx.JSON(http.StatusNotFound,gin.H{})
		return 
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "Failed to hash password",
			})
		return 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	secret :=[] byte(os.Getenv("SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
				"error": "Failed to create token",
			})
		return 
	}

	ctx.JSON(http.StatusOK,gin.H{
		"token":tokenString,
	})
}

func Validate(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message": "I`m logged",
	})
}