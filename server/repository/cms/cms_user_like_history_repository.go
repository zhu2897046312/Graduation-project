package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsUserLikeHistoryRepository struct {
	*base.BaseRepository
}

func NewCmsUserLikeHistoryRepository(DB *gorm.DB) *CmsUserLikeHistoryRepository {
	return &CmsUserLikeHistoryRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建点赞记录
func (r *CmsUserLikeHistoryRepository) Create(history *cms.CmsUserLikeHistory) error {
	return r.DB.Create(history).Error
}

// 更新点赞记录
func (r *CmsUserLikeHistoryRepository) Update(history *cms.CmsUserLikeHistory) error {
	return r.DB.Save(history).Error
}

// 根据用户ID获取点赞记录
func (r *CmsUserLikeHistoryRepository) FindByUserID(userID int64) ([]cms.CmsUserLikeHistory, error) {
	var histories []cms.CmsUserLikeHistory
	err := r.DB.Where("user_id = ?", userID).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}

// 根据文档ID获取点赞记录
func (r *CmsUserLikeHistoryRepository) FindByDocumentID(documentID int64) ([]cms.CmsUserLikeHistory, error) {
	var histories []cms.CmsUserLikeHistory
	err := r.DB.Where("document_id = ?", documentID).
		Order("created_time DESC").
		Find(&histories).Error
	return histories, err
}

// 检查用户是否点赞过文档
func (r *CmsUserLikeHistoryRepository) HasLiked(userID, documentID int64) (bool, error) {
	var count int64
	err := r.DB.Model(&cms.CmsUserLikeHistory{}).
		Where("user_id = ? AND document_id = ?", userID, documentID).
		Count(&count).Error
	return count > 0, err
}

// 获取文档点赞数
func (r *CmsUserLikeHistoryRepository) CountByDocumentID(documentID int64) (int64, error) {
	var count int64
	err := r.DB.Model(&cms.CmsUserLikeHistory{}).
		Where("document_id = ? AND state = 1", documentID).
		Count(&count).Error
	return count, err
}