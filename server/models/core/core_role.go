package core

import (
	"time"
	"encoding/json"
)
type CoreRole struct {
    ID          int64           `gorm:"primaryKey;autoIncrement" json:"id"`
    RoleName    string          `gorm:"size:128;not null;default:''" json:"role_name"`
    Permission  json.RawMessage `gorm:"type:json" json:"permission"`
    RoleStatus  int8            `gorm:"not null;default:0" json:"role_status"`
    Remark      string          `gorm:"size:200;not null;default:''" json:"remark"`
    CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
    UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
    DeletedTime    *time.Time      `json:"deleted_time"`
}

func (CoreRole) TableName() string {
    return "core_role"
}