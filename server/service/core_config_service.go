package service

import (
	"errors"
	"server/models/core"
)

type CoreConfigService struct {
	*Service
}

func NewCoreConfigService(base *Service) *CoreConfigService {
	return &CoreConfigService{Service: base}
}

// CreateConfig 创建配置项
func (s *CoreConfigService) CreateConfig(config *core.CoreConfig) error {
	if config.ConfigKey == "" {
		return errors.New("配置键不能为空")
	}
	if config.ConfigValue == "" {
		return errors.New("配置值不能为空")
	}
	return s.repoFactory.GetCoreConfigRepository().Create(config)
}

// UpdateConfig 更新配置项
func (s *CoreConfigService) UpdateConfig(config *core.CoreConfig) error {
	if config.ConfigKey == "" {
		return errors.New("配置键不能为空")
	}
	
	// 检查配置是否存在
	existing, err := s.repoFactory.GetCoreConfigRepository().FindByKey(config.ConfigKey)
	if err != nil || existing.ID != config.ID {
		return errors.New("配置不存在或键值冲突")
	}
	
	return s.repoFactory.GetCoreConfigRepository().Update(config)
}

func (s *CoreConfigService) GetAllConfigs() ([]core.CoreConfig, error) {
	return s.repoFactory.GetCoreConfigRepository().FindAll()
}
// GetConfigByKey 根据Key获取配置
func (s *CoreConfigService) GetConfigByKey(key string) (*core.CoreConfig, error) {
	if key == "" {
		return nil, errors.New("配置键不能为空")
	}
	return s.repoFactory.GetCoreConfigRepository().FindByKey(key)
}

// BatchUpdateConfigs 批量更新配置
func (s *CoreConfigService) BatchUpdateConfigs(configs []core.CoreConfig) error {
	if len(configs) == 0 {
		return errors.New("配置列表不能为空")
	}
	
	// 验证每个配置项
	for _, config := range configs {
		if config.ID <= 0 {
			return errors.New("存在无效的配置ID")
		}
		if config.ConfigKey == "" {
			return errors.New("配置键不能为空")
		}
	}
	
	return s.repoFactory.GetCoreConfigRepository().BatchUpdate(configs)
}