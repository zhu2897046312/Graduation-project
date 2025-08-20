package sp

import (
	"time"
	"encoding/json"
)
type SpProduct struct {
    ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    CategoryID     uint           `gorm:"not null" json:"categoryId"`
    Title          string         `gorm:"size:200;not null" json:"title"`
    State          uint8          `gorm:"not null;default:1" json:"state"`
    Price          float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    OriginalPrice  float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"originalPrice"`
    CostPrice      float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"costPrice"`
    Stock          uint           `gorm:"not null;default:0" json:"stock"`
    OpenSku        uint8          `gorm:"not null;default:0" json:"openSku"`
    Picture        string         `gorm:"size:200;not null" json:"picture"`
    PictureGallery json.RawMessage `gorm:"type:json" json:"pictureGallery"`
    Description    string         `gorm:"size:200;not null" json:"description"`
    SoldNum        uint16         `json:"soldNum"`
    Version        uint           `gorm:"not null;default:0" json:"version"`
    SortNum        uint16         `gorm:"not null;default:0" json:"sortNum"`
    Hot            uint8          `gorm:"not null;default:0" json:"hot"`
    PutawayTime    *time.Time     `json:"putawayTime"`
    CreatedTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime    *time.Time     `json:"deletedTime"`
    DetailURL      string         `gorm:"size:512" json:"detailUrl"`
    PriceLocked    bool           `gorm:"default:0" json:"priceLocked"`
}
func (SpProduct) TableName() string {
    return "sp_product"
}