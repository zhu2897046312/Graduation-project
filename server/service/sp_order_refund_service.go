package service

import (
	"errors"
	"server/models/sp"
	"server/models/common"
	"time"
	"math/rand"
	"fmt"
)

type SpOrderRefundService struct {
	*Service
}

func NewSpOrderRefundService(base *Service) *SpOrderRefundService {
	return &SpOrderRefundService{Service: base}
}

// CreateRefund 创建退款记录
func (s *SpOrderRefundService) CreateRefund(refund *sp.SpOrderRefund) error {
	if refund.OrderID == 0 {
		return errors.New("订单ID不能为空")
	}
	if refund.RefundAmount <= 0 {
		return errors.New("退款金额必须大于0")
	}

	// 生成退款单号：yyMMdd + 6位随机数（不足补零）
	datePart := time.Now().Format("060102") // yyMMdd 格式
	randomPart := fmt.Sprintf("%06d", rand.Intn(999999)) // 6位随机数
	refund.RefundNo = datePart + randomPart
	
	refund.RefundTime = time.Now()
	refund.CreatedTime = time.Now()
	refund.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderRefundRepository().Create(refund)
}

// UpdateRefund 更新退款记录
func (s *SpOrderRefundService) UpdateRefund(refund *sp.SpOrderRefund) error {
	if refund.ID == 0 {
		return errors.New("退款记录ID不能为空")
	}
	if refund.RefundAmount <= 0 {
		return errors.New("退款金额必须大于0")
	}
	if refund.Status != 1 && refund.Status != 2 && refund.Status != 3 {
		return errors.New("无效的退款状态")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpOrderRefundRepository().FindByRefundNo(refund.RefundNo)
	if err != nil {
		return errors.New("退款记录不存在")
	}
	
	refund.CreatedTime = existing.CreatedTime
	refund.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpOrderRefundRepository().Update(refund)
}

func (s *SpOrderRefundService) GetRefundByOrderIDOne(orderID common.MyID) (*sp.SpOrderRefund, error) {
	if orderID == 0 {
		return nil, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderRefundRepository().FindByOrderID(orderID)
}
// GetRefundByOrderID 根据订单ID获取退款记录
func (s *SpOrderRefundService) GetRefundByOrderID(orderID common.MyID) ([]sp.SpOrderRefund,int64, error) {
	if orderID == 0 {
		return nil,0, errors.New("无效的订单ID")
	}
	return s.repoFactory.GetSpOrderRefundRepository().ListByOrderID(orderID)
}

// GetRefundByRefundNo 根据退款单号获取退款记录
func (s *SpOrderRefundService) GetRefundByRefundNo(refundNo string) (*sp.SpOrderRefund, error) {
	if refundNo == "" {
		return nil, errors.New("退款单号不能为空")
	}
	return s.repoFactory.GetSpOrderRefundRepository().FindByRefundNo(refundNo)
}

// UpdateRefundStatus 更新退款状态
func (s *SpOrderRefundService) UpdateRefundStatus(id common.MyID, status uint8) error {
	if status != 1 && status != 2 && status != 3 {
		return errors.New("无效的退款状态")
	}
	return s.repoFactory.GetSpOrderRefundRepository().UpdateStatus(id, status)
}

// UpdateRefundAmount 更新退款金额
func (s *SpOrderRefundService) UpdateRefundAmount(id common.MyID, amount float64) error {
	if amount <= 0 {
		return errors.New("退款金额必须大于0")
	}
	return s.repoFactory.GetSpOrderRefundRepository().UpdateRefundAmount(id, amount)
}

func (s *SpOrderRefundService) ListWithPagination(ordersID []common.MyID, refundNo string, status uint) ([]sp.SpOrderRefund, int64, error) {
	return s.repoFactory.GetSpOrderRefundRepository().ListWithPagination(ordersID, refundNo, status)
}