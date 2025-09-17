package handlers

import (
	"server/models/core"
	"server/service"
	"server/models/common"
	"server/utils"
	"sort"

	"github.com/gin-gonic/gin"
)

type CorePermissionHandler struct {
	service *service.CorePermissionService
}

type PermissionCreateRequest struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Remark string `json:"remark"`
	Pid    int64  `json:"pid"`
}

type PermissionUpdateRequest struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Code   string `json:"code"`
	Remark string `json:"remark"`
	Pid    int64  `json:"pid"`
}

func NewCorePermissionHandler(service *service.CorePermissionService) *CorePermissionHandler {
	return &CorePermissionHandler{service: service}
}

// 创建权限
func (h *CorePermissionHandler) CreatePermission(c *gin.Context) {
	var req PermissionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	permission := core.CorePermission{
		Title:  req.Title,
		Code:   req.Code,
		Remark: req.Remark,
		Pid:    common.MyID(req.Pid),
	}
	if err := h.service.CreatePermission(&permission); err != nil {
		Error(c, 8001, err.Error())
		return
	}

	Success(c, permission)
}

// 更新权限
func (h *CorePermissionHandler) UpdatePermission(c *gin.Context) {
	var req PermissionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}
	permission := core.CorePermission{
		ID:     common.MyID(req.ID),
		Title:  req.Title,
		Code:   req.Code,
		Remark: req.Remark,
		Pid:    common.MyID(req.Pid),
	}
	if err := h.service.UpdatePermission(&permission); err != nil {
		Error(c, 8001, err.Error())
		return
	}

	Success(c, permission)
}

// 获取权限详情
func (h *CorePermissionHandler) GetPermission(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}

	uID := utils.ConvertToUint(id)

	permission, err := h.service.GetPermissionByID(common.MyID(uID))
	if err != nil {
		Error(c, 8003, "权限不存在")
		return
	}

	Success(c, permission)
}

func (h *CorePermissionHandler) List(c *gin.Context) {
	onlyTop := c.Query("onlyTop") == "true" // 获取是否只返回顶级分类

	// 获取所有分类
	permissions, err := h.service.GetAll()
	if err != nil {
		Error(c, 22007, "获取分类失败")
		return
	}

	// 构建与Java类似的扁平列表
	result := h.buildFlatList(permissions, onlyTop)

	Success(c, result)
}

// 构建扁平列表结构
func (h *CorePermissionHandler) buildFlatList(allPermissions []core.CorePermission, onlyTop bool) []core.CorePermission {
	// 先获取顶级分类（PID=0）
	var topList []core.CorePermission
	for _, perm := range allPermissions {
		if perm.Pid == 0 {
			topList = append(topList, perm)
		}
	}

	// 按ID排序顶级分类
	sort.Slice(topList, func(i, j int) bool {
		return topList[i].ID < topList[j].ID
	})

	if onlyTop {
		// 如果只需要顶级分类，直接返回
		return topList
	}

	// 获取非顶级分类
	var secList []core.CorePermission
	for _, perm := range allPermissions {
		if perm.Pid != 0 {
			secList = append(secList, perm)
		}
	}

	// 按ID排序非顶级分类
	sort.Slice(secList, func(i, j int) bool {
		return secList[i].ID < secList[j].ID
	})

	// 构建结果列表
	var result []core.CorePermission
	for _, top := range topList {
		// 添加顶级分类
		result = append(result, top)

		// 添加该顶级分类下的子分类
		for _, sec := range secList {
			if sec.Pid == common.MyID(top.ID) {
				result = append(result, sec)
			}
		}
	}

	return result
}

func (h *CorePermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		InvalidParams(c)
		return
	}

	uID := utils.ConvertToUint(id)
	if uID == 0 {
		InvalidParams(c)
		return
	}			
	if err := h.service.Delete(common.MyID(uID)); err != nil {
		Error(c, 8002, err.Error())
		return
	}	
	Success(c, nil)
}

func (h *CorePermissionHandler) TopList(c *gin.Context) {
	onlyTop := c.Query("onlyTop") == "true" // 获取是否只返回顶级分类

	// 获取所有分类
	permissions, err := h.service.GetAll()
	if err != nil {
		Error(c, 22007, "获取分类失败")
		return
	}

	// 构建与Java类似的扁平列表
	result := h.buildFlatList(permissions, onlyTop)

	Success(c, result)
}