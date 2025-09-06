package sp

import (
	"time"
)
type SpProductProperty struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID   uint      `gorm:"not null" json:"product_id"`
    Title       string    `gorm:"size:200;not null" json:"title"`
    Value       string    `gorm:"size:200;not null" json:"value"`
    SortNum     uint16    `gorm:"not null;default:0" json:"sort_num"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}
func (SpProductProperty) TableName() string {
    return "sp_product_property"
}