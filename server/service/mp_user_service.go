package service

import (
	"errors"
	"server/models/mp"
)

type MpUserService struct {
	*Service
}

func NewMpUserService(base *Service) *MpUserService {
	return &MpUserService{Service: base}
}

// CreateUser 创建用户
func (s *MpUserService) CreateUser(user *mp.MpUser) error {
	if user.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if user.Password == "" {
		return errors.New("密码不能为空")
	}
	
	// 检查邮箱是否已注册
	existing, _ := s.repoFactory.GetMpUserRepository().FindByEmail(user.Email)
	if existing != nil {
		return errors.New("邮箱已被注册")
	}
	
	return s.repoFactory.GetMpUserRepository().Create(user)
}

// UpdateUser 更新用户
func (s *MpUserService) UpdateUser(user *mp.MpUser) error {
	if user.ID <= 0 {
		return errors.New("无效的用户ID")
	}
	
	// 检查用户是否存在
	existing, err := s.repoFactory.GetMpUserRepository().FindByID(user.ID)
	if err != nil {
		return errors.New("用户不存在")
	}
	
	// 保留原始密码（如果未修改）
	if user.Password == "" {
		user.Password = existing.Password
	}
	
	return s.repoFactory.GetMpUserRepository().Update(user)
}

// GetUserByID 根据ID获取用户
func (s *MpUserService) GetUserByID(id int) (*mp.MpUser, error) {
	if id <= 0 {
		return nil, errors.New("无效的用户ID")
	}
	return s.repoFactory.GetMpUserRepository().FindByID(id)
}

// GetUserByEmail 根据Email获取用户
func (s *MpUserService) GetUserByEmail(email string) (*mp.MpUser, error) {
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	return s.repoFactory.GetMpUserRepository().FindByEmail(email)
}

// UpdateUserStatus 更新用户状态
func (s *MpUserService) UpdateUserStatus(id int, status int8) error {
	if id <= 0 {
		return errors.New("无效的用户ID")
	}
	if status < 0 || status > 2 {
		return errors.New("无效的状态值")
	}
	
	return s.repoFactory.GetMpUserRepository().UpdateStatus(id, status)
}

// UpdateUserPassword 更新用户密码
func (s *MpUserService) UpdateUserPassword(id int, newPassword string) error {
	if id <= 0 {
		return errors.New("无效的用户ID")
	}
	if len(newPassword) < 6 {
		return errors.New("密码长度不能少于6位")
	}
	
	return s.repoFactory.GetMpUserRepository().UpdatePassword(id, newPassword)
}

// VerifyUserEmail 验证用户邮箱
func (s *MpUserService) VerifyUserEmail(id int) error {
	if id <= 0 {
		return errors.New("无效的用户ID")
	}
	return s.repoFactory.GetMpUserRepository().VerifyEmail(id)
}

// UpdateUserToken 更新用户token
func (s *MpUserService) UpdateUserToken(id int, token string) error {
	if id <= 0 {
		return errors.New("无效的用户ID")
	}
	if token == "" {
		return errors.New("令牌不能为空")
	}
	return s.repoFactory.GetMpUserRepository().UpdateToken(id, token)
}