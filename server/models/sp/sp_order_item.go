package sp

import (
	"time"
)

type SpOrderItem struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title         string    `gorm:"size:200;not null;default:''" json:"title"`
    SkuTitle      string    `gorm:"size:200;not null;default:''" json:"skuTitle"`
    SkuCode       string    `gorm:"size:200;not null;default:''" json:"skuCode"`
    Thumb         string    `gorm:"size:500;not null;default:''" json:"thumb"`
    OrderID       uint      `gorm:"not null;default:0" json:"orderId"`
    ProductID     uint      `gorm:"not null;default:0" json:"productId"`
    SkuID         uint      `gorm:"not null;default:0" json:"skuId"`
    TotalAmount   float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"totalAmount"`
    PayAmount     float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"payAmount"`
    Quantity      uint      `gorm:"not null;default:0" json:"quantity"`
    Price         float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    CostPrice     float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"costPrice"`
    OriginalPrice float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"originalPrice"`
    CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}
func (SpOrderItem) TableName() string {
    return "sp_order_item"
}