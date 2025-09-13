package shop

import (
	"server/models/common"
	"time"
)

type ShopTagMate struct {
	ID             common.MyID `gorm:"primaryKey" json:"id"`
	Content        string      `gorm:"type:text;not null" json:"content"`
	SeoTitle       string      `gorm:"size:200;not null;default:''" json:"seo_title"`
	SeoKeyword     string      `gorm:"size:200;not null;default:''" json:"seo_keyword"`
	SeoDescription string      `gorm:"size:200;not null;default:''" json:"seo_description"`
	CreatedTime    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
	DeletedTime    *time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"deletedTime"`
}

func (ShopTagMate) TableName() string {
	return "shop_tag_mate"
}
