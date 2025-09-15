package repository

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	cms_category        *CmsCategoryRepository
	// cms_comment         *CmsCommentRepository
	cms_document        *CmsDocumentRepository
	// cms_place           *CmsAssociatedPlaceRepository
	cms_archive         *CmsDocumentArchiveRepository
	// cms_doc_tag         *CmsDocumentTagRepository
	// cms_doc_video       *CmsDocumentVideoRepository
	// cms_file            *CmsFileRepository
	cms_recommend       *CmsRecommendRepository
	cms_recommend_idx   *CmsRecommendIndexRepository
	// cms_scenic_spot     *CmsScenicSpotRepository
	// cms_tag             *CmsTagRepository
	// cms_user_like       *CmsUserLikeHistoryRepository
	core_admin          *CoreAdminRepository
	core_admin_role     *CoreAdminRoleIndexRepository
	core_config         *CoreConfigRepository
	core_dept           *CoreDeptRepository
	core_permission     *CorePermissionRepository
	// core_request_log    *CoreRequestLogRepository
	core_role           *CoreRoleRepository
	// mp_order            *MpOrderRepository
	// mp_pay_config       *MpPayConfigRepository
	// mp_product          *MpProductRepository
	// mp_reset_pwd        *MpResetPwdTokensRepository
	mp_user             *MpUserRepository
	mp_user_token       *MpUserTokenRepository
	// paypal_order_logs   *PaypalOrderLogsRepository
	// paypal_webhook_logs *PaypalWebhookLogsRepository
	shop_tag            *ShopTagRepository
	shop_tag_index      *ShopTagIndexRepository
	shop_tag_mate       *ShopTagMateRepository
	sp_category         *SpCategoryRepository
	sp_order            *SpOrderRepository
	sp_order_item       *SpOrderItemRepository
	// sp_order_history    *SpOrderOperateHistoryRepository
	sp_order_address    *SpOrderReceiveAddressRepository
	sp_order_refund     *SpOrderRefundRepository
	sp_prod_attr        *SpProdAttributesRepository
	sp_prod_attr_val    *SpProdAttributesValueRepository
	sp_product          *SpProductRepository
	sp_product_content  *SpProductContentRepository
	sp_product_prop     *SpProductPropertyRepository
	sp_sku              *SpSkuRepository
	sp_sku_index        *SpSkuIndexRepository
	// sp_user_address     *SpUserAddressRepository
	sp_user_cart        *SpUserCartRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{db: db}
}

func (f *RepositoryFactory) GetDB() *gorm.DB {
	return f.db
}

func (f *RepositoryFactory) GetCmsCategoryRepository() *CmsCategoryRepository {
	if f.cms_category == nil {
		f.cms_category = NewCmsCategoryRepository(f.db)
	}
	return f.cms_category
}

// func (f *RepositoryFactory) GetCmsCommentRepository() *CmsCommentRepository {
// 	if f.cms_comment == nil {
// 		f.cms_comment = NewCmsCommentRepository(f.db)
// 	}
// 	return f.cms_comment
// }

func (f *RepositoryFactory) GetCmsDocumentRepository() *CmsDocumentRepository {
	if f.cms_document == nil {
		f.cms_document = NewCmsDocumentRepository(f.db)
	}
	return f.cms_document
}

// 新增方法
// func (f *RepositoryFactory) GetCmsPlaceRepository() *CmsAssociatedPlaceRepository {
// 	if f.cms_place == nil {
// 		f.cms_place = NewCmsAssociatedPlaceRepository(f.db)
// 	}
// 	return f.cms_place
// }

func (f *RepositoryFactory) GetCmsDocumentArchiveRepository() *CmsDocumentArchiveRepository {
	if f.cms_archive == nil {
		f.cms_archive = NewCmsDocumentArchiveRepository(f.db)
	}
	return f.cms_archive
}

// func (f *RepositoryFactory) GetCmsDocumentTagRepository() *CmsDocumentTagRepository {
// 	if f.cms_doc_tag == nil {
// 		f.cms_doc_tag = NewCmsDocumentTagRepository(f.db)
// 	}
// 	return f.cms_doc_tag
// }

// func (f *RepositoryFactory) GetCmsDocumentVideoRepository() *CmsDocumentVideoRepository {
// 	if f.cms_doc_video == nil {
// 		f.cms_doc_video = NewCmsDocumentVideoRepository(f.db)
// 	}
// 	return f.cms_doc_video
// }

// func (f *RepositoryFactory) GetCmsFileRepository() *CmsFileRepository {
// 	if f.cms_file == nil {
// 		f.cms_file = NewCmsFileRepository(f.db)
// 	}
// 	return f.cms_file
// }

func (f *RepositoryFactory) GetCmsRecommendRepository() *CmsRecommendRepository {
	if f.cms_recommend == nil {
		f.cms_recommend = NewCmsRecommendRepository(f.db)
	}
	return f.cms_recommend
}

func (f *RepositoryFactory) GetCmsRecommendIndexRepository() *CmsRecommendIndexRepository {
	if f.cms_recommend_idx == nil {
		f.cms_recommend_idx = NewCmsRecommendIndexRepository(f.db)
	}
	return f.cms_recommend_idx
}

// func (f *RepositoryFactory) GetCmsScenicSpotRepository() *CmsScenicSpotRepository {
// 	if f.cms_scenic_spot == nil {
// 		f.cms_scenic_spot = NewCmsScenicSpotRepository(f.db)
// 	}
// 	return f.cms_scenic_spot
// }

// func (f *RepositoryFactory) GetCmsTagRepository() *CmsTagRepository {
// 	if f.cms_tag == nil {
// 		f.cms_tag = NewCmsTagRepository(f.db)
// 	}
// 	return f.cms_tag
// }

