package mp

import (
	"time"
)
type MpOrder struct {
    ID           string     `gorm:"primaryKey;size:128" json:"id"`
    ProductID    int64      `gorm:"not null;default:0" json:"productId"`
    UserID       int64      `gorm:"not null;default:0" json:"userId"`
    PayPrice     int        `gorm:"not null;default:0" json:"payPrice"`
    State        int8       `gorm:"not null;default:0" json:"state"`
    PayConfigID  int64      `gorm:"not null;default:0" json:"payConfigId"`
    ThirdID      string     `gorm:"size:64;not null;default:''" json:"thirdId"`
    ErrMsg       string     `gorm:"size:128;not null;default:''" json:"errMsg"`
    Remark       string     `gorm:"size:250;not null;default:''" json:"remark"`
    FailTime     *time.Time `json:"failTime"`
    CloseTime    *time.Time `json:"closeTime"`
    CompleteTime *time.Time `json:"completeTime"`
    CreatedTime  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}