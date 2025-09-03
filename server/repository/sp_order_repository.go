package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
	"time"
)

type SpOrderRepository struct {
	*BaseRepository
}

func NewSpOrderRepository(db *gorm.DB) *SpOrderRepository {
	return &SpOrderRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建订单
func (r *SpOrderRepository) Create(order *sp.SpOrder) error {
	return r.db.Create(order).Error
}

// 更新订单
func (r *SpOrderRepository) Update(order *sp.SpOrder) error {
	return r.db.Updates(order).Error
}

// 根据ID获取订单
func (r *SpOrderRepository) FindByID(id uint) (*sp.SpOrder, error) {
	var order sp.SpOrder
	err := r.db.First(&order, id).Error
	return &order, err
}

// 根据订单号获取订单
func (r *SpOrderRepository) FindByCode(code string) (*sp.SpOrder, error) {
	var order sp.SpOrder
	err := r.db.Where("code = ?", code).First(&order).Error
	return &order, err
}

// 根据用户ID获取订单列表
func (r *SpOrderRepository) FindByUserID(userID uint) ([]sp.SpOrder, error) {
	var orders []sp.SpOrder
	err := r.db.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 根据状态获取订单列表
func (r *SpOrderRepository) FindByState(state uint8) ([]sp.SpOrder, error) {
	var orders []sp.SpOrder
	err := r.db.Where("state = ?", state).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 更新订单状态
func (r *SpOrderRepository) UpdateState(id uint, state uint8,remark string) error {
	updates := map[string]interface{}{
		"state": state,
		"remark": remark,
	}
	
	switch state {
	case 2: // 已支付
		now := time.Now()
		updates["payment_time"] = &now
	case 3: // 已发货
		now := time.Now()
		updates["delivery_time"] = &now
	case 4: // 已完成
		now := time.Now()
		updates["receive_time"] = &now
	}
	
	return r.db.Model(&sp.SpOrder{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// 更新物流信息
func (r *SpOrderRepository) UpdateDeliveryInfo(id uint, company, sn string) error {
	return r.db.Model(&sp.SpOrder{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"delivery_company": company,
			"delivery_sn":      sn,
		}).Error
}

func (r *SpOrderRepository) ListWithPagination(params sp.ListOrdersQueryParam) ([]sp.SpOrder, int64, error) {
	var products []sp.SpOrder
	var total int64

	// 设置默认值
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 || params.PageSize > 100 {
		params.PageSize = 10
	}
	offset := (params.Page - 1) * params.PageSize

	// 构建查询
	query := r.db.Model(&sp.SpOrder{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if params.NikeName != "" {
		query = query.Where("nickname LIKE ?", "%"+params.NikeName+"%")
	}

	if params.State != 0 {
		query = query.Where("state = ?", params.State)
	}

	if params.Code != "" {
		query = query.Where("Code LIKE ?", "%"+params.Code+"%")
	}

	if params.Email != "" {
		query = query.Where("email LIKE ?", "%"+params.Email+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Offset(offset).
		Limit(params.PageSize).
		Find(&products).Error

	return products, total, err
}