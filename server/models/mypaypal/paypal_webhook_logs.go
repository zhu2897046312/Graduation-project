package mypaypal

import (
	"time"
	"encoding/json"
)
type PaypalWebhookLogs struct {
    ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
    LocalOrderID  string         `gorm:"size:64" json:"localOrderId"`
    PaypalOrderID string         `gorm:"size:64" json:"paypalOrderId"`
    EventID       string         `gorm:"size:64;not null;unique" json:"eventId"`
    EventType     string         `gorm:"size:128;not null" json:"eventType"`
    EventBody     json.RawMessage `gorm:"type:json;not null" json:"eventBody"`
    CreateTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"createTime"`
    ProcessResult string         `gorm:"size:255" json:"processResult"`
}
func (PaypalWebhookLogs) TableName() string {
    return "paypal_webhook_logs"
}