package repository

import (
	"gorm.io/gorm"
	"server/models/cms"
)

type CmsFileRepository struct {
	*BaseRepository
}

func NewCmsFileRepository(db *gorm.DB) *CmsFileRepository {
	return &CmsFileRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// 创建文件记录
func (r *CmsFileRepository) Create(file *cms.CmsFile) error {
	return r.db.Create(file).Error
}

// 更新文件记录
func (r *CmsFileRepository) Update(file *cms.CmsFile) error {
	return r.db.Save(file).Error
}

// 根据ID获取文件
func (r *CmsFileRepository) FindByID(id int64) (*cms.CmsFile, error) {
	var file cms.CmsFile
	err := r.db.First(&file, id).Error
	return &file, err
}

// 根据MD5查找文件
func (r *CmsFileRepository) FindByMD5(md5 string) (*cms.CmsFile, error) {
	var file cms.CmsFile
	err := r.db.Where("file_md5 = ?", md5).First(&file).Error
	return &file, err
}