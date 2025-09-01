package cms

import (
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	db                  *gorm.DB
	cms_category        *CmsCategoryRepository
	cms_comment         *CmsCommentRepository
	cms_document        *CmsDocumentRepository
	cms_place           *CmsAssociatedPlaceRepository
	cms_archive         *CmsDocumentArchiveRepository
	cms_doc_tag         *CmsDocumentTagRepository
	cms_doc_video       *CmsDocumentVideoRepository
	cms_file            *CmsFileRepository
	cms_recommend       *CmsRecommendRepository
	cms_recommend_idx   *CmsRecommendIndexRepository
	cms_scenic_spot     *CmsScenicSpotRepository
	cms_tag             *CmsTagRepository
	cms_user_like       *CmsUserLikeHistoryRepository
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

func (f *RepositoryFactory) GetCmsCommentRepository() *CmsCommentRepository {
	if f.cms_comment == nil {
		f.cms_comment = NewCmsCommentRepository(f.db)
	}
	return f.cms_comment
}

func (f *RepositoryFactory) GetCmsDocumentRepository() *CmsDocumentRepository {
	if f.cms_document == nil {
		f.cms_document = NewCmsDocumentRepository(f.db)
	}
	return f.cms_document
}

// 新增方法
func (f *RepositoryFactory) GetCmsPlaceRepository() *CmsAssociatedPlaceRepository {
	if f.cms_place == nil {
		f.cms_place = NewCmsAssociatedPlaceRepository(f.db)
	}
	return f.cms_place
}

func (f *RepositoryFactory) GetCmsDocumentArchiveRepository() *CmsDocumentArchiveRepository {
	if f.cms_archive == nil {
		f.cms_archive = NewCmsDocumentArchiveRepository(f.db)
	}
	return f.cms_archive
}

func (f *RepositoryFactory) GetCmsDocumentTagRepository() *CmsDocumentTagRepository {
	if f.cms_doc_tag == nil {
		f.cms_doc_tag = NewCmsDocumentTagRepository(f.db)
	}
	return f.cms_doc_tag
}

func (f *RepositoryFactory) GetCmsDocumentVideoRepository() *CmsDocumentVideoRepository {
	if f.cms_doc_video == nil {
		f.cms_doc_video = NewCmsDocumentVideoRepository(f.db)
	}
	return f.cms_doc_video
}

func (f *RepositoryFactory) GetCmsFileRepository() *CmsFileRepository {
	if f.cms_file == nil {
		f.cms_file = NewCmsFileRepository(f.db)
	}
	return f.cms_file
}

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

func (f *RepositoryFactory) GetCmsScenicSpotRepository() *CmsScenicSpotRepository {
	if f.cms_scenic_spot == nil {
		f.cms_scenic_spot = NewCmsScenicSpotRepository(f.db)
	}
	return f.cms_scenic_spot
}

func (f *RepositoryFactory) GetCmsTagRepository() *CmsTagRepository {
	if f.cms_tag == nil {
		f.cms_tag = NewCmsTagRepository(f.db)
	}
	return f.cms_tag
}

func (f *RepositoryFactory) GetCmsUserLikeHistoryRepository() *CmsUserLikeHistoryRepository {
	if f.cms_user_like == nil {
		f.cms_user_like = NewCmsUserLikeHistoryRepository(f.db)
	}
	return f.cms_user_like
}