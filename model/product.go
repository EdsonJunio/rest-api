package model

type Product struct {
	ID    int     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name  string  `gorm:"column:product_name" json:"name"`
	Price float64 `gorm:"column:price" json:"price"`
}

func (Product) TableName() string {
	return "product"
}
