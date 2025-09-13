package sp

import (
	"server/models/common"
	"time"
)

type SpOrderOperateHistory struct {
	ID          common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     common.MyID `gorm:"not null;default:0" json:"order_id"`
	OperateUser string      `gorm:"size:200;not null;default:''" json:"operate_user"`
	Remark      string      `gorm:"size:200;not null;default:''" json:"remark"`
	IP          string      `gorm:"size:200;not null;default:''" json:"ip"`
	CreatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
	UpdatedTime time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}

func (SpOrderOperateHistory) TableName() string {
	return "sp_order_operate_history"
}
