package sp

import (
	"server/models/common"
	"time"
)

type SpSkuIndex struct {
	ID                    common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID             common.MyID `gorm:"not null" json:"product_id"`
	SkuID                 common.MyID `gorm:"not null" json:"sku_id"`
	ProdAttributesID      common.MyID `gorm:"not null" json:"prod_attributes_id"`
	ProdAttributesValueID common.MyID `gorm:"not null" json:"prod_attributes_value_id"`
	CreatedTime           time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime           time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (SpSkuIndex) TableName() string {
	return "sp_sku_index"
}
