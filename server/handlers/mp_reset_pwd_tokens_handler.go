package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mp"
	"server/service"
	"time"
)

type MpResetPwdTokensHandler struct {
	service *service.MpResetPwdTokensService
}

func NewMpResetPwdTokensHandler(service *service.MpResetPwdTokensService) *MpResetPwdTokensHandler {
	return &MpResetPwdTokensHandler{service: service}
}

// 创建重置令牌
func (h *MpResetPwdTokensHandler) CreateResetToken(c *gin.Context) {
	var token mp.MpResetPwdTokens
	if err := c.ShouldBindJSON(&token); err != nil {
		InvalidParams(c)
		return
	}

	// 设置默认过期时间为1小时后
	if token.ExpirationTime.IsZero() {
		token.ExpirationTime = time.Now().Add(time.Hour)
	}

	if err := h.service.CreateResetToken(&token); err != nil {
		Error(c, 14001, err.Error())
		return
	}

	Success(c, token)
}

// 根据令牌获取记录
func (h *MpResetPwdTokensHandler) GetTokenRecord(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		InvalidParams(c)
		return
	}

	record, err := h.service.GetTokenRecord(token)
	if err != nil {
		Error(c, 14002, "令牌记录不存在")
		return
	}

	Success(c, record)
}

// 根据邮箱获取记录
func (h *MpResetPwdTokensHandler) GetTokenByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		InvalidParams(c)
		return
	}

	record, err := h.service.GetTokenByEmail(email)
	if err != nil {
		Error(c, 14003, "令牌记录不存在")
		return
	}

	Success(c, record)
}

// 增加令牌计数
func (h *MpResetPwdTokensHandler) IncrementTokenCount(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		InvalidParams(c)
		return
	}

	if err := h.service.IncrementTokenCount(email); err != nil {
		Error(c, 14004, err.Error())
		return
	}

	Success(c, nil)
}

// 删除过期令牌
func (h *MpResetPwdTokensHandler) DeleteExpiredTokens(c *gin.Context) {
	if err := h.service.DeleteExpiredTokens(); err != nil {
		Error(c, 14005, err.Error())
		return
	}

	Success(c, nil)
}

// 删除邮箱令牌
func (h *MpResetPwdTokensHandler) DeleteTokenByEmail(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteTokenByEmail(email); err != nil {
		Error(c, 14006, err.Error())
		return
	}

	Success(c, nil)
}