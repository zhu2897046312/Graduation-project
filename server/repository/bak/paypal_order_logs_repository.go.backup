package repository

import (
	"gorm.io/gorm"
	"server/models/paypal"
)

type PaypalOrderLogsRepository struct {
	*BaseRepository
}

func NewPaypalOrderLogsRepository(db *gorm.DB) *PaypalOrderLogsRepository {
	return &PaypalOrderLogsRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建PayPal订单日志
func (r *PaypalOrderLogsRepository) Create(log *paypal.PaypalOrderLogs) error {
	return r.db.Create(log).Error
}

// 根据本地订单ID获取日志
func (r *PaypalOrderLogsRepository) FindByLocalOrderID(localOrderID string) ([]paypal.PaypalOrderLogs, error) {
	var logs []paypal.PaypalOrderLogs
	err := r.db.Where("local_order_id = ?", localOrderID).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 根据PayPal订单ID获取日志
func (r *PaypalOrderLogsRepository) FindByPaypalOrderID(paypalOrderID string) (*paypal.PaypalOrderLogs, error) {
	var log paypal.PaypalOrderLogs
	err := r.db.Where("paypal_order_id = ?", paypalOrderID).First(&log).Error
	return &log, err
}

// 获取所有日志
func (r *PaypalOrderLogsRepository) FindAll() ([]paypal.PaypalOrderLogs, error) {
	var logs []paypal.PaypalOrderLogs
	err := r.db.Order("create_time DESC").Find(&logs).Error
	return logs, err
}