package repository

import (
	"gorm.io/gorm"
	"server/models/core"
)

type CoreConfigRepository struct {
	*BaseRepository
}

func NewCoreConfigRepository(db *gorm.DB) *CoreConfigRepository {
	return &CoreConfigRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建配置项
func (r *CoreConfigRepository) Create(config *core.CoreConfig) error {
	return r.db.Create(config).Error
}

// 更新配置项
func (r *CoreConfigRepository) Update(config *core.CoreConfig) error {
	return r.db.Updates(config).Error
}

// 根据Key获取配置
func (r *CoreConfigRepository) FindByKey(key string) (*core.CoreConfig, error) {
	var config core.CoreConfig
	err := r.db.Where("config_key = ?", key).First(&config).Error
	return &config, err
}

// 获取所有配置项
func (r *CoreConfigRepository) FindAll() ([]core.CoreConfig, error) {
	var configs []core.CoreConfig
	err := r.db.Order("sort_num ASC").Find(&configs).Error
	return configs, err
}

// 批量更新配置
func (r *CoreConfigRepository) BatchUpdate(configs []core.CoreConfig) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, config := range configs {
			if err := tx.Save(&config).Error; err != nil {
				return err
			}
		}
		return nil
	})
}