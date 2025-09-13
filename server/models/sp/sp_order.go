package sp

import (
	"time"
	"server/models/common"
)

type SpOrder struct {
	ID               common.MyID       `gorm:"primaryKey;autoIncrement" json:"id"`
	Code             string     `gorm:"size:128;not null;default:''" json:"code"`
	UserID           common.MyID       `gorm:"not null;default:0" json:"user_id"`
	Nickname         string     `gorm:"size:100;not null" json:"nickname"`
	Email            string     `gorm:"size:100;not null" json:"email"`
	TotalAmount      float64    `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"total_amount"`
	PayAmount        float64    `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"pay_amount"`
	Freight          float64    `gorm:"type:decimal(10,2);unsigned;not null;default:0.00" json:"freight"`
	PayType          uint16     `gorm:"not null;default:0" json:"pay_type"`
	SourceType       int16      `gorm:"not null;default:0" json:"source_type"`
	State            uint8      `gorm:"not null;default:0" json:"state"`
	DisputeStatus    uint8      `gorm:"not null;default:0" json:"dispute_status"`
	PaymentTime      *time.Time `json:"payment_time"`
	DeliveryTime     *time.Time `json:"delivery_time"`
	ReceiveTime      *time.Time `json:"receive_time"`
	DeliveryCompany  string     `gorm:"size:128;not null;default:''" json:"delivery_company"`
	DeliverySn       string     `gorm:"size:200;not null;default:''" json:"delivery_sn"`
	Remark           string     `gorm:"size:230;not null;default:''" json:"remark"`
	VisitorQueryCode string     `gorm:"size:128;not null;default:''" json:"visitor_query_code"`
	CreatedTime      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_time"`
	UpdatedTime      time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_time"`
	DeletedTime      *time.Time `json:"deleted_time"`
}
type ListOrdersQueryParam struct {
	NikeName string `json:"nickname"`
	Email    string `json:"email"`
	Code     string `json:"code"`
	State    uint8  `json:"state"`
	Page     int    `json:"page_no"`
	PageSize int    `json:"page_size"`
}

func (SpOrder) TableName() string {
	return "sp_order"
}
