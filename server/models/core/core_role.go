package core

import (
	"time"
	"encoding/json"
)
type CoreRole struct {
    ID          int64           `gorm:"primaryKey;autoIncrement" json:"id"`
    RoleName    string          `gorm:"size:128;not null;default:''" json:"roleName"`
    Permission  json.RawMessage `gorm:"type:json" json:"permission"`
    RoleStatus  int8            `gorm:"not null;default:0" json:"roleStatus"`
    Remark      string          `gorm:"size:200;not null;default:''" json:"remark"`
    CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}