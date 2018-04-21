package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type QueryInfo struct {
	QueryName    string `gorm:"primary_key"`
	ShopName     string `gorm:"type:varchar(16);not null;"`
	ProductName  string `gorm:"type:varchar(128);not null;"`
	ProductPrice int    `gorm:"type:integer;not null;"`
	UpdateTime   time.Time
}

type ProductInfo struct {
	ShopName     string `gorm:"type:varchar(16);not null;"`
	ProductName  string `gorm:"type:varchar(128);not null;"`
	ProductPrice int    `gorm:"type:integer;not null;"`
}
