package repository

import (
	"gorm.io/gorm"
	"server/models/sp"
)

type SpOrderOperateHistoryRepository struct {
	*BaseRepository
}

func NewSpOrderOperateHistoryRepository(db *gorm.DB) *SpOrderOperateHistoryRepository {
	return &SpOrderOperateHistoryRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建操作历史
func (r *SpOrderOperateHistoryRepository) Create(history *sp.SpOrderOperateHistory) error {
	return r.db.Create(history).Error
}

// 根据订单ID获取操作历史
func (r *SpOrderOperateHistoryRepository) FindByOrderID(orderID uint) ([]sp.SpOrderOperateHistory, error) {
	var histories []sp.SpOrderOperateHistory
	err := r.db.Where("order_id = ?", orderID).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}

// 根据操作人获取操作历史
func (r *SpOrderOperateHistoryRepository) FindByOperateUser(user string) ([]sp.SpOrderOperateHistory, error) {
	var histories []sp.SpOrderOperateHistory
	err := r.db.Where("operate_user = ?", user).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}