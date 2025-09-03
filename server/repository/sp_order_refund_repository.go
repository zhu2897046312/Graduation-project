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

func (r *SpOrderRefundRepository) ListWithPagination(ordersID []uint, refundNo string, status uint) ([]sp.SpOrderRefund, int64, error) {
	var products []sp.SpOrderRefund
	var total int64

	// 构建查询
	query := r.db.Model(&sp.SpOrderRefund{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	// 订单ID列表匹配（如果提供了订单ID列表）
	if len(ordersID) > 0 {
		query = query.Where("order_id IN (?)", ordersID)
	}
	// 退款单号模糊匹配（如果提供了退款单号）
	if refundNo != "" {
		query = query.Where("refund_no LIKE ?", "%"+refundNo+"%")
	}
	// 状态精确匹配（如果提供了状态值）
	if status != 0 {
		query = query.Where("status = ?", status)
	}

	// 获取符合条件的总记录数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取匹配的退款记录列表
	err := query.Find(&products).Error

	return products, total, err
}

func (r *SpOrderRefundRepository) ListByOrderID(orderID uint) ([]sp.SpOrderRefund, int64, error) {
	var products []sp.SpOrderRefund
	var total int64

	// 构建查询
	query := r.db.Model(&sp.SpOrderRefund{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if orderID != 0 {
		query = query.Where("order_id = ?", orderID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Find(&products).Error

	return products, total, err
}