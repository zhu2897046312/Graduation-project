package core

import (
	"encoding/json"
	"server/models/common"
	"time"
)

type CoreRole struct {
	ID          common.MyID     `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName    string          `gorm:"size:128;not null;default:''" json:"role_name"`
	Permission  json.RawMessage `gorm:"type:json" json:"permission"`
	RoleStatus  common.MyState  `gorm:"not null;default:0" json:"role_status"`
	Remark      string          `gorm:"size:200;not null;default:''" json:"remark"`
	CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time      `json:"deleted_time"`
}

func (CoreRole) TableName() string {
	return "core_role"
}
