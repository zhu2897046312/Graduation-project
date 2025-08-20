package shop

import (
	"time"
)
type ShopTagMate struct {
    ID             int       `gorm:"primaryKey" json:"id"`
    Content        string    `gorm:"type:text;not null" json:"content"`
    SeoTitle       string    `gorm:"size:200;not null;default:''" json:"seoTitle"`
    SeoKeyword     string    `gorm:"size:200;not null;default:''" json:"seoKeyword"`
    SeoDescription string    `gorm:"size:200;not null;default:''" json:"seoDescription"`
    CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}
func (ShopTagMate) TableName() string {
    return "shop_tag_mate"
}