package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mp"
	"server/service"
	"strconv"
)

type MpUserTokenHandler struct {
	service *service.MpUserTokenService
}

func NewMpUserTokenHandler(service *service.MpUserTokenService) *MpUserTokenHandler {
	return &MpUserTokenHandler{service: service}
}

// 创建用户令牌
func (h *MpUserTokenHandler) CreateUserToken(c *gin.Context) {
	var token mp.MpUserToken
	if err := c.ShouldBindJSON(&token); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateUserToken(&token); err != nil {
		Error(c, 16001, err.Error())
		return
	}

	Success(c, token)
}

// 根据令牌获取详情
func (h *MpUserTokenHandler) GetToken(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		InvalidParams(c)
		return
	}

	tokenRecord, err := h.service.GetTokenByValue(token)
	if err != nil {
		Error(c, 16002, "令牌不存在")
		return
	}

	Success(c, tokenRecord)
}

// 获取用户所有令牌
func (h *MpUserTokenHandler) GetUserTokens(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	tokens, err := h.service.GetTokensByUserID(uint(userID))
	if err != nil {
		Error(c, 16003, "获取令牌失败")
		return
	}

	Success(c, tokens)
}

// 删除令牌
func (h *MpUserTokenHandler) DeleteToken(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteToken(uint(id)); err != nil {
		Error(c, 16004, err.Error())
		return
	}

	Success(c, nil)
}

// 删除用户所有令牌
func (h *MpUserTokenHandler) DeleteUserTokens(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil || userID == 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.DeleteAllTokensByUserID(uint(userID)); err != nil {
		Error(c, 16005, err.Error())
		return
	}

	Success(c, nil)
}

// 清理过期令牌
func (h *MpUserTokenHandler) CleanupExpiredTokens(c *gin.Context) {
	if err := h.service.CleanupExpiredTokens(); err != nil {
		Error(c, 16006, err.Error())
		return
	}

	Success(c, nil)
}