package paypal

import (
	"gorm.io/gorm"
	"server/models/paypal"
	"server/repository/base"
)

type PaypalWebhookLogsRepository struct {
	*base.BaseRepository
}

func NewPaypalWebhookLogsRepository(DB *gorm.DB) *PaypalWebhookLogsRepository {
	return &PaypalWebhookLogsRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建PayPal Webhook日志
func (r *PaypalWebhookLogsRepository) Create(log *paypal.PaypalWebhookLogs) error {
	return r.DB.Create(log).Error
}

// 根据事件ID获取日志
func (r *PaypalWebhookLogsRepository) FindByEventID(eventID string) (*paypal.PaypalWebhookLogs, error) {
	var log paypal.PaypalWebhookLogs
	err := r.DB.Where("event_id = ?", eventID).First(&log).Error
	return &log, err
}

// 根据本地订单ID获取日志
func (r *PaypalWebhookLogsRepository) FindByLocalOrderID(localOrderID string) ([]paypal.PaypalWebhookLogs, error) {
	var logs []paypal.PaypalWebhookLogs
	err := r.DB.Where("local_order_id = ?", localOrderID).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 根据PayPal订单ID获取日志
func (r *PaypalWebhookLogsRepository) FindByPaypalOrderID(paypalOrderID string) ([]paypal.PaypalWebhookLogs, error) {
	var logs []paypal.PaypalWebhookLogs
	err := r.DB.Where("paypal_order_id = ?", paypalOrderID).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 根据事件类型获取日志
func (r *PaypalWebhookLogsRepository) FindByEventType(eventType string) ([]paypal.PaypalWebhookLogs, error) {
	var logs []paypal.PaypalWebhookLogs
	err := r.DB.Where("event_type = ?", eventType).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 更新处理结果
func (r *PaypalWebhookLogsRepository) UpdateProcessResult(id uint, result string) error {
	return r.DB.Model(&paypal.PaypalWebhookLogs{}).
		Where("id = ?", id).
		Update("process_result", result).Error
}