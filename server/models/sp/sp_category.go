package sp

import (
	"time"
)
type SpCategory struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Pid           uint      `gorm:"not null;default:0" json:"pid"`
    Title         string    `gorm:"size:200;not null;default:''" json:"title"`
    Code          string    `gorm:"size:200;not null;default:''" json:"code"`
    State         uint8     `gorm:"not null;default:1" json:"state"`
    Icon          string    `gorm:"size:200;not null;default:''" json:"icon"`
    Picture       string    `gorm:"size:200;not null;default:''" json:"picture"`
    Description   string    `gorm:"size:200;not null;default:''" json:"description"`
    SortNum       uint16    `gorm:"not null;default:0" json:"sortNum"`
    SeoTitle      string    `gorm:"size:200;not null;default:''" json:"seoTitle"`
    SeoKeyword    string    `gorm:"size:200;not null;default:''" json:"seoKeyword"`
    SeoDescription string   `gorm:"size:200;not null;default:''" json:"seoDescription"`
    CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime   *time.Time `json:"deletedTime"`
}