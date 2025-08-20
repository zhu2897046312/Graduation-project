package sp

import (
	"time"
)
type SpProductContent struct {
    ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID      uint      `gorm:"not null" json:"productId"`
    Content        string    `gorm:"type:text;not null" json:"content"`
    SeoTitle       string    `gorm:"size:200;not null;default:''" json:"seoTitle"`
    SeoKeyword     string    `gorm:"size:200;not null;default:''" json:"seoKeyword"`
    SeoDescription string    `gorm:"size:200;not null;default:''" json:"seoDescription"`
    CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}
func (SpProductContent) TableName() string {
    return "sp_product_content"
}