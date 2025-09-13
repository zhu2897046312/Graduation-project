package core

import (
	"encoding/json"
	"server/models/common"
	"time"
)

type CorePermission struct {
	ID          common.MyID     `gorm:"primaryKey;autoIncrement" json:"id"`
	Pid         common.MyID     `gorm:"not null;default:0" json:"pid"`
	Title       string          `gorm:"size:128;not null;default:''" json:"title"`
	Code        string          `gorm:"size:128;not null;default:''" json:"code"`
	Urls        json.RawMessage `gorm:"type:json" json:"urls"`
	Remark      string          `gorm:"size:200;not null;default:''" json:"remark"`
	CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (CorePermission) TableName() string {
	return "core_permission"
}
