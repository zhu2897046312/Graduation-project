package cms

import (
	"server/models/common"
	"time"
)

type CmsRecommendIndex struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	RecommendID common.MyID `gorm:"not null;default:0" json:"recommend_id"`
	ProductID   common.MyID `gorm:"not null;default:0" json:"product_id"`
	DocumentID  common.MyID `gorm:"not null;default:0" json:"document_id"`
	Title       string      `gorm:"size:200;not null;default:''" json:"title"`
	Thumb       string      `gorm:"size:200;not null;default:''" json:"thumb"`
	Link        string      `gorm:"size:200;not null;default:''" json:"link"`
	State       int8        `gorm:"not null;default:0" json:"state"`
	SortNum     int         `gorm:"not null;default:0" json:"sort_num"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime *time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_time"`
}
type RecommendIndexQueryParams struct {
	Title       string      `json:"title"`
	RecommendID common.MyID `json:"recommend_id"`
	Page        int         `json:"page_no"`
	PageSize    int         `json:"page_size"`
}

func (CmsRecommendIndex) TableName() string {
	return "cms_recommend_index"
}
