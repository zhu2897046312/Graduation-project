package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpProductService struct {
	*Service
}

func NewSpProductService(base *Service) *SpProductService {
	return &SpProductService{Service: base}
}

// CreateProduct 创建商品
func (s *SpProductService) CreateProduct(product *sp.SpProduct) error {
	if product.Title == "" {
		return errors.New("商品名称不能为空")
	}
	if product.CategoryID == 0 {
		return errors.New("分类ID不能为空")
	}
	if product.Price <= 0 {
		return errors.New("商品价格必须大于0")
	}
	
	// 设置默认值
	if product.SortNum == 0 {
		product.SortNum = 99
	}
	if product.State == 0 {
		product.State = 1 // 默认上架
	}
	
	product.CreatedTime = time.Now()
	product.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpProductRepository().Create(product)
}

// UpdateProduct 更新商品
func (s *SpProductService) UpdateProduct(product *sp.SpProduct) error {
	if product.ID == 0 {
		return errors.New("商品ID不能为空")
	}
	if product.Title == "" {
		return errors.New("商品名称不能为空")
	}
	if product.Price <= 0 {
		return errors.New("商品价格必须大于0")
	}
	if product.State != 1 && product.State != 2 {
		return errors.New("无效的商品状态")
	}
	
	// 保留原始创建时间
	existing, err := s.repoFactory.GetSpProductRepository().FindByID(product.ID)
	if err != nil {
		return errors.New("商品不存在")
	}
	
	product.CreatedTime = existing.CreatedTime
	product.UpdatedTime = time.Now()
	
	return s.repoFactory.GetSpProductRepository().Update(product)
}

// GetProductByID 根据ID获取商品详情
func (s *SpProductService) GetProductByID(id uint) (*sp.SpProduct, error) {
	if id == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpProductRepository().FindByID(id)
}

// GetProductsByCategoryID 根据分类ID获取商品
func (s *SpProductService) GetProductsByCategoryID(categoryID uint) ([]sp.SpProduct, error) {
	if categoryID == 0 {
		return nil, errors.New("无效的分类ID")
	}
	return s.repoFactory.GetSpProductRepository().FindByCategoryID(categoryID)
}

// GetHotProducts 获取热门商品
func (s *SpProductService) GetHotProducts(limit int) ([]sp.SpProduct, error) {
	if limit <= 0 || limit > 100 {
		return nil, errors.New("限制数量必须在1-100之间")
	}
	return s.repoFactory.GetSpProductRepository().FindHotProducts(limit)
}

// ListProducts 分页获取商品
func (s *SpProductService) ListProducts(page, pageSize int) ([]sp.SpProduct, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.repoFactory.GetSpProductRepository().ListWithPagination(page, pageSize)
}

// UpdateStock 更新商品库存
func (s *SpProductService) UpdateStock(id uint, stock int) error {
	if stock < 0 {
		return errors.New("库存不能为负数")
	}
	return s.repoFactory.GetSpProductRepository().UpdateStock(id, stock)
}

// UpdateState 更新商品状态
func (s *SpProductService) UpdateState(id uint, state uint8) error {
	if state != 1 && state != 2 {
		return errors.New("无效的商品状态")
	}
	return s.repoFactory.GetSpProductRepository().UpdateState(id, state)
}

// IncrementSoldNum 增加销量
func (s *SpProductService) IncrementSoldNum(id uint, num uint16) error {
	if num <= 0 {
		return errors.New("增加数量必须大于0")
	}
	return s.repoFactory.GetSpProductRepository().IncrementSoldNum(id, num)
}