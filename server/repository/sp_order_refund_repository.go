package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
	"time"
)

type SpOrderRefundRepository struct {
	*BaseRepository
}

func NewSpOrderRefundRepository(db *gorm.DB) *SpOrderRefundRepository {
	return &SpOrderRefundRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建退款记录
func (r *SpOrderRefundRepository) Create(refund *sp.SpOrderRefund) error {
	return r.db.Create(refund).Error
}

// 更新退款记录
func (r *SpOrderRefundRepository) Update(refund *sp.SpOrderRefund) error {
	return r.db.Save(refund).Error
}

// 根据订单ID获取退款记录
func (r *SpOrderRefundRepository) FindByOrderID(orderID uint) (*sp.SpOrderRefund, error) {
	var refund sp.SpOrderRefund
	err := r.db.Where("order_id = ?", orderID).First(&refund).Error
	return &refund, err
}

// 根据退款单号获取退款记录
func (r *SpOrderRefundRepository) FindByRefundNo(refundNo string) (*sp.SpOrderRefund, error) {
	var refund sp.SpOrderRefund
	err := r.db.Where("refund_no = ?", refundNo).First(&refund).Error
	return &refund, err
}

// 更新退款状态
func (r *SpOrderRefundRepository) UpdateStatus(id uint, status uint8) error {
	updates := map[string]interface{}{"status": status}
	
	if status == 2 { // 退款成功
		now := time.Now()
		updates["refund_time"] = &now
	}
	
	return r.db.Model(&sp.SpOrderRefund{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// 更新退款金额
func (r *SpOrderRefundRepository) UpdateRefundAmount(id uint, amount float64) error {
	return r.db.Model(&sp.SpOrderRefund{}).
		Where("id = ?", id).
		Update("refund_amount", amount).Error
}