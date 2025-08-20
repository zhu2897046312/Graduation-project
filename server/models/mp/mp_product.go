package mp

import (
	"time"
)
type MpProduct struct {
    ID            int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Title         string    `gorm:"size:255;not null;default:''" json:"title"`
    Price         int       `gorm:"not null;default:0" json:"price"`
    ProductType   int8      `gorm:"not null;default:0" json:"productType"`
    TerminalType  int8      `gorm:"not null;default:1" json:"terminalType"`
    ProductCode   string    `gorm:"size:255;not null;default:''" json:"productCode"`
    ProductConfig string    `gorm:"type:text" json:"productConfig"`
    Code          string    `gorm:"size:64;not null;default:''" json:"code"`
    SortNum       int       `gorm:"not null;default:0" json:"sortNum"`
    State         int       `gorm:"not null;default:0" json:"state"`
    CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    ShowType      int8      `gorm:"not null;default:0" json:"showType"`
}
func (MpProduct) TableName() string {
    return "mp_product"
}