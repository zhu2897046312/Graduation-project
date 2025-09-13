package mp

import (
	"server/models/common"
	"time"
)

type MpPayConfig struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string      `gorm:"size:255;not null;default:''" json:"title"`
	Photo       string      `gorm:"size:255;not null;default:''" json:"photo"`
	Code        string      `gorm:"size:255;not null;default:''" json:"code"`
	State       int8        `gorm:"not null;default:0" json:"state"`
	SortNum     int         `gorm:"not null;default:0" json:"sortNum"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (MpPayConfig) TableName() string {
	return "mp_pay_config"
}
