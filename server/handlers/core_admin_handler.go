package handlers

import (
	"fmt"
	"net/http"
	"server/models/core"
	"server/service"
	"server/utils"
	"strconv"
	"server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type CoreAdminHandler struct {
	service *service.CoreAdminService
	rdb     *redis.Client // 添加Redis客户端
}

func NewCoreAdminHandler(service *service.CoreAdminService, rdb *redis.Client) *CoreAdminHandler {
	return &CoreAdminHandler{
		service: service,
		rdb:     rdb, // 初始化Redis客户端
	}
}

// 创建管理员
func (h *CoreAdminHandler) CreateAdmin(c *gin.Context) {
	var admin core.CoreAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		InvalidParams(c)
		return
	}

	if err := h.service.CreateAdmin(&admin); err != nil {
		Error(c, 5001, err.Error())
		return
	}

	Success(c, admin)
}

// 更新管理员
func (h *CoreAdminHandler) UpdateAdmin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	var admin core.CoreAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		InvalidParams(c)
		return
	}
	admin.ID = id

	if err := h.service.UpdateAdmin(&admin); err != nil {
		Error(c, 5002, err.Error())
		return
	}

	Success(c, admin)
}

// 获取管理员详情
func (h *CoreAdminHandler) GetAdmin(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		InvalidParams(c)
		return
	}

	admin, err := h.service.GetAdminByID(id)
	if err != nil {
		Error(c, 5003, "管理员不存在")
		return
	}

	Success(c, admin)
}

// 更新管理员状态
func (h *CoreAdminHandler) UpdateAdminStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

	if err := h.service.UpdateAdminStatus(id, req.Status); err != nil {
		Error(c, 5004, err.Error())
		return
	}

	Success(c, nil)
}

// 更新管理员密码
func (h *CoreAdminHandler) UpdateAdminPassword(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
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

	if err := h.service.UpdateAdminPassword(id, req.NewPassword); err != nil {
		Error(c, 5005, err.Error())
		return
	}

	Success(c, nil)
}

// LoginRequest 定义登录请求结构体
type LoginRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"pwd" binding:"required"`
}

// LoginResponse 定义登录响应结构体
type LoginResponse struct {
	Token string `json:"token"`
}

// Login 管理员登录
func (h *CoreAdminHandler) LoginAdmin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	fmt.Println(req)
	admin, err := h.service.GetAdminByAccount(req.Account)
	if err != nil {
		Error(c, 5006, "账号或密码错误")
		return
	}

	// 验证账号密码
	if !utils.VerifyPassword(req.Password, admin.Pwd) {
		Error(c, 5006, "账号或密码错误")
		return
	}

	// 创建会话信息
	sessionInfo := &utils.CoreLoginUserInfoModel{
		ID: admin.ID,
		Nickname: admin.Nickname,
		Permission: string(admin.Permission),
		Avatar: "",
		DeptID: admin.DeptID,
	}

	// 存储到Redis并获取token
	token, err := utils.PutSession(h.rdb, admin.ID, sessionInfo)
	if err != nil {
		Error(c, 5006, "登录失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"result": token,
		"error": nil,
		"message": nil,
		"time": 1755765633795,
	})
}

// 在 CoreAdminHandler 结构体添加方法

// GetAdminInfo 获取当前登录管理员信息
func (h *CoreAdminHandler) GetAdminInfo(c *gin.Context) {
	// 从上下文中获取用户信息
	user := middleware.GetUserFromContext(c)
	if user == nil {
		Error(c, 5006, "获取管理员信息失败")
		return
	}

	// 根据用户ID获取完整的用户信息
	admin, err := h.service.GetAdminByID(user.ID)
	if err != nil {
		Error(c, 5006, "获取管理员信息失败")
		return
	}
	
	Success(c, admin)
}

func (h *CoreAdminHandler) GetEnumDict(c *gin.Context) {
	utils.GetEnumDict()
	c.JSON(http.StatusOK, utils.GetEnumDict())
}