package service

import (
	"errors"
	"server/models/mp"
)

type MpPayConfigService struct {
	*Service
}

func NewMpPayConfigService(base *Service) *MpPayConfigService {
	return &MpPayConfigService{Service: base}
}

// CreatePayConfig 创建支付配置
func (s *MpPayConfigService) CreatePayConfig(config *mp.MpPayConfig) error {
	if config.Title == "" {
		return errors.New("支付方式名称不能为空")
	}
	if config.Code == "" {
		return errors.New("支付方式代码不能为空")
	}
	
	// 检查代码是否已存在
	existing, _ := s.repoFactory.GetMpPayConfigRepository().FindByCode(config.Code)
	if existing != nil {
		return errors.New("支付方式代码已存在")
	}
	
	return s.repoFactory.GetMpPayConfigRepository().Create(config)
}

// UpdatePayConfig 更新支付配置
func (s *MpPayConfigService) UpdatePayConfig(config *mp.MpPayConfig) error {
	if config.ID <= 0 {
		return errors.New("无效的配置ID")
	}
	if config.Title == "" {
		return errors.New("支付方式名称不能为空")
	}
	
	// 检查配置是否存在
	existing, err := s.repoFactory.GetMpPayConfigRepository().FindByID(config.ID)
	if err != nil {
		return errors.New("支付配置不存在")
	}
	
	// 检查代码冲突
	if existing.Code != config.Code {
		codeCheck, _ := s.repoFactory.GetMpPayConfigRepository().FindByCode(config.Code)
		if codeCheck != nil {
			return errors.New("支付方式代码已存在")
		}
	}
	
	return s.repoFactory.GetMpPayConfigRepository().Update(config)
}

// GetPayConfigByID 根据ID获取支付配置
func (s *MpPayConfigService) GetPayConfigByID(id int64) (*mp.MpPayConfig, error) {
	if id <= 0 {
		return nil, errors.New("无效的配置ID")
	}
	return s.repoFactory.GetMpPayConfigRepository().FindByID(id)
}

// GetActivePayConfigs 获取所有启用的支付配置
func (s *MpPayConfigService) GetActivePayConfigs() ([]mp.MpPayConfig, error) {
	return s.repoFactory.GetMpPayConfigRepository().FindActive()
}

// GetPayConfigByCode 根据Code获取支付配置
func (s *MpPayConfigService) GetPayConfigByCode(code string) (*mp.MpPayConfig, error) {
	if code == "" {
		return nil, errors.New("支付方式代码不能为空")
	}
	return s.repoFactory.GetMpPayConfigRepository().FindByCode(code)
}

// UpdatePayConfigState 更新支付配置状态
func (s *MpPayConfigService) UpdatePayConfigState(id int64, state int8) error {
	if id <= 0 {
		return errors.New("无效的配置ID")
	}
	if state < 0 || state > 1 {
		return errors.New("无效的状态值")
	}
	
	return s.repoFactory.GetMpPayConfigRepository().UpdateState(id, state)
}