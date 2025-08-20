package cms

import (
	"time"
)
type CmsUserLikeHistory struct {
    ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    DocumentID  int64     `gorm:"not null;default:0" json:"documentId"`
    UserID      int64     `gorm:"not null;default:0" json:"userId"`
    IP          string    `gorm:"size:200;not null;default:''" json:"ip"`
    State       int8      `gorm:"not null;default:0" json:"state"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (CmsUserLikeHistory) TableName() string {
	return "cms_user_like_history"
}