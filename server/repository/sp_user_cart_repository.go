package repository

import (
	"server/models/sp"
	"server/models/common"
	"gorm.io/gorm"
)

type SpUserCartRepository struct {
	*BaseRepository
}

func NewSpUserCartRepository(db *gorm.DB) *SpUserCartRepository {
	return &SpUserCartRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 添加到购物车
func (r *SpUserCartRepository) AddToCart(cart *sp.SpUserCart) error {
	return r.db.Create(cart).Error
}

// 更新购物车项
func (r *SpUserCartRepository) Update(cart *sp.SpUserCart) error {
	return r.db.Updates(cart).Error
}

// 更新购物车项数量
func (r *SpUserCartRepository) UpdateQuantity(id common.MyID, quantity uint) error {
	return r.db.Model(&sp.SpUserCart{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"quantity":     quantity,
			"total_amount": gorm.Expr("price * ?", quantity),
			"pay_amount":   gorm.Expr("price * ?", quantity),
		}).Error
}

// 根据用户ID获取购物车项
func (r *SpUserCartRepository) FindByUserID(userID common.MyID) ([]sp.SpUserCart, error) {
	var cartItems []sp.SpUserCart
	err := r.db.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&cartItems).Error
	return cartItems, err
}

// 根据用户ID和产品ID获取购物车项
func (r *SpUserCartRepository) FindByUserAndProduct(userId common.MyID, fingerprint string, productID common.MyID, skuID common.MyID) (*sp.SpUserCart, error) {
	var cartItem sp.SpUserCart
	db := r.db.Model(&sp.SpUserCart{})
	
	// 动态添加查询条件
	if userId != 0 {
		db = db.Where("user_id = ?", userId)
	}
	if fingerprint != "" {
		db = db.Where("fingerprint = ?", fingerprint)
	}
	if productID != 0 {
		db = db.Where("product_id = ?", productID)
	}
	if skuID != 0 {
		db = db.Where("sku_id = ?", skuID)
	}
	
	err := db.First(&cartItem).Error
	return &cartItem, err
}

// 根据用户ID和SKU ID获取购物车项
func (r *SpUserCartRepository) FindByUserAndSku(userID, skuID common.MyID) (*sp.SpUserCart, error) {
	var cartItem sp.SpUserCart
	err := r.db.Where("user_id = ? AND sku_id = ?", userID, skuID).
		First(&cartItem).Error
	return &cartItem, err
}

// 删除购物车项
func (r *SpUserCartRepository) Delete(id common.MyID) error {
	return r.db.Delete(&sp.SpUserCart{}, id).Error
}

// 清空用户购物车
func (r *SpUserCartRepository) ClearUserCart(userID common.MyID) error {
	return r.db.Where("user_id = ?", userID).
		Delete(&sp.SpUserCart{}).Error
}

func (r *SpUserCartRepository) ClearFingerprintCart(Fingerprint string) error {
	return r.db.Where("fingerprint = ?", Fingerprint).
		Delete(&sp.SpUserCart{}).Error
}

// 合并游客购物车到用户购物车
func (r *SpUserCartRepository) MergeGuestCart(userID common.MyID, fingerprint string) error {
	return r.db.Model(&sp.SpUserCart{}).
		Where("fingerprint = ?", fingerprint).
		Update("user_id", userID).Error
}

func (r *SpUserCartRepository) ListWithPagination(userId common.MyID, fingerprint string) ([]sp.SpUserCart, int64, error) {
	var tags []sp.SpUserCart
	var total int64

	// 构建查询
	query := r.db.Model(&sp.SpUserCart{}).Where("deleted_time IS NULL")

	// 应用过滤条件
	if userId != 0 {
		query = query.Where("user_id = ?", userId)
	}

	if fingerprint != "" {
		query = query.Where("fingerprint LIKE ?", "%"+fingerprint+"%")
	}

	// if productID != 0 {
	// 	query = query.Where("product_id = ?", productID)
	// }

	// if skuID != 0 {
	// 	query = query.Where("sku_id = ?", skuID)
	// }

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Find(&tags).Error
	return tags, total, err
}
