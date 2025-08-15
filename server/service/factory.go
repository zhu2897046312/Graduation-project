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

func (f *ServiceFactory) GetCmsAssociatedPlaceService() *CmsAssociatedPlaceService {
	return NewCmsAssociatedPlaceService(f.base)
}

func (f *ServiceFactory) GetCmsCategoryService() *CmsCategoryService {
	return NewCmsCategoryService(f.base)
}

func (f *ServiceFactory) GetCmsCommentService() *CmsCommentService {
	return NewCmsCommentService(f.base)
}

func (f *ServiceFactory) GetCmsDocumentArchiveService() *CmsDocumentArchiveService {
	return NewCmsDocumentArchiveService(f.base)
}

func (f *ServiceFactory) GetCmsDocumentService() *CmsDocumentService {
	return NewCmsDocumentService(f.base)
}

func (f *ServiceFactory) GetCmsDocumentTagService() *CmsDocumentTagService {
	return NewCmsDocumentTagService(f.base)
}

func (f *ServiceFactory) GetCmsDocumentVideoService() *CmsDocumentVideoService {
	return NewCmsDocumentVideoService(f.base)
}

func (f *ServiceFactory) GetCmsFileService() *CmsFileService {
	return NewCmsFileService(f.base)
}

func (f *ServiceFactory) GetCmsRecommendIndexService() *CmsRecommendIndexService {
	return NewCmsRecommendIndexService(f.base)
}

func (f *ServiceFactory) GetCmsRecommendService() *CmsRecommendService {
	return NewCmsRecommendService(f.base)
}

func (f *ServiceFactory) GetCmsScenicSpotService() *CmsScenicSpotService {
	return NewCmsScenicSpotService(f.base)
}

func (f *ServiceFactory) GetCmsTagService() *CmsTagService {
	return NewCmsTagService(f.base)
}

func (f *ServiceFactory) GetCmsUserLikeHistoryService() *CmsUserLikeHistoryService {
	return NewCmsUserLikeHistoryService(f.base)
}