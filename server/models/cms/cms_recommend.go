package cms

import (
	"server/models/common"
	"time"
)

type CmsRecommend struct {
	ID          common.MyID    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"size:200;not null;default:''" json:"title"`
	Code        string         `gorm:"size:200;not null;default:''" json:"code"`
	Thumb       string         `gorm:"size:200;not null;default:''" json:"thumb"`
	Description string         `gorm:"size:200;not null;default:''" json:"description"`
	State       common.MyState `gorm:"not null;default:0" json:"state"`
	MoreLink    string         `gorm:"size:200;not null;default:''" json:"more_link"`
	CreatedTime time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time     `json:"deletedTime"`
}
type RecommendQueryParams struct {
	Page     int `json:"page_no"`
	PageSize int `json:"page_size"`
}

func (CmsRecommend) TableName() string {
	return "cms_recommend"
}
