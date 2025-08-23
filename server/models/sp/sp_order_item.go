package sp

import (
	"time"
)

type SpOrderItem struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title         string    `gorm:"size:200;not null;default:''" json:"title"`
    SkuTitle      string    `gorm:"size:200;not null;default:''" json:"sku_title"`
    SkuCode       string    `gorm:"size:200;not null;default:''" json:"sku_code"`
    Thumb         string    `gorm:"size:500;not null;default:''" json:"thumb"`
    OrderID       uint      `gorm:"not null;default:0" json:"order_id"`
    ProductID     uint      `gorm:"not null;default:0" json:"product_id"`
    SkuID         uint      `gorm:"not null;default:0" json:"sku_id"`
    TotalAmount   float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"total_amount"`
    PayAmount     float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"pay_amount"`
    Quantity      uint      `gorm:"not null;default:0" json:"quantity"`
    Price         float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    CostPrice     float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"cost_price"`
    OriginalPrice float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"original_price"`
    CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
    UpdatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}
func (SpOrderItem) TableName() string {
    return "sp_order_item"
}