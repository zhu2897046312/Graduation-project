package sp

import (
	"time"
)
type SpSku struct {
    ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID    uint       `gorm:"not null" json:"productId"`
    SkuCode      string     `gorm:"size:200;not null" json:"skuCode"`
    Title        string     `gorm:"size:200;not null" json:"title"`
    Price        float64    `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    OriginalPrice float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"originalPrice"`
    CostPrice    float64    `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"costPrice"`
    Stock        uint       `gorm:"not null;default:0" json:"stock"`
    DefaultShow  uint8      `gorm:"not null;default:0" json:"defaultShow"`
    State        uint8      `gorm:"not null;default:1" json:"state"`
    Version      uint       `gorm:"not null;default:0" json:"version"`
    CreatedTime  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime  *time.Time `json:"deletedTime"`
}
func (SpSku) TableName() string {
    return "sp_sku"
}