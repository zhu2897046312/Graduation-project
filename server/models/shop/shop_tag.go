package shop

import (
	"time"
)

type ShopTag struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string     `gorm:"size:200;not null;default:''" json:"title"`
	Code        string     `gorm:"size:200;not null;default:''" json:"code"`
	Thumb       string     `gorm:"size:200;not null;default:''" json:"thumb"`
	State       int8       `gorm:"not null;default:0" json:"state"`
	ReadNum     int        `gorm:"not null;default:0" json:"read_num"`
	SortNum     int        `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time `json:"deleted_time"`
	MatchWord   string     `gorm:"size:255" json:"match_word"`
}

type TagQueryParams struct {
	Page      int
	PageSize  int
	State     int
	Title     string
	Code      string
	SortBy    string // 排序字段
	SortOrder string // 排序方向: asc, desc
}

func (ShopTag) TableName() string {
	return "shop_tag"
}
