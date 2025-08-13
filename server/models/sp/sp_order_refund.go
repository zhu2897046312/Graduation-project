package sp

import (
	"time"
	"encoding/json"
)
type SpOrderRefund struct {
    ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    OrderID      uint           `gorm:"not null;default:0" json:"orderId"`
    RefundAmount float64        `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"refundAmount"`
    Status       uint8          `gorm:"not null;default:0" json:"status"`
    Reason       string         `gorm:"size:500" json:"reason"`
    RefundTime   *time.Time     `json:"refundTime"`
    Images       json.RawMessage `gorm:"type:json" json:"images"`
    CreatedTime  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
    RefundNo     string         `gorm:"size:64" json:"refundNo"`
}