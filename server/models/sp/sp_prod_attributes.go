package sp

import (
	"server/models/common"
	"time"
)

type SpProdAttributes struct {
	ID          common.MyID      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string           `gorm:"size:200;not null" json:"title"`
	SortNum     common.MySortNum `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time       `json:"deleted_time"`
}

func (SpProdAttributes) TableName() string {
	return "sp_prod_attributes"
}
