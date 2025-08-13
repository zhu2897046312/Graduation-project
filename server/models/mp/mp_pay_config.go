package mp

import (
	"time"
)
type MpPayConfig struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Title       string    `gorm:"size:255;not null;default:''" json:"title"`
    Photo       string    `gorm:"size:255;not null;default:''" json:"photo"`
    Code        string    `gorm:"size:255;not null;default:''" json:"code"`
    State       int8      `gorm:"not null;default:0" json:"state"`
    SortNum     int       `gorm:"not null;default:0" json:"sortNum"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}