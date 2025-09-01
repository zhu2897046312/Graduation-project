package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CoreRequestLogRepository struct {
	*base.BaseRepository
}

func NewCoreRequestLogRepository(DB *gorm.DB) *CoreRequestLogRepository {
	return &CoreRequestLogRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建请求日志
func (r *CoreRequestLogRepository) Create(log *core.CoreRequestLog) error {
	return r.DB.Create(log).Error
}

// 分页查询请求日志
func (r *CoreRequestLogRepository) ListWithPagination(page, pageSize int, tag, method string) ([]core.CoreRequestLog, int64, error) {
	var logs []core.CoreRequestLog
	var total int64

	offset := (page - 1) * pageSize

	query := r.DB.Model(&core.CoreRequestLog{})
	if tag != "" {
		query = query.Where("tag = ?", tag)
	}
	if method != "" {
		query = query.Where("request_method = ?", method)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Offset(offset).
		Limit(pageSize).
		Order("create_time DESC").
		Find(&logs).Error

	return logs, total, err
}

// 根据IP查询日志
func (r *CoreRequestLogRepository) FindByIP(ip string) ([]core.CoreRequestLog, error) {
	var logs []core.CoreRequestLog
	err := r.DB.Where("ip = ?", ip).
		Order("create_time DESC").
		Find(&logs).Error
	return logs, err
}

// 删除过期日志
func (r *CoreRequestLogRepository) DeleteOlderThan(days int) error {
	return r.DB.Where("create_time < DATE_SUB(NOW(), INTERVAL ? DAY)", days).
		Delete(&core.CoreRequestLog{}).Error
}