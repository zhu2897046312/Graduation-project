package handlers

import (
	"github.com/gin-gonic/gin"
	"server/models/mp"
	"server/service"
	"strconv"
)

type MpUserHandler struct {
	service *service.MpUserService
}

func NewMpUserHandler(service *service.MpUserService) *MpUserHandler {
	return &MpUserHandler{service: service}
}

// 创建用户
func (h *MpUserHandler) CreateUser(c *gin.Context) {
	var user mp.MpUser
	if err := c.ShouldBindJSON(&user); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		Error(c, 15001, err.Error())
		return
	}

	Success(c, user)
}

// 更新用户
func (h *MpUserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var user mp.MpUser
	if err := c.ShouldBindJSON(&user); err != nil {
		InvalidParams(c)
		return
	}
	user.ID = id

	if err := h.service.UpdateUser(&user); err != nil {
		Error(c, 15002, err.Error())
		return
	}

	Success(c, user)
}

// 获取用户详情
func (h *MpUserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		Error(c, 15003, "用户不存在")
		return
	}

	Success(c, user)
}

// 根据邮箱获取用户
func (h *MpUserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		InvalidParams(c)
		return
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		Error(c, 15004, "用户不存在")
		return
	}

	Success(c, user)
}

// 更新用户状态
func (h *MpUserHandler) UpdateUserStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Status int8 `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateUserStatus(id, req.Status); err != nil {
		Error(c, 15005, err.Error())
		return
	}

	Success(c, nil)
}

// 更新用户密码
func (h *MpUserHandler) UpdateUserPassword(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateUserPassword(id, req.NewPassword); err != nil {
		Error(c, 15006, err.Error())
		return
	}

	Success(c, nil)
}

// 验证用户邮箱
func (h *MpUserHandler) VerifyUserEmail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	if err := h.service.VerifyUserEmail(id); err != nil {
		Error(c, 15007, err.Error())
		return
	}

	Success(c, nil)
}

// 更新用户Token
func (h *MpUserHandler) UpdateUserToken(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var req struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.UpdateUserToken(id, req.Token); err != nil {
		Error(c, 15008, err.Error())
		return
	}

	Success(c, nil)
}