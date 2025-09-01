package cms

import (
    "gorm.io/gorm"
    "server/models/cms"
    "server/repository/base"
)

type CmsCommentRepository struct {
    *base.BaseRepository
}

func NewCmsCommentRepository(DB *gorm.DB) *CmsCommentRepository {
    return &CmsCommentRepository{
        BaseRepository: base.NewBaseRepository(DB),
    }
}

// 根据文档ID获取评论
func (r *CmsCommentRepository) FindByDocumentID(documentID int64) ([]cms.CmsComment, error) {
    var comments []cms.CmsComment
    err := r.DB.Where("document_id = ?", documentID).Find(&comments).Error
    return comments, err
}

// 根据用户ID获取评论
func (r *CmsCommentRepository) FindByUserID(userID int64) ([]cms.CmsComment, error) {
    var comments []cms.CmsComment
    err := r.DB.Where("user_id = ?", userID).Find(&comments).Error
    return comments, err
}

// 获取顶级评论（非回复）
func (r *CmsCommentRepository) FindTopLevelComments() ([]cms.CmsComment, error) {
    var comments []cms.CmsComment
    err := r.DB.Where("reply_id = 0").Find(&comments).Error
    return comments, err
}

// 获取评论回复
func (r *CmsCommentRepository) FindReplies(commentID int64) ([]cms.CmsComment, error) {
    var comments []cms.CmsComment
    err := r.DB.Where("reply_id = ?", commentID).Find(&comments).Error
    return comments, err
}

// 分页获取评论
func (r *CmsCommentRepository) ListWithPagination(page, pageSize int) ([]cms.CmsComment, int64, error) {
    var comments []cms.CmsComment
    var total int64
    
    offset := (page - 1) * pageSize

    // 获取总数
    if err := r.DB.Model(&cms.CmsComment{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // 获取分页数据
    err := r.DB.Offset(offset).
        Limit(pageSize).
        Order("created_time DESC").
        Find(&comments).Error

    return comments, total, err
}