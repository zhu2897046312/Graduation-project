package mp

import (
	"time"
)
type MpUserToken struct {
    ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID      uint      `gorm:"not null;default:0" json:"userId"`
    IP          string    `gorm:"size:200;not null;default:''" json:"ip"`
    Token       string    `gorm:"size:200;not null;default:''" json:"token"`
    UserAgent   string    `gorm:"size:250;not null;default:''" json:"userAgent"`
    ExpireTime  *time.Time `json:"expireTime"`
    CreatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}