package service

import (
	"errors"
	"server/models/cms"
)

type CmsFileService struct {
	*Service
}

func NewCmsFileService(base *Service) *CmsFileService {
	return &CmsFileService{Service: base}
}

// CreateFile 创建文件记录
func (s *CmsFileService) CreateFile(file *cms.CmsFile) error {
	// 验证必填字段
	if file.FilePath == "" {
		return errors.New("文件路径不能为空")
	}
	if file.FileMd5 == "" {
		return errors.New("文件MD5不能为空")
	}
	
	return s.repoFactory.GetCmsFileRepository().Create(file)
}

// UpdateFile 更新文件记录
func (s *CmsFileService) UpdateFile(file *cms.CmsFile) error {
	// 验证必填字段
	if file.FilePath == "" {
		return errors.New("文件路径不能为空")
	}
	
	// 检查文件是否存在
	_, err := s.repoFactory.GetCmsFileRepository().FindByID(file.ID)
	if err != nil {
		return errors.New("文件记录不存在")
	}
	
	// 保留原始创建时间（如果模型中有该字段）
	// file.CreatedAt = existing.CreatedAt
	
	return s.repoFactory.GetCmsFileRepository().Update(file)
}

// GetFileByID 根据ID获取文件
func (s *CmsFileService) GetFileByID(id int64) (*cms.CmsFile, error) {
	if id <= 0 {
		return nil, errors.New("无效的文件ID")
	}
	return s.repoFactory.GetCmsFileRepository().FindByID(id)
}

// GetFileByMD5 根据MD5获取文件
func (s *CmsFileService) GetFileByMD5(md5 string) (*cms.CmsFile, error) {
	if md5 == "" || len(md5) != 32 {
		return nil, errors.New("无效的MD5值")
	}
	return s.repoFactory.GetCmsFileRepository().FindByMD5(md5)
}