package core

import (
	"time"
	"encoding/json"
)
type CoreAdmin struct {
    ID           int64           `gorm:"primaryKey;autoIncrement" json:"id"`
    Nickname     string          `gorm:"size:128;not null;default:''" json:"nickname"`
    Account      string          `gorm:"size:128;not null;default:''" json:"account"`
    Pwd          string          `gorm:"size:130;not null;default:''" json:"pwd"`
    Mobile       string          `gorm:"size:32;not null;default:''" json:"mobile"`
    DeptID       int64           `gorm:"not null;default:0" json:"deptId"`
    AdminStatus  int8            `gorm:"not null;default:0" json:"adminStatus"`
    Permission   json.RawMessage `gorm:"type:json" json:"permission"`
    LastPwd      *time.Time      `json:"lastPwd"`
    CreatedTime  time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (CoreAdmin) TableName() string {
    return "core_admin"
}