// func (f *RepositoryFactory) GetCmsUserLikeHistoryRepository() *CmsUserLikeHistoryRepository {
// 	if f.cms_user_like == nil {
// 		f.cms_user_like = NewCmsUserLikeHistoryRepository(f.db)
// 	}
// 	return f.cms_user_like
// }

func (f *RepositoryFactory) GetCoreAdminRepository() *CoreAdminRepository {
	if f.core_admin == nil {
		f.core_admin = NewCoreAdminRepository(f.db)
	}
	return f.core_admin
}

func (f *RepositoryFactory) GetCoreAdminRoleIndexRepository() *CoreAdminRoleIndexRepository {
	if f.core_admin_role == nil {
		f.core_admin_role = NewCoreAdminRoleIndexRepository(f.db)
	}
	return f.core_admin_role
}

func (f *RepositoryFactory) GetCoreConfigRepository() *CoreConfigRepository {
	if f.core_config == nil {
		f.core_config = NewCoreConfigRepository(f.db)
	}
	return f.core_config
}

func (f *RepositoryFactory) GetCoreDeptRepository() *CoreDeptRepository {
	if f.core_dept == nil {
		f.core_dept = NewCoreDeptRepository(f.db)
	}
	return f.core_dept
}

func (f *RepositoryFactory) GetCorePermissionRepository() *CorePermissionRepository {
	if f.core_permission == nil {
		f.core_permission = NewCorePermissionRepository(f.db)
	}
	return f.core_permission
}

// func (f *RepositoryFactory) GetCoreRequestLogRepository() *CoreRequestLogRepository {
// 	if f.core_request_log == nil {
// 		f.core_request_log = NewCoreRequestLogRepository(f.db)
// 	}
// 	return f.core_request_log
// }

func (f *RepositoryFactory) GetCoreRoleRepository() *CoreRoleRepository {
	if f.core_role == nil {
		f.core_role = NewCoreRoleRepository(f.db)
	}
	return f.core_role
}

// func (f *RepositoryFactory) GetMpOrderRepository() *MpOrderRepository {
// 	if f.mp_order == nil {
// 		f.mp_order = NewMpOrderRepository(f.db)
// 	}
// 	return f.mp_order
// }

// func (f *RepositoryFactory) GetMpPayConfigRepository() *MpPayConfigRepository {
// 	if f.mp_pay_config == nil {
// 		f.mp_pay_config = NewMpPayConfigRepository(f.db)
// 	}
// 	return f.mp_pay_config
// }

// func (f *RepositoryFactory) GetMpProductRepository() *MpProductRepository {
// 	if f.mp_product == nil {
// 		f.mp_product = NewMpProductRepository(f.db)
// 	}
// 	return f.mp_product
// }

// func (f *RepositoryFactory) GetMpResetPwdTokensRepository() *MpResetPwdTokensRepository {
// 	if f.mp_reset_pwd == nil {
// 		f.mp_reset_pwd = NewMpResetPwdTokensRepository(f.db)
// 	}
// 	return f.mp_reset_pwd
// }

func (f *RepositoryFactory) GetMpUserRepository() *MpUserRepository {
	if f.mp_user == nil {
		f.mp_user = NewMpUserRepository(f.db)
	}
	return f.mp_user
}

func (f *RepositoryFactory) GetMpUserTokenRepository() *MpUserTokenRepository {
	if f.mp_user_token == nil {
		f.mp_user_token = NewMpUserTokenRepository(f.db)
	}
	return f.mp_user_token
}

// func (f *RepositoryFactory) GetPaypalOrderLogsRepository() *PaypalOrderLogsRepository {
// 	if f.paypal_order_logs == nil {
// 		f.paypal_order_logs = NewPaypalOrderLogsRepository(f.db)
// 	}
// 	return f.paypal_order_logs
// }

// func (f *RepositoryFactory) GetPaypalWebhookLogsRepository() *PaypalWebhookLogsRepository {
// 	if f.paypal_webhook_logs == nil {
// 		f.paypal_webhook_logs = NewPaypalWebhookLogsRepository(f.db)
// 	}
// 	return f.paypal_webhook_logs
// }

func (f *RepositoryFactory) GetShopTagRepository() *ShopTagRepository {
	if f.shop_tag == nil {
		f.shop_tag = NewShopTagRepository(f.db)
	}
	return f.shop_tag
}

func (f *RepositoryFactory) GetShopTagIndexRepository() *ShopTagIndexRepository {
	if f.shop_tag_index == nil {
		f.shop_tag_index = NewShopTagIndexRepository(f.db)
	}
	return f.shop_tag_index
}

func (f *RepositoryFactory) GetShopTagMateRepository() *ShopTagMateRepository {
	if f.shop_tag_mate == nil {
		f.shop_tag_mate = NewShopTagMateRepository(f.db)
	}
	return f.shop_tag_mate
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

// func (f *RepositoryFactory) GetSpOrderOperateHistoryRepository() *SpOrderOperateHistoryRepository {
// 	if f.sp_order_history == nil {
// 		f.sp_order_history = NewSpOrderOperateHistoryRepository(f.db)
// 	}
// 	return f.sp_order_history
// }

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

// func (f *RepositoryFactory) GetSpUserAddressRepository() *SpUserAddressRepository {
// 	if f.sp_user_address == nil {
// 		f.sp_user_address = NewSpUserAddressRepository(f.db)
// 	}
// 	return f.sp_user_address
// }

func (f *RepositoryFactory) GetSpUserCartRepository() *SpUserCartRepository {
	if f.sp_user_cart == nil {
		f.sp_user_cart = NewSpUserCartRepository(f.db)
	}
	return f.sp_user_cart
}
