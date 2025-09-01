package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsTagRepository struct {
	*base.BaseRepository
}

func NewCmsTagRepository(DB *gorm.DB) *CmsTagRepository {
	return &CmsTagRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建标签
func (r *CmsTagRepository) Create(tag *cms.CmsTag) error {
	return r.DB.Create(tag).Error
}

// 更新标签
func (r *CmsTagRepository) Update(tag *cms.CmsTag) error {
	return r.DB.Save(tag).Error
}

// 根据ID获取标签
func (r *CmsTagRepository) FindByID(id int64) (*cms.CmsTag, error) {
	var tag cms.CmsTag
	err := r.DB.First(&tag, id).Error
	return &tag, err
}

// 根据状态获取标签列表
func (r *CmsTagRepository) FindByState(state int8) ([]cms.CmsTag, error) {
	var tags []cms.CmsTag
	err := r.DB.Where("state = ?", state).
		Order("sort_num ASC").
		Find(&tags).Error
	return tags, err
}

// 增加标签阅读量
func (r *CmsTagRepository) IncrementReadNum(id int64) error {
	return r.DB.Model(&cms.CmsTag{}).
		Where("id = ?", id).
		Update("read_num", gorm.Expr("read_num + ?", 1)).Error
}

// 搜索标签
func (r *CmsTagRepository) Search(keyword string) ([]cms.CmsTag, error) {
	var tags []cms.CmsTag
	err := r.DB.Where("title LIKE ? OR keyword LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Order("sort_num ASC").
		Find(&tags).Error
	return tags, err
}