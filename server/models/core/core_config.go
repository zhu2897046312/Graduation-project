package core

import (
	"time"
)
type CoreConfig struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    ConfigKey   string    `gorm:"size:255;not null;default:''" json:"configKey"`
    ConfigValue string    `gorm:"type:text" json:"configValue"`
    Remark      string    `gorm:"size:255;not null;default:''" json:"remark"`
    SortNum     int       `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (CoreConfig) TableName() string {
    return "core_config"
}