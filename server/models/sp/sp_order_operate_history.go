package sp

import (
	"time"
)
type SpOrderOperateHistory struct {
    ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    OrderID      uint      `gorm:"not null;default:0" json:"orderId"`
    OperateUser  string    `gorm:"size:200;not null;default:''" json:"operateUser"`
    Remark       string    `gorm:"size:200;not null;default:''" json:"remark"`
    IP           string    `gorm:"size:200;not null;default:''" json:"ip"`
    CreatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}