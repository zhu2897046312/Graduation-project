package sp

import (
	"time"
)

type SpProdAttributesValue struct {
	ID               uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProdAttributesID uint       `gorm:"not null" json:"prod_attributes_id"`
	Title            string     `gorm:"size:200;not null" json:"title"`
	SortNum          uint16     `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime      *time.Time `json:"deleted_tcdime"`
}

func (SpProdAttributesValue) TableName() string {
	return "sp_prod_attributes_value"
}

type SpProdAttributesQueryParams struct {
	Page             int
	PageSize         int
	ProdAttributesID uint
}
