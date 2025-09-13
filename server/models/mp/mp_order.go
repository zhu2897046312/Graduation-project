package mp

import (
	"server/models/common"
	"time"
)

type MpOrder struct {
	ID           common.MyID `gorm:"primaryKey;size:128" json:"id"`
	ProductID    common.MyID `gorm:"not null;default:0" json:"productId"`
	UserID       common.MyID `gorm:"not null;default:0" json:"userId"`
	PayConfigID  common.MyID `gorm:"not null;default:0" json:"payConfigId"`
	ThirdID      common.MyID `gorm:"size:64;not null;default:''" json:"thirdId"`
	PayPrice     int         `gorm:"not null;default:0" json:"payPrice"`
	State        int8        `gorm:"not null;default:0" json:"state"`
	ErrMsg       string      `gorm:"size:128;not null;default:''" json:"errMsg"`
	Remark       string      `gorm:"size:250;not null;default:''" json:"remark"`
	FailTime     *time.Time  `json:"failTime"`
	CloseTime    *time.Time  `json:"closeTime"`
	CompleteTime *time.Time  `json:"completeTime"`
	CreatedTime  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (MpOrder) TableName() string {
	return "mp_order"
}
