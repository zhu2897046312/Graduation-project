package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpOrderOperateHistoryRepository struct {
	*base.BaseRepository
}

func NewSpOrderOperateHistoryRepository(DB *gorm.DB) *SpOrderOperateHistoryRepository {
	return &SpOrderOperateHistoryRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建操作历史
func (r *SpOrderOperateHistoryRepository) Create(history *sp.SpOrderOperateHistory) error {
	return r.DB.Create(history).Error
}

// 根据订单ID获取操作历史
func (r *SpOrderOperateHistoryRepository) FindByOrderID(orderID uint) ([]sp.SpOrderOperateHistory, error) {
	var histories []sp.SpOrderOperateHistory
	err := r.DB.Where("order_id = ?", orderID).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}

// 根据操作人获取操作历史
func (r *SpOrderOperateHistoryRepository) FindByOperateUser(user string) ([]sp.SpOrderOperateHistory, error) {
	var histories []sp.SpOrderOperateHistory
	err := r.DB.Where("operate_user = ?", user).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}