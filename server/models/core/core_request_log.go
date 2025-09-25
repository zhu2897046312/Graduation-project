package core

import (
	"server/models/common"
	"time"
)

type CoreRequestLog struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Source          int8           `gorm:"not null;default:0" json:"source"`
	Tag             string         `gorm:"size:128;not null;default:''" json:"tag"`
	Title           string         `gorm:"size:255;not null;default:''" json:"title"`
	RequestURL      string         `gorm:"size:255;not null;default:''" json:"requestUrl"`
	RequestMethod   string         `gorm:"size:10;not null;default:''" json:"requestMethod"`
	RequestParams   string         `gorm:"type:longtext" json:"requestParams"`
	IP              string         `gorm:"size:200;not null;default:''" json:"ip"`
	UseTime         int            `gorm:"not null;default:0" json:"useTime"`
	State           common.MyState `gorm:"not null;default:0" json:"state"`
	ResponseContent string         `gorm:"type:longtext" json:"responseContent"`
	Token           string         `gorm:"size:200" json:"token"`
	CreateTime      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"createTime"`
}

func (CoreRequestLog) TableName() string {
	return "core_request_log"
}
