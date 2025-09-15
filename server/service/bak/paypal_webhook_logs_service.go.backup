package service

import (
	"errors"
	"server/models/paypal"
)

type PaypalWebhookLogsService struct {
	*Service
}

func NewPaypalWebhookLogsService(base *Service) *PaypalWebhookLogsService {
	return &PaypalWebhookLogsService{Service: base}
}

// CreateWebhookLog 创建PayPal Webhook日志
func (s *PaypalWebhookLogsService) CreateWebhookLog(log *paypal.PaypalWebhookLogs) error {
	if log.EventID == "" {
		return errors.New("事件ID不能为空")
	}
	if log.EventType == "" {
		return errors.New("事件类型不能为空")
	}
	
	return s.repoFactory.GetPaypalWebhookLogsRepository().Create(log)
}

// GetLogByEventID 根据事件ID获取日志
func (s *PaypalWebhookLogsService) GetLogByEventID(eventID string) (*paypal.PaypalWebhookLogs, error) {
	if eventID == "" {
		return nil, errors.New("事件ID不能为空")
	}
	return s.repoFactory.GetPaypalWebhookLogsRepository().FindByEventID(eventID)
}

// GetLogsByLocalOrderID 根据本地订单ID获取日志
func (s *PaypalWebhookLogsService) GetLogsByLocalOrderID(localOrderID string) ([]paypal.PaypalWebhookLogs, error) {
	if localOrderID == "" {
		return nil, errors.New("本地订单ID不能为空")
	}
	return s.repoFactory.GetPaypalWebhookLogsRepository().FindByLocalOrderID(localOrderID)
}

// GetLogsByPaypalOrderID 根据PayPal订单ID获取日志
func (s *PaypalWebhookLogsService) GetLogsByPaypalOrderID(paypalOrderID string) ([]paypal.PaypalWebhookLogs, error) {
	if paypalOrderID == "" {
		return nil, errors.New("PayPal订单ID不能为空")
	}
	return s.repoFactory.GetPaypalWebhookLogsRepository().FindByPaypalOrderID(paypalOrderID)
}

// GetLogsByEventType 根据事件类型获取日志
func (s *PaypalWebhookLogsService) GetLogsByEventType(eventType string) ([]paypal.PaypalWebhookLogs, error) {
	if eventType == "" {
		return nil, errors.New("事件类型不能为空")
	}
	return s.repoFactory.GetPaypalWebhookLogsRepository().FindByEventType(eventType)
}

// UpdateProcessResult 更新处理结果
func (s *PaypalWebhookLogsService) UpdateProcessResult(id uint, result string) error {
	if id == 0 {
		return errors.New("日志ID不能为空")
	}
	if result == "" {
		return errors.New("处理结果不能为空")
	}
	return s.repoFactory.GetPaypalWebhookLogsRepository().UpdateProcessResult(id, result)
}