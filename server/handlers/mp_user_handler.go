package handlers

import (
	"errors"
	"server/models/mp"
	"server/service"
	"server/models/common"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type MpUserHandler struct {
	service        *service.MpUserService
	mpTokenService *service.MpUserTokenService
}

type MpUserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type MpUserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MpUserLoginParam struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Ip        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

func NewMpUserHandler(service *service.MpUserService, mpTokenService *service.MpUserTokenService) *MpUserHandler {
	return &MpUserHandler{
		service:        service,
		mpTokenService: mpTokenService,
	}
}

// 创建用户
func (h *MpUserHandler) Register(c *gin.Context) {
	var req MpUserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	user := mp.MpUser{
		Email:    req.Email,
		Password: utils.HashPwd(req.Password),
		Nickname: req.Nickname,
	}

	if err := h.service.CreateUser(&user); err != nil {
		Error(c, 15001, err.Error())
		return
	}

	token, err_1 := h.login(&MpUserLoginParam{
		Email:     req.Email,
		Password:  req.Password,
		Ip:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	})
	if err_1 != nil {
		Error(c, 15001, err_1.Error())
		return
	}

	Success(c, token)
}

func (h *MpUserHandler) Login(c *gin.Context) {
	var req MpUserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		InvalidParams(c)
		return
	}

	token, err := h.login(&MpUserLoginParam{
		Email:     req.Email,
		Password:  utils.HashPwd(req.Password),
		Ip:        c.ClientIP(),
		UserAgent: c.Request.UserAgent(),
	})
	if err != nil {
		Error(c, 15001, err.Error())
		return
	}

	Success(c, token)
}

func (h *MpUserHandler) login(LoginRequest *MpUserLoginParam) (string, error) {
	user, err := h.service.GetUserByEmail(LoginRequest.Email)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(LoginRequest.Password, user.Password) {
		return "", errors.New("密码错误")
	}

	token, err_1 := utils.GenerateToken(int64(user.ID), user.Email)
	if err_1 != nil {
		return "", err_1
	}

	expireTime := time.Now()
	tokenEntity := &mp.MpUserToken{
		UserID:      common.MyID(user.ID),
		IP:          LoginRequest.Ip,
		Token:       token,
		ExpireTime:  &expireTime,
		CreatedTime: expireTime,
		UpdatedTime: expireTime,
		UserAgent:   LoginRequest.UserAgent,
	}
	if err_2 := h.mpTokenService.CreateUserToken(tokenEntity); err_2 != nil {
		return "", err_2
	}

	return token, nil
}

