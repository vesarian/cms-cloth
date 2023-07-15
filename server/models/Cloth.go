package models

type Cloth struct {
	GormModel
	Name        string    `gorm:"column:name;not null" json:"name" form:"name" valid:"required"`
	Description string    `gorm:"column:description" json:"description" form:"description"`
	Price       int       `gorm:"column:price;not null" json:"price" form:"price"`
	ImgUrl      string    `gorm:"column:imgUrl" json:"imgUrl" form:"imgUrl"`
	Stock       int       `gorm:"column:stock;not null" json:"stock" form:"stock"`
	UserId      uint      `gorm:"column:user_id" json:"user_id" form:"user_id"`
	User        *User     `json:"-"`
	CategoryId  uint      `gorm:"column:category_id" json:"category_id" form:"category_id"`
	Category    *Category `json:"-"`
}
