package core

import (
	"server/models/common"
	"time"
)

type CoreConfig struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigKey   string      `gorm:"size:255;not null;default:''" json:"config_key"`
	ConfigValue string      `gorm:"type:text" json:"config_value"`
	Remark      string      `gorm:"size:255;not null;default:''" json:"remark"`
	SortNum     int         `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}

func (CoreConfig) TableName() string {
	return "core_config"
}
