package mp

import (
	"time"
)
type MpResetPwdTokens struct {
    ID             int       `gorm:"primaryKey;autoIncrement" json:"id"`
    Email          string    `gorm:"size:255;not null;unique" json:"email"`
    Token          string    `gorm:"size:255;not null" json:"token"`
    ExpirationTime time.Time `gorm:"not null" json:"expirationTime"`
    CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    Count          int       `gorm:"not null;default:1" json:"count"`
}
func (MpResetPwdTokens) TableName() string {
    return "mp_reset_pwd_tokens"
}