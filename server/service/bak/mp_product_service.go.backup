package service

import (
	"errors"
	"server/models/mp"
	"server/models/common"
)

type MpProductService struct {
	*Service
}

func NewMpProductService(base *Service) *MpProductService {
	return &MpProductService{Service: base}
}

// CreateProduct 创建产品
func (s *MpProductService) CreateProduct(product *mp.MpProduct) error {
	if product.Title == "" {
		return errors.New("产品名称不能为空")
	}
	if product.Code == "" {
		return errors.New("产品代码不能为空")
	}
	if product.Price <= 0 {
		return errors.New("产品价格必须大于0")
	}
	
	// 设置默认状态
	if product.State == 0 {
		product.State = 1 // 默认启用
	}
	
	return s.repoFactory.GetMpProductRepository().Create(product)
}

// UpdateProduct 更新产品
func (s *MpProductService) UpdateProduct(product *mp.MpProduct) error {
	if product.ID <= 0 {
		return errors.New("无效的产品ID")
	}
	if product.Title == "" {
		return errors.New("产品名称不能为空")
	}
	
	// 检查产品是否存在
	existing, err := s.repoFactory.GetMpProductRepository().FindByID(product.ID)
	if err != nil {
		return errors.New("产品不存在")
	}
	
	// 检查代码冲突
	if existing.Code != product.Code {
		codeCheck, _ := s.repoFactory.GetMpProductRepository().FindByCode(product.Code)
		if codeCheck != nil {
			return errors.New("产品代码已存在")
		}
	}
	
	return s.repoFactory.GetMpProductRepository().Update(product)
}

// GetProductByID 根据ID获取产品
func (s *MpProductService) GetProductByID(id common.MyID) (*mp.MpProduct, error) {
	if id <= 0 {
		return nil, errors.New("无效的产品ID")
	}
	return s.repoFactory.GetMpProductRepository().FindByID(id)
}

// GetProductsByType 根据产品类型获取产品列表
func (s *MpProductService) GetProductsByType(productType int8) ([]mp.MpProduct, error) {
	if productType < 0 {
		return nil, errors.New("无效的产品类型")
	}
	return s.repoFactory.GetMpProductRepository().FindByType(productType)
}

// GetProductsByTerminal 根据终端类型获取产品列表
func (s *MpProductService) GetProductsByTerminal(terminalType int8) ([]mp.MpProduct, error) {
	if terminalType < 0 {
		return nil, errors.New("无效的终端类型")
	}
	return s.repoFactory.GetMpProductRepository().FindByTerminal(terminalType)
}

// GetProductByCode 根据产品代码获取产品
func (s *MpProductService) GetProductByCode(code string) (*mp.MpProduct, error) {
	if code == "" {
		return nil, errors.New("产品代码不能为空")
	}
	return s.repoFactory.GetMpProductRepository().FindByCode(code)
}

// UpdateProductState 更新产品状态
func (s *MpProductService) UpdateProductState(id common.MyID, state int) error {
	if id <= 0 {
		return errors.New("无效的产品ID")
	}
	if state < 0 || state > 1 {
		return errors.New("无效的状态值")
	}
	
	return s.repoFactory.GetMpProductRepository().UpdateState(id, state)
}