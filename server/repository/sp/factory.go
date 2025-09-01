package sp

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	sp_category         *SpCategoryRepository
	sp_order            *SpOrderRepository
	sp_order_item       *SpOrderItemRepository
	sp_order_history    *SpOrderOperateHistoryRepository
	sp_order_address    *SpOrderReceiveAddressRepository
	sp_order_refund     *SpOrderRefundRepository
	sp_prod_attr        *SpProdAttributesRepository
	sp_prod_attr_val    *SpProdAttributesValueRepository
	sp_product          *SpProductRepository
	sp_product_content  *SpProductContentRepository
	sp_product_prop     *SpProductPropertyRepository
	sp_sku              *SpSkuRepository
	sp_sku_index        *SpSkuIndexRepository
	sp_user_address     *SpUserAddressRepository
	sp_user_cart        *SpUserCartRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetSpCategoryRepository() *SpCategoryRepository {
	if f.sp_category == nil {
		f.sp_category = NewSpCategoryRepository(f.db)
	}
	return f.sp_category
}

func (f *RepositoryFactory) GetSpOrderRepository() *SpOrderRepository {
	if f.sp_order == nil {
		f.sp_order = NewSpOrderRepository(f.db)
	}
	return f.sp_order
}

func (f *RepositoryFactory) GetSpOrderItemRepository() *SpOrderItemRepository {
	if f.sp_order_item == nil {
		f.sp_order_item = NewSpOrderItemRepository(f.db)
	}
	return f.sp_order_item
}

func (f *RepositoryFactory) GetSpOrderOperateHistoryRepository() *SpOrderOperateHistoryRepository {
	if f.sp_order_history == nil {
		f.sp_order_history = NewSpOrderOperateHistoryRepository(f.db)
	}
	return f.sp_order_history
}

func (f *RepositoryFactory) GetSpOrderReceiveAddressRepository() *SpOrderReceiveAddressRepository {
	if f.sp_order_address == nil {
		f.sp_order_address = NewSpOrderReceiveAddressRepository(f.db)
	}
	return f.sp_order_address
}

func (f *RepositoryFactory) GetSpOrderRefundRepository() *SpOrderRefundRepository {
	if f.sp_order_refund == nil {
		f.sp_order_refund = NewSpOrderRefundRepository(f.db)
	}
	return f.sp_order_refund
}

func (f *RepositoryFactory) GetSpProdAttributesRepository() *SpProdAttributesRepository {
	if f.sp_prod_attr == nil {
		f.sp_prod_attr = NewSpProdAttributesRepository(f.db)
	}
	return f.sp_prod_attr
}

func (f *RepositoryFactory) GetSpProdAttributesValueRepository() *SpProdAttributesValueRepository {
	if f.sp_prod_attr_val == nil {
		f.sp_prod_attr_val = NewSpProdAttributesValueRepository(f.db)
	}
	return f.sp_prod_attr_val
}

func (f *RepositoryFactory) GetSpProductRepository() *SpProductRepository {
	if f.sp_product == nil {
		f.sp_product = NewSpProductRepository(f.db)
	}
	return f.sp_product
}

func (f *RepositoryFactory) GetSpProductContentRepository() *SpProductContentRepository {
	if f.sp_product_content == nil {
		f.sp_product_content = NewSpProductContentRepository(f.db)
	}
	return f.sp_product_content
}

func (f *RepositoryFactory) GetSpProductPropertyRepository() *SpProductPropertyRepository {
	if f.sp_product_prop == nil {
		f.sp_product_prop = NewSpProductPropertyRepository(f.db)
	}
	return f.sp_product_prop
}

func (f *RepositoryFactory) GetSpSkuRepository() *SpSkuRepository {
	if f.sp_sku == nil {
		f.sp_sku = NewSpSkuRepository(f.db)
	}
	return f.sp_sku
}

func (f *RepositoryFactory) GetSpSkuIndexRepository() *SpSkuIndexRepository {
	if f.sp_sku_index == nil {
		f.sp_sku_index = NewSpSkuIndexRepository(f.db)
	}
	return f.sp_sku_index
}

func (f *RepositoryFactory) GetSpUserAddressRepository() *SpUserAddressRepository {
	if f.sp_user_address == nil {
		f.sp_user_address = NewSpUserAddressRepository(f.db)
	}
	return f.sp_user_address
}

func (f *RepositoryFactory) GetSpUserCartRepository() *SpUserCartRepository {
	if f.sp_user_cart == nil {
		f.sp_user_cart = NewSpUserCartRepository(f.db)
	}
	return f.sp_user_cart
}
