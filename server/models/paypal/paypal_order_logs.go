package paypal

import (
	"time"
)
type PaypalOrderLogs struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    LocalOrderID  string    `gorm:"size:64;not null" json:"localOrderId"`
    PaypalOrderID string    `gorm:"size:64;not null;unique" json:"paypalOrderId"`
    CreateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createTime"`
}