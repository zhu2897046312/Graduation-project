package sp

import (
	"time"
	"encoding/json"
)
type SpProduct struct {
    ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    CategoryID     uint           `gorm:"not null" json:"category_id"`
    Title          string         `gorm:"size:200;not null" json:"title"`
    State          uint8          `gorm:"not null;default:1" json:"state"`
    Price          float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"price"`
    OriginalPrice  float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"original_price"`
    CostPrice      float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"cost_price"`
    Stock          uint           `gorm:"not null;default:0" json:"stock"`
    OpenSku        uint8          `gorm:"not null;default:0" json:"open_sku"`
    Picture        string         `gorm:"size:200;not null" json:"picture"`
    PictureGallery json.RawMessage `gorm:"type:json" json:"picture_gallery"`
    Description    string         `gorm:"size:200;not null" json:"description"`
    SoldNum        uint16         `json:"sold_num"`
    Version        uint           `gorm:"not null;default:0" json:"version"`
    SortNum        uint16         `gorm:"not null;default:0" json:"sort_num"`
    Hot            uint8          `gorm:"not null;default:0" json:"hot"`
    PutawayTime    *time.Time     `json:"putaway_time"`
    CreatedTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
    UpdatedTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
    DeletedTime    *time.Time     `json:"deleted_time"`
    DetailURL      string         `gorm:"size:512" json:"detail_url"`
    PriceLocked    bool           `gorm:"default:0" json:"price_locked"`
}
// ProductQueryParams 商品查询参数结构体
type ProductQueryParams struct {
	Page       int
	PageSize   int
	CategoryID uint
	State      int
	Title      string
	Hot        *bool
	SortBy     string // 排序字段
	SortOrder  string // 排序方向: asc, desc
}
func (SpProduct) TableName() string {
    return "sp_product"
}