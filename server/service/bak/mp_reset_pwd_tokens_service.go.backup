package service

import (
	"errors"
	"server/models/mp"
	"time"
)

type MpResetPwdTokensService struct {
	*Service
}

func NewMpResetPwdTokensService(base *Service) *MpResetPwdTokensService {
	return &MpResetPwdTokensService{Service: base}
}

// CreateResetToken 创建重置密码令牌
func (s *MpResetPwdTokensService) CreateResetToken(token *mp.MpResetPwdTokens) error {
	if token.Email == "" {
		return errors.New("邮箱不能为空")
	}
	if token.Token == "" {
		return errors.New("令牌不能为空")
	}
	if token.ExpirationTime.Before(time.Now()) {
		return errors.New("过期时间不能早于当前时间")
	}
	
	return s.repoFactory.GetMpResetPwdTokensRepository().Create(token)
}

// GetTokenRecord 根据Token获取重置密码记录
func (s *MpResetPwdTokensService) GetTokenRecord(token string) (*mp.MpResetPwdTokens, error) {
	if token == "" {
		return nil, errors.New("令牌不能为空")
	}
	return s.repoFactory.GetMpResetPwdTokensRepository().FindByToken(token)
}

// GetTokenByEmail 根据Email获取重置密码记录
func (s *MpResetPwdTokensService) GetTokenByEmail(email string) (*mp.MpResetPwdTokens, error) {
	if email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	return s.repoFactory.GetMpResetPwdTokensRepository().FindByEmail(email)
}

// IncrementTokenCount 增加令牌使用计数
func (s *MpResetPwdTokensService) IncrementTokenCount(email string) error {
	if email == "" {
		return errors.New("邮箱不能为空")
	}
	return s.repoFactory.GetMpResetPwdTokensRepository().IncrementCount(email)
}

// DeleteExpiredTokens 删除过期令牌
func (s *MpResetPwdTokensService) DeleteExpiredTokens() error {
	return s.repoFactory.GetMpResetPwdTokensRepository().DeleteExpired()
}

// DeleteTokenByEmail 删除特定令牌
func (s *MpResetPwdTokensService) DeleteTokenByEmail(email string) error {
	if email == "" {
		return errors.New("邮箱不能为空")
	}
	return s.repoFactory.GetMpResetPwdTokensRepository().Delete(email)
}