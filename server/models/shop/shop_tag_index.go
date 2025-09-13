package shop

import (
	"server/models/common"
	"time"
)

type ShopTagIndex struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID   common.MyID `gorm:"not null;default:0" json:"productId"`
	TagID       common.MyID `gorm:"not null;default:0" json:"tagId"`
	SortNum     int         `gorm:"not null;default:0" json:"sortNum"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (ShopTagIndex) TableName() string {
	return "shop_tag_index"
}
