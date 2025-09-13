package handlers

import (
	"fmt"
	"net/http"
	"server/middleware"
	"server/models/core"
	"server/models/common"
	"server/service"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ListCoreAdminRequest struct {
	Nickname    string      `json:"nickname"`
	Account     string      `json:"account"`
	AdminStatus interface{} `json:"admin_status"`
	Page        interface{} `json:"page_no"`
	PageSize    interface{} `json:"page_size"`
}
type CoreAdminHandler struct {
	service       *service.CoreAdminService
	coreDept      *service.CoreDeptService
	coreRole      *service.CoreRoleService
	coreRoleIndex *service.CoreAdminRoleIndexService
	rdb           *redis.Client // 添加Redis客户端
}
type AdminCreateRequest struct {
	Nickname    string  `json:"nickname"`
	Account     string  `json:"account"`
	Password    string  `json:"pwd"`
	Mobile      string  `json:"mobile"`
	DeptID      int64   `json:"dept_id"`
	AdminStatus int8    `json:"admin_status"`
	Roles       []int64 `json:"roles"`
}
type AdminUpdateRequest struct {
	ID          int64   `json:"id"`
	Nickname    string  `json:"nickname"`
	Account     string  `json:"account"`
	Password    string  `json:"pwd"`
	Mobile      string  `json:"mobile"`
	DeptID      int64   `json:"dept_id"`
	AdminStatus int8    `json:"admin_status"`
	Roles       []int64 `json:"roles"`
}

func NewCoreAdminHandler(
	service *service.CoreAdminService,
	coreDept *service.CoreDeptService,
	coreRole *service.CoreRoleService,
	coreRoleIndex *service.CoreAdminRoleIndexService,
	rdb *redis.Client,
) *CoreAdminHandler {
	return &CoreAdminHandler{
		service:       service,
		rdb:           rdb, // 初始化Redis客户端
		coreDept:      coreDept,
		coreRole:      coreRole,
		coreRoleIndex: coreRoleIndex,
	}
}

// 创建管理员
func (h *CoreAdminHandler) CreateAdmin(c *gin.Context) {
	var req AdminCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	pwd := utils.HashPwd(req.Password)
	admin := core.CoreAdmin{
		Nickname:    req.Nickname,
		Account:     req.Account,
		Pwd:         pwd,
		Mobile:      req.Mobile,
		DeptID:      common.MyID(req.DeptID),
		AdminStatus: req.AdminStatus,
	}
	fmt.Printf("pwd: %s", admin.Pwd)
	if _, err := h.coreDept.GetDeptByID(admin.DeptID); err != nil {
		Error(c, 5001, "部门不存在")
		return
	}

	// 创建管理员
	if err := h.service.CreateAdmin(&admin); err != nil {
		Error(c, 5001, err.Error())
		return
	}
	if admin.ID == 0 {
		Error(c, 5001, "创建管理员失败")
		return
	}
	// 创建管理员角色关联
	for _, roleID := range req.Roles {
		role := core.CoreAdminRoleIndex{
			AdminID: admin.ID,
			RoleID:  common.MyID(roleID),
		}
		if err := h.coreRoleIndex.CreateAdminRole(&role); err != nil {
			fmt.Println(err)
		}
	}

	admin.Pwd = ""
	Success(c, admin)
}

// 更新管理员
func (h *CoreAdminHandler) UpdateAdmin(c *gin.Context) {
	var req AdminUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	pwd := utils.HashPwd(req.Password)
	admin := core.CoreAdmin{
		ID:          common.MyID(req.ID),
		Nickname:    req.Nickname,
		Account:     req.Account,
		Pwd:         pwd,
		Mobile:      req.Mobile,
		DeptID:      common.MyID(req.DeptID),
		AdminStatus: req.AdminStatus,
	}
	fmt.Printf("pwd: %s", admin.Pwd)
	if _, err := h.coreDept.GetDeptByID(admin.DeptID); err != nil {
		Error(c, 5001, "部门不存在")
		return
	}

	// 创建管理员
	if err := h.service.UpdateAdmin(&admin); err != nil {
		Error(c, 5001, err.Error())
		return
	}
	if admin.ID == 0 {
		Error(c, 5001, "创建管理员失败")
		return
	}
	
	err := h.coreRoleIndex.DeleteAllRolesByAdminID(common.MyID(admin.ID))
	if err != nil {
		Error(c, 5001, err.Error())
		return
	}
	// 创建管理员角色关联
	for _, roleID := range req.Roles {
		role := core.CoreAdminRoleIndex{
			AdminID: admin.ID,
			RoleID:  common.MyID(roleID),
		}
		if err := h.coreRoleIndex.CreateAdminRole(&role); err != nil {
			fmt.Println(err)
		}
	}
	admin.Pwd = ""
	Success(c, admin)
}

// 获取管理员详情
func (h *CoreAdminHandler) GetAdmin(c *gin.Context) {
	req := c.Query("id")
	if req == "" {
		InvalidParams(c)
		return
	}
	adminID := utils.ConvertToUint(req)

	admin, err := h.service.GetAdminByID(common.MyID(adminID))
	if err != nil {
		Error(c, 5003, "管理员不存在")
		return
	}
	admin.Pwd = ""

	dept, err := h.coreDept.GetDeptByID(admin.DeptID)
	if err != nil {
		Error(c, 5003, "部门不存在")
		return
	}
	roleIndexs, err_1 := h.coreRoleIndex.GetRolesByAdminID(common.MyID(adminID))
	if err_1 != nil {
		Error(c, 5003, "角色不存在")
		return
	}
	var roles []core.CoreRole
	for _, roleIndex := range roleIndexs {
		role, err_2 := h.coreRole.GetRoleByID(roleIndex.RoleID)
		if err_2 != nil {
			Error(c, 5003, "角色不存在")
			return
		}
		roles = append(roles, *role)
	}

	Success(c, gin.H{
		"roles":        roles,
		"dept":         dept,
		"id":           admin.ID,
		"nickname":     admin.Nickname,
		"account":      admin.Account,
		"mobile":       admin.Mobile,
		"admin_status": admin.AdminStatus,
		"dept_id":      admin.DeptID,
		"last_pwd":     admin.LastPwd,
		"created_time": admin.CreatedTime,
		"updated_time": admin.UpdatedTime,
	})
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

	if err := h.service.UpdateAdminStatus(common.MyID(id), req.Status); err != nil {
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

	if err := h.service.UpdateAdminPassword(common.MyID(id), req.NewPassword); err != nil {
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
		ID:         common.MyID(admin.ID),
		Nickname:   admin.Nickname,
		Permission: string(admin.Permission),
		Avatar:     "",
		DeptID:     common.MyID(admin.DeptID),
	}

	// 存储到Redis并获取token
	token, err := utils.PutSession(h.rdb, int64(admin.ID), sessionInfo)
	if err != nil {
		Error(c, 5006, "登录失败")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"result":  token,
		"error":   nil,
		"message": nil,
		"time":    1755765633795,
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

func (h *CoreAdminHandler) List(c *gin.Context) {
	var req ListCoreAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	page := utils.ConvertToUint(req.Page)
	pageSize := utils.ConvertToUint(req.PageSize)
	adminStatus := utils.ConvertToUint(req.AdminStatus)

	params := core.CoreAdminQueryParam{
		Page:        int(page),
		PageSize:    int(pageSize),
		AdminStatus: int8(adminStatus),
		Account:     req.Account,
		Nickname:    req.Nickname,
	}
	admins, total, err := h.service.List(params)
	if err != nil {
		Error(c, 5006, "获取管理员列表失败")
		return
	}
	Success(c, gin.H{
		"list":  admins,
		"total": total,
	})
}
