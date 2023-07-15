package models

type User struct {
	GormModel
	Username string `gorm:"column:username;not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"column:email;not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"column:password;not null" json:"password" form:"password" valid:"required"`
	Role     string `gorm:"column:role;not null" json:"role" form:"role" valid:"required"`
}
