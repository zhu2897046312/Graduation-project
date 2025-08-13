package core

import (
	"time"
	"encoding/json"
)
type CorePermission struct {
    ID          int64           `gorm:"primaryKey;autoIncrement" json:"id"`
    Title       string          `gorm:"size:128;not null;default:''" json:"title"`
    Code        string          `gorm:"size:128;not null;default:''" json:"code"`
    Pid         int64           `gorm:"not null;default:0" json:"pid"`
    Urls        json.RawMessage `gorm:"type:json" json:"urls"`
    Remark      string          `gorm:"size:200;not null;default:''" json:"remark"`
    CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}