package sp

import (
	"time"
)
type SpSkuIndex struct {
    ID                     uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID              uint      `gorm:"not null" json:"productId"`
    SkuID                  uint      `gorm:"not null" json:"skuId"`
    ProdAttributesID       uint      `gorm:"not null" json:"prodAttributesId"`
    ProdAttributesValueID  uint      `gorm:"not null" json:"prodAttributesValueId"`
    CreatedTime            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}