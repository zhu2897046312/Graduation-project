package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFileRequest 文件上传请求
type UploadFileRequest struct {
	// 可以根据需要添加其他表单字段
}

// OssHandler OSS处理程序
type OssHandler struct {
	// 可以注入OSS客户端或其他服务
}

// NewOssHandler 创建OssHandler实例
func NewOssHandler() *OssHandler {
	return &OssHandler{}
}

// UploadFile 处理文件上传
func (h *OssHandler) UploadFile(c *gin.Context) {
	// 单文件上传
	file, err := c.FormFile("file")
	if err != nil {
		InvalidParams(c)
		return
	}

	// 验证文件大小（可选）
	if file.Size > 10<<20 { // 10MB限制
		Error(c, 400, "文件大小不能超过10MB")
		return
	}

	// 获取当前时间并创建多级目录结构
	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")

	// 创建目录路径
	dirPath := fmt.Sprintf("./oss/%s/%s/%s", year, month, day)

	// 确保目录存在
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		Error(c, 500, "创建目录失败: "+err.Error())
		return
	}

	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", generateUUID(), ext)

	// 指定保存路径
	dst := fmt.Sprintf("%s/%s", dirPath, filename)

	// 保存文件到本地
	if err := c.SaveUploadedFile(file, dst); err != nil {
		Error(c, 500, "保存文件失败: "+err.Error())
		return
	}

	// 生成访问URL（直接使用静态文件服务的路径）
	accessURL := fmt.Sprintf("/%s/%s/%s/%s", year, month, day, filename)

	// 返回成功响应
	Success(c, accessURL)
}

// generateUUID 生成UUID文件名
func generateUUID() string {
	return strings.ReplaceAll(fmt.Sprintf("%s", uuid.New()), "-", "")
}

// getFullURL 获取完整的访问URL
func getFullURL(c *gin.Context, path string) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s", scheme, c.Request.Host, path)
}

// 如果需要支持多文件上传，可以使用以下方法
func (h *OssHandler) UploadMultipleFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		Error(c, 400, "获取表单失败: "+err.Error())
		return
	}

	files := form.File["files"] // 注意这里的字段名与前端一致
	var uploadedFiles []gin.H

	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	dirPath := fmt.Sprintf("./oss/%s/%s/%s", year, month, day)

	// 确保目录存在
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		Error(c, 500, "创建目录失败: "+err.Error())
		return
	}

	for _, file := range files {
		// 验证文件大小
		if file.Size > 10<<20 {
			continue // 跳过过大文件
		}

		// 生成唯一文件名
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("%s%s", generateUUID(), ext)
		dst := fmt.Sprintf("%s/%s", dirPath, filename)

		// 保存文件
		if err := c.SaveUploadedFile(file, dst); err != nil {
			continue // 跳过保存失败的文件
		}

		accessURL := fmt.Sprintf("/%s/%s/%s/%s", year, month, day, filename)

		uploadedFiles = append(uploadedFiles, gin.H{
			"original_name": file.Filename,
			"filename":      filename,
			"size":          file.Size,
			"url":           accessURL,
			"full_url":      getFullURL(c, accessURL),
			"path":          dst,
		})
	}

	Success(c, gin.H{
		"files": uploadedFiles,
		"count": len(uploadedFiles),
	})
}

// DeleteFile 删除文件
func (h *OssHandler) DeleteFile(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		InvalidParams(c)
		return
	}

	// 安全检查
	if strings.Contains(filePath, "..") || !strings.HasPrefix(filePath, "./oss/") {
		Error(c, 403, "禁止操作")
		return
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			Error(c, 404, "文件不存在")
		} else {
			Error(c, 500, "删除文件失败: "+err.Error())
		}
		return
	}

	Success(c, "文件删除成功")
}

// GetFileInfo 获取文件信息
func (h *OssHandler) GetFileInfo(c *gin.Context) {
	filePath := c.Query("path")
	if filePath == "" {
		InvalidParams(c)
		return
	}

	// 安全检查
	if strings.Contains(filePath, "..") || !strings.HasPrefix(filePath, "./oss/") {
		Error(c, 403, "禁止访问")
		return
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			Error(c, 404, "文件不存在")
		} else {
			Error(c, 500, "获取文件信息失败: "+err.Error())
		}
		return
	}

	Success(c, gin.H{
		"filename":    fileInfo.Name(),
		"size":        fileInfo.Size(),
		"mod_time":    fileInfo.ModTime(),
		"is_dir":      fileInfo.IsDir(),
		"path":        filePath,
	})
}