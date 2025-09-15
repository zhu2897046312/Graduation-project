package repository

import (
	"time"
	"gorm.io/gorm"
	"server/models/mp"
	"server/models/common"
)

type MpOrderRepository struct {
	*BaseRepository
}

func NewMpOrderRepository(db *gorm.DB) *MpOrderRepository {
	return &MpOrderRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建订单
func (r *MpOrderRepository) Create(order *mp.MpOrder) error {
	return r.db.Create(order).Error
}

// 更新订单
func (r *MpOrderRepository) Update(order *mp.MpOrder) error {
	return r.db.Save(order).Error
}

// 根据订单ID获取订单
func (r *MpOrderRepository) FindByID(id common.MyID) (*mp.MpOrder, error) {
	var order mp.MpOrder
	err := r.db.First(&order, id).Error
	return &order, err
}

// 根据用户ID获取订单列表
func (r *MpOrderRepository) FindByUserID(userID common.MyID) ([]mp.MpOrder, error) {
	var orders []mp.MpOrder
	err := r.db.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 根据状态获取订单列表
func (r *MpOrderRepository) FindByState(state int8) ([]mp.MpOrder, error) {
	var orders []mp.MpOrder
	err := r.db.Where("state = ?", state).
		Order("created_time DESC").
		Find(&orders).Error
	return orders, err
}

// 更新订单状态
func (r *MpOrderRepository) UpdateState(id common.MyID, state int8) error {
	updates := map[string]interface{}{"state": state}
	
	switch state {
	case 1: // 支付成功
		now := time.Now()
		updates["complete_time"] = &now
	case 2: // 支付失败
		now := time.Now()
		updates["fail_time"] = &now
	case 3: // 订单关闭
		now := time.Now()
		updates["close_time"] = &now
	}
	
	return r.db.Model(&mp.MpOrder{}).
		Where("id = ?", id).
		Updates(updates).Error
}

// 根据第三方支付ID获取订单
func (r *MpOrderRepository) FindByThirdID(thirdID common.MyID) (*mp.MpOrder, error) {
	var order mp.MpOrder
	err := r.db.Where("third_id = ?", thirdID).First(&order).Error
	return &order, err
}