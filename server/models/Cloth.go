package models

type Cloth struct {
	GormModel
	Name        string `gorm:"column:name;not null" json:"name" form:"name" valid:"required"`
	Description string `gorm:"column:description" json:"description" form:"description"`
	Price       int    `gorm:"column:price;not null" json:"price" form:"price"`
	ImgUrl      string `gorm:"column:imgUrl" json:"imgUrl" form:"imgUrl"`
	Stock       int    `gorm:"column:stock;not null" json:"stock" form:"stock"`
	UserId    uint   `gorm:"column:UserId" json:"UserId" form:"UserId`
	User        *User  `json:"-"`
	CategoryId  uint   `gorm:"column:categoryId" json:"categoryId" form:"categoryId"`
	Category    *Category
}
