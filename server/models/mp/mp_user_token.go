package mp

import (
	"server/models/common"
	"time"
)

type MpUserToken struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      common.MyID `gorm:"not null;default:0" json:"userId"`
	IP          string      `gorm:"size:200;not null;default:''" json:"ip"`
	Token       string      `gorm:"size:200;not null;default:''" json:"token"`
	UserAgent   string      `gorm:"size:250;not null;default:''" json:"userAgent"`
	ExpireTime  *time.Time  `json:"expireTime"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (MpUserToken) TableName() string {
	return "mp_user_token"
}
