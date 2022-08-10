package main

import (
	AuthController "example/go-my-imgstore/controller"
	PictureController "example/go-my-imgstore/controller/upload"
	UserController "example/go-my-imgstore/controller/user"
	"example/go-my-imgstore/middleware"
	"example/go-my-imgstore/orm"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Binding from JSON
type Register struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	Username   string
	Password   string
	Avatar     string
	Avatarname string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orm.InitDB()

	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/userall", UserController.UserAll)
	authorized.GET("/profile", UserController.Profile)

	authorized.GET("/profile/pictures", PictureController.GetAllPicture)
	authorized.POST("/profile/pictures", PictureController.UploadPicture)
	authorized.PUT("/profile/pictures", PictureController.UpdatePictureName)
	authorized.DELETE("/profile/pictures/:id", PictureController.DeletePicture)
	r.Run() //"localhost:8080"
}
