package shop

import (
	"time"
)
type ShopTag struct {
    ID         int        `gorm:"primaryKey;autoIncrement" json:"id"`
    Title      string     `gorm:"size:200;not null;default:''" json:"title"`
    Code       string     `gorm:"size:200;not null;default:''" json:"code"`
    Thumb      string     `gorm:"size:200;not null;default:''" json:"thumb"`
    State      int8       `gorm:"not null;default:0" json:"state"`
    ReadNum    int        `gorm:"not null;default:0" json:"readNum"`
    SortNum    int        `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime *time.Time `json:"deletedTime"`
    MatchWord   string     `gorm:"size:255" json:"matchWord"`
}
func (ShopTag) TableName() string {
    return "shop_tag"
}