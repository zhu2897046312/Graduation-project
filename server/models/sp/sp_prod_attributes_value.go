package sp

import (
	"time"
)
type SpProdAttributesValue struct {
    ID                uint       `gorm:"primaryKey;autoIncrement" json:"id"`
    ProdAttributesID  uint       `gorm:"not null" json:"prodAttributesId"`
    Title             string     `gorm:"size:200;not null" json:"title"`
    SortNum           uint16     `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime       *time.Time `json:"deletedTime"`
}