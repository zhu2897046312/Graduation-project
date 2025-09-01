package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
	"time"
)

type SpOrderRepository struct {
	*base.BaseRepository
}

func NewSpOrderRepository(DB *gorm.DB) *SpOrderRepository {
	return &SpOrderRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建订单
func (r *SpOrderRepository) Create(order *sp.SpOrder) error {
	return r.DB.Create(order).Error
}

// 更新订单
func (r *SpOrderRepository) Update(order *sp.SpOrder) error {
	return r.DB.Save(order).Error
}

// 根据ID获取订单
func (r *SpOrderRepository) FindByID(id uint) (*sp.SpOrder, error) {
	var order sp.SpOrder
	err := r.DB.First(&order, id).Error
	return &order, err
}

// 根据订单号获取订单
func (r *SpOrderRepository) FindByCode(code string) (*sp.SpOrder, error) {
	var order sp.SpOrder
	err := r.DB.Where("code = ?", code).First(&order).Error
	return &order, err
}

// 根据用户ID获取订单列表
func (r *SpOrderRepository) FindByUserID(userID uint) ([]sp.SpOrder, error) {
	var orders []sp.SpOrder
	err := r.DB.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 根据状态获取订单列表
func (r *SpOrderRepository) FindByState(state uint8) ([]sp.SpOrder, error) {
	var orders []sp.SpOrder
	err := r.DB.Where("state = ?", state).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 更新订单状态
func (r *SpOrderRepository) UpdateState(id uint, state uint8) error {
	updates := map[string]interface{}{"state": state}
	
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
	
	return r.DB.Model(&sp.SpOrder{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// 更新物流信息
func (r *SpOrderRepository) UpdateDeliveryInfo(id uint, company, sn string) error {
	return r.DB.Model(&sp.SpOrder{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"delivery_company": company,
			"delivery_sn":      sn,
		}).Error
}