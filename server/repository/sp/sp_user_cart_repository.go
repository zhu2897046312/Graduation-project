package sp

import (
	"gorm.io/gorm"
	"server/models/sp"
	"server/repository/base"
)

type SpUserCartRepository struct {
	*base.BaseRepository
}

func NewSpUserCartRepository(DB *gorm.DB) *SpUserCartRepository {
	return &SpUserCartRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 添加到购物车
func (r *SpUserCartRepository) AddToCart(cart *sp.SpUserCart) error {
	return r.DB.Create(cart).Error
}

// 更新购物车项
func (r *SpUserCartRepository) Update(cart *sp.SpUserCart) error {
	return r.DB.Save(cart).Error
}

// 更新购物车项数量
func (r *SpUserCartRepository) UpdateQuantity(id uint, quantity uint) error {
	return r.DB.Model(&sp.SpUserCart{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"quantity":    quantity,
			"total_amount": gorm.Expr("price * ?", quantity),
			"pay_amount":   gorm.Expr("price * ?", quantity),
		}).Error
}

// 根据用户ID获取购物车项
func (r *SpUserCartRepository) FindByUserID(userID uint) ([]sp.SpUserCart, error) {
	var cartItems []sp.SpUserCart
	err := r.DB.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&cartItems).Error
	return cartItems, err
}

// 根据用户ID和产品ID获取购物车项
func (r *SpUserCartRepository) FindByUserAndProduct(userID, productID uint) (*sp.SpUserCart, error) {
	var cartItem sp.SpUserCart
	err := r.DB.Where("user_id = ? AND product_id = ?", userID, productID).
		First(&cartItem).Error
	return &cartItem, err
}

// 根据用户ID和SKU ID获取购物车项
func (r *SpUserCartRepository) FindByUserAndSku(userID, skuID uint) (*sp.SpUserCart, error) {
	var cartItem sp.SpUserCart
	err := r.DB.Where("user_id = ? AND sku_id = ?", userID, skuID).
		First(&cartItem).Error
	return &cartItem, err
}

// 删除购物车项
func (r *SpUserCartRepository) Delete(id uint) error {
	return r.DB.Delete(&sp.SpUserCart{}, id).Error
}

// 清空用户购物车
func (r *SpUserCartRepository) ClearUserCart(userID uint) error {
	return r.DB.Where("user_id = ?", userID).
		Delete(&sp.SpUserCart{}).Error
}

// 合并游客购物车到用户购物车
func (r *SpUserCartRepository) MergeGuestCart(userID uint, fingerprint string) error {
	return r.DB.Model(&sp.SpUserCart{}).
		Where("fingerprint = ?", fingerprint).
		Update("user_id", userID).Error
}