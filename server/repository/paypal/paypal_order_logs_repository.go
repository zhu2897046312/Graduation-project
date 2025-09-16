package paypal

import (
	"gorm.io/gorm"
	"server/models/mypaypal"
	"server/repository/base"
)

type PaypalOrderLogsRepository struct {
	*base.BaseRepository
}

func NewPaypalOrderLogsRepository(DB *gorm.DB) *PaypalOrderLogsRepository {
	return &PaypalOrderLogsRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建PayPal订单日志
func (r *PaypalOrderLogsRepository) Create(log *mypaypal.PaypalOrderLogs) error {
	return r.DB.Create(log).Error
}

// 根据本地订单ID获取日志
func (r *PaypalOrderLogsRepository) FindByLocalOrderID(localOrderID string) ([]mypaypal.PaypalOrderLogs, error) {
	var logs []mypaypal.PaypalOrderLogs
	err := r.DB.Where("local_order_id = ?", localOrderID).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 根据PayPal订单ID获取日志
func (r *PaypalOrderLogsRepository) FindByPaypalOrderID(paypalOrderID string) (*mypaypal.PaypalOrderLogs, error) {
	var log mypaypal.PaypalOrderLogs
	err := r.DB.Where("paypal_order_id = ?", paypalOrderID).First(&log).Error
	return &log, err
}

// 获取所有日志
func (r *PaypalOrderLogsRepository) FindAll() ([]mypaypal.PaypalOrderLogs, error) {
	var logs []mypaypal.PaypalOrderLogs
	err := r.DB.Order("create_time DESC").Find(&logs).Error
	return logs, err
}