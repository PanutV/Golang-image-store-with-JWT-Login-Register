package user

import (
	"example/go-my-imgstore/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"Ststus": "ok", "users": users})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId")
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"Ststus": "ok", "user": user})
}
