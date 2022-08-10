package orm

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Picture  []Picture
}

type Picture struct {
	gorm.Model
	UserID      int
	Picture     string
	Picturename string
}
