package shop

import (
	"time"
)
type ShopTagIndex struct {
    ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
    ProductID   int       `gorm:"not null;default:0" json:"productId"`
    TagID       int       `gorm:"not null;default:0" json:"tagId"`
    SortNum     int       `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}