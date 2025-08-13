package sp

import (
	"time"
)
type SpUserCart struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID        uint      `gorm:"not null;default:0" json:"userId"`
    Fingerprint   string    `gorm:"size:255;not null;default:''" json:"fingerprint"`
    Title         string    `gorm:"size:200;not null" json:"title"`
    SkuTitle      string    `gorm:"size:200;not null;default:''" json:"skuTitle"`
    SkuCode       string    `gorm:"size:200;not null;default:''" json:"skuCode"`
    Thumb         string    `gorm:"size:500;not null;default:''" json:"thumb"`
    ProductID     uint      `gorm:"not null;default:0" json:"productId"`
    SkuID         uint      `gorm:"not null;default:0" json:"skuId"`
    TotalAmount   float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"totalAmount"`
    PayAmount     float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"payAmount"`
    Quantity      uint      `gorm:"not null;default:0" json:"quantity"`
    Price         float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    OriginalPrice float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"originalPrice"`
    CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}