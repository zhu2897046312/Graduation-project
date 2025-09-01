package cms

import (
	"gorm.io/gorm"
	"server/models/cms"
	"server/repository/base"
)

type CmsFileRepository struct {
	*base.BaseRepository
}

func NewCmsFileRepository(DB *gorm.DB) *CmsFileRepository {
	return &CmsFileRepository{
		BaseRepository: base.NewBaseRepository(DB),
	}
}

// 创建文件记录
func (r *CmsFileRepository) Create(file *cms.CmsFile) error {
	return r.DB.Create(file).Error
}

// 更新文件记录
func (r *CmsFileRepository) Update(file *cms.CmsFile) error {
	return r.DB.Save(file).Error
}

// 根据ID获取文件
func (r *CmsFileRepository) FindByID(id int64) (*cms.CmsFile, error) {
	var file cms.CmsFile
	err := r.DB.First(&file, id).Error
	return &file, err
}

// 根据MD5查找文件
func (r *CmsFileRepository) FindByMD5(md5 string) (*cms.CmsFile, error) {
	var file cms.CmsFile
	err := r.DB.Where("file_md5 = ?", md5).First(&file).Error
	return &file, err
}