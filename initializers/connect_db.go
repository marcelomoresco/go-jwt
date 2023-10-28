package initializers

import (
	"os"

	model "github.com/marcelomoresco/go-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(){
	var err error
	dsn:=os.Getenv("DB")
	DB, err := gorm.Open(postgres.New(postgres.Config{
	DSN: dsn,
	}))

	if err != nil{
		panic("Errro conect db")
	}
	DB.AutoMigrate(&model.User{})


}