package sp

import (
	"server/models/common"
	"time"
)

type SpOrderReceiveAddress struct {
	ID            common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       common.MyID `gorm:"not null;default:0" json:"order_id"`
	FirstName     string      `gorm:"size:64;not null;default:''" json:"first_name"`
	LastName      string      `gorm:"size:64;not null;default:''" json:"last_name"`
	Email         string      `gorm:"size:128;not null;default:''" json:"email"`
	Phone         string      `gorm:"size:64;not null;default:''" json:"phone"`
	Province      string      `gorm:"size:64;not null;default:''" json:"province"`
	City          string      `gorm:"size:64;not null;default:''" json:"city"`
	Region        string      `gorm:"size:64;not null;default:''" json:"region"`
	DetailAddress string      `gorm:"size:200;not null;default:''" json:"detail_address"`
	CreatedTime   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	Country       string      `gorm:"size:64;not null;default:''" json:"country"`
	PostalCode    string      `gorm:"size:64;not null;default:''" json:"postal_code"`
}

func (SpOrderReceiveAddress) TableName() string {
	return "sp_order_receive_address"
}
