package core

import (
	"time"
	"encoding/json"
)
type CoreDept struct {
    ID             int64           `gorm:"primaryKey;autoIncrement" json:"id"`
    Pid            int64           `gorm:"not null;default:0" json:"pid"`
    DeptName       string          `gorm:"size:128;not null;default:''" json:"deptName"`
    ConnectName    string          `gorm:"size:128;not null;default:''" json:"connectName"`
    ConnectMobile  string          `gorm:"size:32;not null;default:''" json:"connectMobile"`
    ConnectAddress string          `gorm:"size:256;not null;default:''" json:"connectAddress"`
    LeaderName     string          `gorm:"size:128;not null;default:''" json:"leaderName"`
    Thumb          string          `gorm:"size:250;not null;default:''" json:"thumb"`
    Content        string          `gorm:"type:text" json:"content"`
    Organize       json.RawMessage `gorm:"type:json" json:"organize"`
    Level          int16           `gorm:"not null;default:0" json:"level"`
    SortNum        int             `gorm:"not null;default:0" json:"sortNum"`
    Remark         string          `gorm:"size:200;not null;default:''" json:"remark"`
    CreatedTime    time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime    time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime    *time.Time      `json:"deletedTime"`
}