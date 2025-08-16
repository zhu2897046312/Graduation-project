package service

import (
	"errors"
	"server/models/sp"
	"time"
)

type SpUserCartService struct {
	*Service
}

func NewSpUserCartService(base *Service) *SpUserCartService {
	return &SpUserCartService{Service: base}
}

// AddToCart 添加到购物车
func (s *SpUserCartService) AddToCart(cart *sp.SpUserCart) error {
	if cart.UserID == 0 && cart.Fingerprint == "" {
		return errors.New("用户ID或设备指纹必须提供一项")
	}
	if cart.ProductID == 0 {
		return errors.New("商品ID不能为空")
	}
	if cart.SkuID == 0 {
		return errors.New("SKU ID不能为空")
	}
	if cart.Price <= 0 {
		return errors.New("价格必须大于0")
	}
	if cart.Quantity <= 0 {
		return errors.New("数量必须大于0")
	}
	
	cart.CreatedTime = time.Now()
	cart.UpdatedTime = time.Now()
	cart.TotalAmount = cart.Price * float64(cart.Quantity)
	cart.PayAmount = cart.Price * float64(cart.Quantity)
	
	return s.repoFactory.GetSpUserCartRepository().AddToCart(cart)
}

// UpdateCartItem 更新购物车项
func (s *SpUserCartService) UpdateCartItem(cart *sp.SpUserCart) error {
	if cart.ID == 0 {
		return errors.New("购物车项ID不能为空")
	}
	if cart.Price <= 0 {
		return errors.New("价格必须大于0")
	}
	if cart.Quantity <= 0 {
		return errors.New("数量必须大于0")
	}
	
	cart.UpdatedTime = time.Now()
	cart.TotalAmount = cart.Price * float64(cart.Quantity)
	cart.PayAmount = cart.Price * float64(cart.Quantity)
	
	return s.repoFactory.GetSpUserCartRepository().Update(cart)
}

// UpdateQuantity 更新购物车项数量
func (s *SpUserCartService) UpdateQuantity(id uint, quantity uint) error {
	if id == 0 {
		return errors.New("购物车项ID不能为空")
	}
	if quantity <= 0 {
		return errors.New("数量必须大于0")
	}
	return s.repoFactory.GetSpUserCartRepository().UpdateQuantity(id, quantity)
}

// GetCartItemsByUserID 根据用户ID获取购物车项
func (s *SpUserCartService) GetCartItemsByUserID(userID uint) ([]sp.SpUserCart, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpUserCartRepository().FindByUserID(userID)
}

// GetCartItemByProduct 根据用户ID和产品ID获取购物车项
func (s *SpUserCartService) GetCartItemByProduct(userID, productID uint) (*sp.SpUserCart, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	if productID == 0 {
		return nil, errors.New("无效的商品ID")
	}
	return s.repoFactory.GetSpUserCartRepository().FindByUserAndProduct(userID, productID)
}

// GetCartItemBySku 根据用户ID和SKU ID获取购物车项
func (s *SpUserCartService) GetCartItemBySku(userID, skuID uint) (*sp.SpUserCart, error) {
	if userID == 0 {
		return nil, errors.New("无效的用户ID")
	}
	if skuID == 0 {
		return nil, errors.New("无效的SKU ID")
	}
	return s.repoFactory.GetSpUserCartRepository().FindByUserAndSku(userID, skuID)
}

// DeleteCartItem 删除购物车项
func (s *SpUserCartService) DeleteCartItem(id uint) error {
	if id == 0 {
		return errors.New("无效的购物车项ID")
	}
	return s.repoFactory.GetSpUserCartRepository().Delete(id)
}

// ClearCart 清空用户购物车
func (s *SpUserCartService) ClearCart(userID uint) error {
	if userID == 0 {
		return errors.New("无效的用户ID")
	}
	return s.repoFactory.GetSpUserCartRepository().ClearUserCart(userID)
}

// MergeGuestCart 合并游客购物车到用户购物车
func (s *SpUserCartService) MergeGuestCart(userID uint, fingerprint string) error {
	if userID == 0 {
		return errors.New("无效的用户ID")
	}
	if fingerprint == "" {
		return errors.New("设备指纹不能为空")
	}
	return s.repoFactory.GetSpUserCartRepository().MergeGuestCart(userID, fingerprint)
}