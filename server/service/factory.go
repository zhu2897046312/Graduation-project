package service

import (
	"github.com/go-redis/redis/v8"
)

type ServiceFactory struct {
	base *Service
	redisClient *redis.Client
}

func NewServiceFactory(base *Service, redisClient *redis.Client) *ServiceFactory {
	return &ServiceFactory{
		base:     base,
		redisClient: redisClient,
	}
}

// func (f *ServiceFactory) GetCmsAssociatedPlaceService() *CmsAssociatedPlaceService {
// 	return NewCmsAssociatedPlaceService(f.base)
// }

// func (f *ServiceFactory) GetCmsCategoryService() *CmsCategoryService {
// 	return NewCmsCategoryService(f.base)
// }

// func (f *ServiceFactory) GetCmsCommentService() *CmsCommentService {
// 	return NewCmsCommentService(f.base)
// }

func (f *ServiceFactory) GetCmsDocumentArchiveService() *CmsDocumentArchiveService {
	return NewCmsDocumentArchiveService(f.base)
}

func (f *ServiceFactory) GetCmsDocumentService() *CmsDocumentService {
	return NewCmsDocumentService(f.base)
}

// func (f *ServiceFactory) GetCmsDocumentTagService() *CmsDocumentTagService {
// 	return NewCmsDocumentTagService(f.base)
// }

// func (f *ServiceFactory) GetCmsDocumentVideoService() *CmsDocumentVideoService {
// 	return NewCmsDocumentVideoService(f.base)
// }

// func (f *ServiceFactory) GetCmsFileService() *CmsFileService {
// 	return NewCmsFileService(f.base)
// }

func (f *ServiceFactory) GetCmsRecommendIndexService() *CmsRecommendIndexService {
	return NewCmsRecommendIndexService(f.base)
}

func (f *ServiceFactory) GetCmsRecommendService() *CmsRecommendService {
	return NewCmsRecommendService(f.base)
}

// func (f *ServiceFactory) GetCmsScenicSpotService() *CmsScenicSpotService {
// 	return NewCmsScenicSpotService(f.base)
// }

// func (f *ServiceFactory) GetCmsTagService() *CmsTagService {
// 	return NewCmsTagService(f.base)
// }

// func (f *ServiceFactory) GetCmsUserLikeHistoryService() *CmsUserLikeHistoryService {
// 	return NewCmsUserLikeHistoryService(f.base)
// }

func (f *ServiceFactory) GetCoreAdminService() *CoreAdminService {
	return NewCoreAdminService(f.base)
}

func (f *ServiceFactory) GetCoreAdminRoleIndexService() *CoreAdminRoleIndexService {
	return NewCoreAdminRoleIndexService(f.base)
}

func (f *ServiceFactory) GetCoreConfigService() *CoreConfigService {
	return NewCoreConfigService(f.base)
}

func (f *ServiceFactory) GetCoreDeptService() *CoreDeptService {
	return NewCoreDeptService(f.base)
}

func (f *ServiceFactory) GetCorePermissionService() *CorePermissionService {
	return NewCorePermissionService(f.base)
}

// func (f *ServiceFactory) GetCoreRequestLogService() *CoreRequestLogService {
// 	return NewCoreRequestLogService(f.base)
// }

func (f *ServiceFactory) GetCoreRoleService() *CoreRoleService {
	return NewCoreRoleService(f.base)
}

// 添加MP服务的工厂方法
// func (f *ServiceFactory) GetMpOrderService() *MpOrderService {
// 	return NewMpOrderService(f.base)
// }

// func (f *ServiceFactory) GetMpPayConfigService() *MpPayConfigService {
// 	return NewMpPayConfigService(f.base)
// }

// func (f *ServiceFactory) GetMpProductService() *MpProductService {
// 	return NewMpProductService(f.base)
// }

// func (f *ServiceFactory) GetMpResetPwdTokensService() *MpResetPwdTokensService {
// 	return NewMpResetPwdTokensService(f.base)
// }

func (f *ServiceFactory) GetMpUserService() *MpUserService {
	return NewMpUserService(f.base)
}

func (f *ServiceFactory) GetMpUserTokenService() *MpUserTokenService {
	return NewMpUserTokenService(f.base)
}

// 添加PayPal服务的工厂方法
func (f *ServiceFactory) GetPaypalOrderLogsService() *PaypalOrderLogsService {
	return NewPaypalOrderLogsService(f.base)
}

func (f *ServiceFactory) GetPaypalWebhookLogsService() *PaypalWebhookLogsService {
	return NewPaypalWebhookLogsService(f.base)
}

// 添加商店服务的工厂方法
func (f *ServiceFactory) GetShopTagIndexService() *ShopTagIndexService {
	return NewShopTagIndexService(f.base)
}

func (f *ServiceFactory) GetShopTagMateService() *ShopTagMateService {
	return NewShopTagMateService(f.base)
}

func (f *ServiceFactory) GetShopTagService() *ShopTagService {
	return NewShopTagService(f.base)
}

func (f *ServiceFactory) GetSpCategoryService() *SpCategoryService {
	return NewSpCategoryService(f.base)
}

func (f *ServiceFactory) GetSpOrderItemService() *SpOrderItemService {
	return NewSpOrderItemService(f.base)
}

// func (f *ServiceFactory) GetSpOrderOperateHistoryService() *SpOrderOperateHistoryService {
// 	return NewSpOrderOperateHistoryService(f.base)
// }

func (f *ServiceFactory) GetSpOrderReceiveAddressService() *SpOrderReceiveAddressService {
	return NewSpOrderReceiveAddressService(f.base)
}

func (f *ServiceFactory) GetSpOrderRefundService() *SpOrderRefundService {
	return NewSpOrderRefundService(f.base)
}

func (f *ServiceFactory) GetSpOrderService() *SpOrderService {
	return NewSpOrderService(f.base)
}

func (f *ServiceFactory) GetSpProdAttributesService() *SpProdAttributesService {
	return NewSpProdAttributesService(f.base)
}

func (f *ServiceFactory) GetSpProdAttributesValueService() *SpProdAttributesValueService {
	return NewSpProdAttributesValueService(f.base)
}

func (f *ServiceFactory) GetSpProductContentService() *SpProductContentService {
	return NewSpProductContentService(f.base)
}

func (f *ServiceFactory) GetSpProductPropertyService() *SpProductPropertyService {
	return NewSpProductPropertyService(f.base)
}

func (f *ServiceFactory) GetSpProductService() *SpProductService {
	return NewSpProductService(f.base)
}

func (f *ServiceFactory) GetSpSkuIndexService() *SpSkuIndexService {
	return NewSpSkuIndexService(f.base)
}

func (f *ServiceFactory) GetSpSkuService() *SpSkuService {
	return NewSpSkuService(f.base)
}

func (f *ServiceFactory) GetSpUserAddressService() *SpUserAddressService {
	return NewSpUserAddressService(f.base)
}

func (f *ServiceFactory) GetSpUserCartService() *SpUserCartService {
	return NewSpUserCartService(f.base)
}