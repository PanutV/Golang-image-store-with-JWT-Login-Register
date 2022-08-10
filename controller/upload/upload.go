package upload

import (
	"example/go-my-imgstore/orm"

	"net/http"

	"github.com/gin-gonic/gin"
)

type any = interface{}

type UploadBody struct {
	Picture     string `form:"picture" json:"picture" binding:"required"`
	Picturename string `form:"picturename" json:"picturename" binding:"required"`
}

func GetAllPicture(c *gin.Context) {
	userId := c.MustGet("userId")
	var pictures []orm.Picture
	orm.Db.Where("user_id = ?", userId).Find(&pictures)
	c.JSON(http.StatusOK, pictures)
}

func UploadPicture(c *gin.Context) {

	var picture UploadBody
	if err := c.ShouldBindJSON(&picture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.MustGet("userId")
	value := userId.(float64)
	var y int = int(value)

	pic := orm.Picture{Picture: picture.Picture, Picturename: picture.Picturename, UserID: y}
	result := orm.Db.Create(&pic)
	c.JSON(200, gin.H{"RowsAffected": result.RowsAffected, "Error": result.Error, "pic": pic})

}

func UpdatePictureName(c *gin.Context) {
	var picture orm.Picture
	var updatedPicture orm.Picture
	if err := c.ShouldBindJSON(&picture); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orm.Db.First(&updatedPicture, picture.ID)
	updatedPicture.Picturename = picture.Picturename
	orm.Db.Save(updatedPicture)
	c.JSON(200, gin.H{"Status": "Updated Picture's Name", "updatedPicture": updatedPicture})
}
func DeletePicture(c *gin.Context) {
	id := c.Param("id")
	var picture orm.Picture
	orm.Db.First(&picture, id)
	orm.Db.Delete(&picture)
	userId := c.MustGet("userId")
	var pictures []orm.Picture
	orm.Db.Where("user_id = ?", userId).Find(&pictures)
	c.JSON(http.StatusOK, pictures)
}

