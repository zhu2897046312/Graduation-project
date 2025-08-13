package cms

import (
	"time"
)
type CmsRecommend struct {
    ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
    Title       string     `gorm:"size:200;not null;default:''" json:"title"`
    Code        string     `gorm:"size:200;not null;default:''" json:"code"`
    Thumb       string     `gorm:"size:200;not null;default:''" json:"thumb"`
    Description string     `gorm:"size:200;not null;default:''" json:"description"`
    State       int8       `gorm:"not null;default:0" json:"state"`
    MoreLink    string     `gorm:"size:200;not null;default:''" json:"moreLink"`
    CreatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    DeletedTime *time.Time `json:"deletedTime"`
}