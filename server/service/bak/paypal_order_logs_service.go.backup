package service

import (
	"errors"
	"server/models/paypal"
)

type PaypalOrderLogsService struct {
	*Service
}

func NewPaypalOrderLogsService(base *Service) *PaypalOrderLogsService {
	return &PaypalOrderLogsService{Service: base}
}

// CreateOrderLog 创建PayPal订单日志
func (s *PaypalOrderLogsService) CreateOrderLog(log *paypal.PaypalOrderLogs) error {
	if log.LocalOrderID == "" {
		return errors.New("本地订单ID不能为空")
	}
	if log.PaypalOrderID == "" {
		return errors.New("PayPal订单ID不能为空")
	}
	
	return s.repoFactory.GetPaypalOrderLogsRepository().Create(log)
}

// GetLogsByLocalOrderID 根据本地订单ID获取日志
func (s *PaypalOrderLogsService) GetLogsByLocalOrderID(localOrderID string) ([]paypal.PaypalOrderLogs, error) {
	if localOrderID == "" {
		return nil, errors.New("本地订单ID不能为空")
	}
	return s.repoFactory.GetPaypalOrderLogsRepository().FindByLocalOrderID(localOrderID)
}

// GetLogByPaypalOrderID 根据PayPal订单ID获取日志
func (s *PaypalOrderLogsService) GetLogByPaypalOrderID(paypalOrderID string) (*paypal.PaypalOrderLogs, error) {
	if paypalOrderID == "" {
		return nil, errors.New("PayPal订单ID不能为空")
	}
	return s.repoFactory.GetPaypalOrderLogsRepository().FindByPaypalOrderID(paypalOrderID)
}

// GetAllOrderLogs 获取所有日志
func (s *PaypalOrderLogsService) GetAllOrderLogs() ([]paypal.PaypalOrderLogs, error) {
	return s.repoFactory.GetPaypalOrderLogsRepository().FindAll()
}