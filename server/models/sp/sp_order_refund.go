package sp

import (
	"encoding/json"
	"server/models/common"
	"time"
)

type SpOrderRefund struct {
	ID           common.MyID     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      common.MyID     `gorm:"not null;default:0" json:"order_id"`
	RefundAmount float64         `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"refund_amount"`
	Status       uint8           `gorm:"not null;default:0" json:"status"`
	Reason       string          `gorm:"size:500" json:"reason"`
	RefundTime   time.Time       `json:"refund_time"`
	Images       json.RawMessage `gorm:"type:json" json:"images"`
	CreatedTime  time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime  time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime  *time.Time      `json:"deleted_time"`
	RefundNo     string          `gorm:"size:64" json:"refund_no"`
}
type ListSpOrderRefundQueyParam struct {
	RefundNo  string `json:"refund_no"`
	OrderCode string `json:"order_code"`
	Status    uint   `json:"status"`
	NikeName  string `json:"nikename"`
	Email     string `json:"email"`
}

func (SpOrderRefund) TableName() string {
	return "sp_order_refund"
}
