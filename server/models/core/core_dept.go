package core

import (
	"encoding/json"
	"time"
)

type CoreDept struct {
	ID             int64           `gorm:"primaryKey;autoIncrement" json:"id"`
	Pid            int64           `gorm:"not null;default:0" json:"pid"`
	DeptName       string          `gorm:"size:128;not null;default:''" json:"dept_name"`
	ConnectName    string          `gorm:"size:128;not null;default:''" json:"connect_name"`
	ConnectMobile  string          `gorm:"size:32;not null;default:''" json:"connect_mobile"`
	ConnectAddress string          `gorm:"size:256;not null;default:''" json:"connect_address"`
	LeaderName     string          `gorm:"size:128;not null;default:''" json:"leader_name"`
	Thumb          string          `gorm:"size:250;not null;default:''" json:"thumb"`
	Content        string          `gorm:"type:text" json:"content"`
	Organize       json.RawMessage `gorm:"type:json" json:"organize"`
	Level          int16           `gorm:"not null;default:0" json:"level"`
	SortNum        int             `gorm:"not null;default:0" json:"sort_num"`
	Remark         string          `gorm:"size:200;not null;default:''" json:"remark"`
	CreatedTime    time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime    time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime    *time.Time      `json:"deleted_time"`
}

func (CoreDept) TableName() string {
	return "core_dept"
}
