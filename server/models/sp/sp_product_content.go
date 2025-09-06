package sp

import (
	"time"
)
type SpProductContent struct {
    ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID      uint      `gorm:"not null" json:"product_id"`
    Content        string    `gorm:"type:text;not null" json:"content"`
    SeoTitle       string    `gorm:"size:200;not null;default:''" json:"seo_title"`
    SeoKeyword     string    `gorm:"size:200;not null;default:''" json:"seo_keyword"`
    SeoDescription string    `gorm:"size:200;not null;default:''" json:"seo_description"`
    CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
    UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
}
func (SpProductContent) TableName() string {
    return "sp_product_content"
}