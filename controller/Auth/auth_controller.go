package controller

import (
	"net/http"

	"HRbackend/constant/share"
	auth "HRbackend/request/Auth"
	service "HRbackend/service/Auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController() AuthController {
	return AuthController{
		service: service.NewAuthService(),
	}
}

func (ac AuthController) Login(c *gin.Context) {
	var input auth.LoginReq

	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := ac.service.Login(input)
	if err != nil {
		share.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"user": gin.H{
			"id":      result.ID,
			"name":    result.Name,
			"phone":   result.Phone,
			"role_id": result.RoleID,
			"parts":   result.Parts,
		},
		"token": result.Token,
	})
}

func Loginversion2(c *gin.Context) {
	var req auth.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	authService := service.NewAuthService()
	result, err := authService.Login(req)
	if err != nil {
		share.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"user": gin.H{
			"id":      result.ID,
			"name":    result.Name,
			"phone":   result.Phone,
			"role_id": result.RoleID,
			"parts":   result.Parts,
		},
		"token": result.Token,
	})

}
