package shop

import (
	"server/models/common"
	"time"
)

type ShopTag struct {
	ID          common.MyID      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string           `gorm:"size:200;not null;default:''" json:"title"`
	Code        string           `gorm:"size:200;not null;default:''" json:"code"`
	Thumb       string           `gorm:"size:200;not null;default:''" json:"thumb"`
	State       common.MyState   `gorm:"not null;default:0" json:"state"`
	ReadNum     common.MyNumber  `gorm:"not null;default:0" json:"read_num"`
	SortNum     common.MySortNum `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time       `json:"deleted_time"`
	MatchWord   string           `gorm:"size:255" json:"match_word"`
}

type TagQueryParams struct {
	Page      int
	PageSize  int
	State     common.MyState
	Title     string
	Code      string
	SortBy    string // 排序字段
	SortOrder string // 排序方向: asc, desc
}

func (ShopTag) TableName() string {
	return "shop_tag"
}
