package sp

import (
	"time"
)
type SpOrder struct {
    ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Code           string    `gorm:"size:128;not null;default:''" json:"code"`
    UserID         uint      `gorm:"not null;default:0" json:"userId"`
    Nickname       string    `gorm:"size:100;not null" json:"nickname"`
    Email          string    `gorm:"size:100;not null" json:"email"`
    TotalAmount    float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"totalAmount"`
    PayAmount      float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"payAmount"`
    Freight        float64   `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"freight"`
    PayType        uint16    `gorm:"not null;default:0" json:"payType"`
    SourceType     int16     `gorm:"not null;default:0" json:"sourceType"`
    State          uint8     `gorm:"not null;default:0" json:"state"`
    DisputeStatus  uint8     `gorm:"not null;default:0" json:"disputeStatus"`
    PaymentTime    *time.Time `json:"paymentTime"`
    DeliveryTime   *time.Time `json:"deliveryTime"`
    ReceiveTime    *time.Time `json:"receiveTime"`
    DeliveryCompany string    `gorm:"size:128;not null;default:''" json:"deliveryCompany"`
    DeliverySn     string    `gorm:"size:200;not null;default:''" json:"deliverySn"`
    Remark         string    `gorm:"size:230;not null;default:''" json:"remark"`
    CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`
    UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`
}
func (SpOrder) TableName() string {
    return "sp_order"
}