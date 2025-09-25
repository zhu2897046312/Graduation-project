package sp

import (
	"server/models/common"
	"time"
)

type SpSku struct {
	ID            common.MyID     `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     common.MyID     `gorm:"not null" json:"product_id"`
	SkuCode       string          `gorm:"size:200;not null" json:"sku_code"`
	Title         string          `gorm:"size:200;not null" json:"title"`
	Price         float64         `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
	OriginalPrice float64         `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"original_price"`
	CostPrice     float64         `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"cost_price"`
	Stock         common.MyNumber `gorm:"not null;default:0" json:"stock"`
	DefaultShow   common.MyState  `gorm:"not null;default:0" json:"default_show"`
	State         common.MyState  `gorm:"not null;default:1" json:"state"`
	Version       uint            `gorm:"not null;default:0" json:"version"`
	CreatedTime   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime   time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime   *time.Time      `json:"deleted_time"`
}

func (SpSku) TableName() string {
	return "sp_sku"
}
