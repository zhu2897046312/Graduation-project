package service

import (
	"errors"
	"server/models/core"
)

type CoreRequestLogService struct {
	*Service
}

func NewCoreRequestLogService(base *Service) *CoreRequestLogService {
	return &CoreRequestLogService{Service: base}
}

// CreateRequestLog 创建请求日志
func (s *CoreRequestLogService) CreateRequestLog(log *core.CoreRequestLog) error {
	if log.IP == "" {
		return errors.New("IP地址不能为空")
	}
	if log.RequestMethod == "" {
		return errors.New("请求方法不能为空")
	}
	if log.RequestURL == "" {
		return errors.New("请求路径不能为空")
	}
	
	return s.repoFactory.GetCoreRequestLogRepository().Create(log)
}

// ListRequestLogs 分页查询请求日志
func (s *CoreRequestLogService) ListRequestLogs(page, pageSize int, tag, method string) ([]core.CoreRequestLog, int64, error) {
	// 验证分页参数
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	
	return s.repoFactory.GetCoreRequestLogRepository().ListWithPagination(page, pageSize, tag, method)
}

// GetLogsByIP 根据IP查询日志
func (s *CoreRequestLogService) GetLogsByIP(ip string) ([]core.CoreRequestLog, error) {
	if ip == "" {
		return nil, errors.New("IP地址不能为空")
	}
	return s.repoFactory.GetCoreRequestLogRepository().FindByIP(ip)
}

// CleanupOldLogs 删除过期日志
func (s *CoreRequestLogService) CleanupOldLogs(days int) error {
	if days <= 0 {
		return errors.New("保留天数必须大于0")
	}
	return s.repoFactory.GetCoreRequestLogRepository().DeleteOlderThan(days)
}