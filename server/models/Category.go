package models

type Category struct {
	GormModel
	Name string `gorm:"column:name; not null" json:"name" form:"name" valid:"required"`
}
