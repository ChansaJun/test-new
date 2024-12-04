package entity

import (
	"time"

	"gorm.io/gorm"
)

type SouvenirShop struct {
	gorm.Model
	GoodsName        string    `json:"goods_name"`
	GoodsPrice       float32   `json:"goods_price"`
	GoodsAmoung      string    `json:"goods_amoung"`
	GoodsDescription string    `json:"goods_description"`
	DateAdded        time.Time `json:"date_added"`
	ImageUrl         string    `json:"image_url"`
}
