package core

import (
	"encoding/json"
	"server/models/common"
	"time"
)

type CoreAdmin struct {
	ID          common.MyID     `gorm:"primaryKey;autoIncrement" json:"id"`
	DeptID      common.MyID     `gorm:"not null;default:0" json:"dept_id"`
	Nickname    string          `gorm:"size:128;not null;default:''" json:"nickname"`
	Account     string          `gorm:"size:128;not null;default:''" json:"account"`
	Pwd         string          `gorm:"size:130;not null;default:''" json:"pwd"`
	Mobile      string          `gorm:"size:32;not null;default:''" json:"mobile"`
	AdminStatus int8            `gorm:"not null;default:0" json:"admin_status"`
	Permission  json.RawMessage `gorm:"type:json" json:"permission"`
	LastPwd     *time.Time      `json:"last_pwd"`
	CreatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time      `json:"deleted_time"`
}
type CoreAdminQueryParam struct {
	Nickname    string
	Account     string
	AdminStatus int8
	Page        int
	PageSize    int
}

func (CoreAdmin) TableName() string {
	return "core_admin"
}
