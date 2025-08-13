package mp

import (
	"time"
)
type MpUser struct {
    ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
    Nickname    string     `gorm:"size:255;not null;default:''" json:"nickname"`
    Avatar      string     `gorm:"size:255;not null;default:''" json:"avatar"`
    Email       string     `gorm:"size:250;not null;default:'';unique" json:"email"`
    EmailVerify int        `gorm:"not null;default:0" json:"emailVerify"`
    Password    string     `gorm:"size:250;not null;default:''" json:"password"`
    UserType    int8       `gorm:"not null;default:0" json:"userType"`
    UserStatus  int8       `gorm:"not null;default:0" json:"userStatus"`
    CreatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    Token       string     `gorm:"size:100" json:"token"`
}