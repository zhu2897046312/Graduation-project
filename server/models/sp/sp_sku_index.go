package sp

import (
	"time"
)
type SpSkuIndex struct {
    ID                     uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID              uint      `gorm:"not null" json:"product_id"`
    SkuID                  uint      `gorm:"not null" json:"sku_id"`
    ProdAttributesID       uint      `gorm:"not null" json:"prod_attributes_id"`
    ProdAttributesValueID  uint      `gorm:"not null" json:"prod_attributes_value_id"`
    CreatedTime            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}
func (SpSkuIndex) TableName() string {
    return "sp_sku_index"
}