package service

import (
	"errors"
	"server/models/mp"
)

type MpUserTokenService struct {
	*Service
}

func NewMpUserTokenService(base *Service) *MpUserTokenService {
	return &MpUserTokenService{Service: base}
}

// CreateUserToken 创建用户令牌
func (s *MpUserTokenService) CreateUserToken(token *mp.MpUserToken) error {
	if token.UserID == 0 {
		return errors.New("用户ID不能为空")
	}
	if token.Token == "" {
		return errors.New("令牌不能为空")
	}
	if token.ExpireTime.IsZero() {
		return errors.New("过期时间不能为空")
	}
	
	return s.repoFactory.GetMpUserTokenRepository().Create(token)
}

// GetTokenByValue 根据Token获取令牌
func (s *MpUserTokenService) GetTokenByValue(token string) (*mp.MpUserToken, error) {
	if token == "" {
		return nil, errors.New("令牌不能为空")
	}
	return s.repoFactory.GetMpUserTokenRepository().FindByToken(token)
}

// GetTokensByUserID 根据用户ID获取令牌列表
func (s *MpUserTokenService) GetTokensByUserID(userID uint) ([]mp.MpUserToken, error) {
	if userID == 0 {
		return nil, errors.New("用户ID不能为空")
	}
	return s.repoFactory.GetMpUserTokenRepository().FindByUserID(userID)
}

// DeleteToken 删除令牌
func (s *MpUserTokenService) DeleteToken(id uint) error {
	if id == 0 {
		return errors.New("令牌ID不能为空")
	}
	return s.repoFactory.GetMpUserTokenRepository().Delete(id)
}

// DeleteAllTokensByUserID 删除用户的所有令牌
func (s *MpUserTokenService) DeleteAllTokensByUserID(userID uint) error {
	if userID == 0 {
		return errors.New("用户ID不能为空")
	}
	return s.repoFactory.GetMpUserTokenRepository().DeleteByUserID(userID)
}

// CleanupExpiredTokens 删除过期令牌
func (s *MpUserTokenService) CleanupExpiredTokens() error {
	return s.repoFactory.GetMpUserTokenRepository().DeleteExpired()
}