package sp

import (
	"server/models/common"
	"time"
)

type SpUserAddress struct {
	ID            common.MyID `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        common.MyID `gorm:"not null;default:0" json:"user_id"`
	Title         string      `gorm:"size:200;not null" json:"title"`
	DefaultStatus int16      `json:"default_status"`
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

type SpUserAddressListParam struct {
	UserID   common.MyID `json:"user_id"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

func (SpUserAddress) TableName() string {
	return "sp_user_address"
}
