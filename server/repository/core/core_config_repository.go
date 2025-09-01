package core

import (
	"gorm.io/gorm"
	"server/models/core"
	"server/repository/base"
)

type CoreConfigRepository struct {
	*base.BaseRepository
}

func NewCoreConfigRepository(DB *gorm.DB) *CoreConfigRepository {
	return &CoreConfigRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建配置项
func (r *CoreConfigRepository) Create(config *core.CoreConfig) error {
	return r.DB.Create(config).Error
}

// 更新配置项
func (r *CoreConfigRepository) Update(config *core.CoreConfig) error {
	return r.DB.Updates(config).Error
}

// 根据Key获取配置
func (r *CoreConfigRepository) FindByKey(key string) (*core.CoreConfig, error) {
	var config core.CoreConfig
	err := r.DB.Where("config_key = ?", key).First(&config).Error
	return &config, err
}

// 获取所有配置项
func (r *CoreConfigRepository) FindAll() ([]core.CoreConfig, error) {
	var configs []core.CoreConfig
	err := r.DB.Order("sort_num ASC").Find(&configs).Error
	return configs, err
}

// 批量更新配置
func (r *CoreConfigRepository) BatchUpdate(configs []core.CoreConfig) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		for _, config := range configs {
			if err := tx.Save(&config).Error; err != nil {
				return err
			}
		}
		return nil
	})
}