package controller

import (
	"HRbackend/constant/share"
	service "HRbackend/service/Province"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	service service.ProvinceService
}

func NewProvinceController() ProvinceController {
	return ProvinceController{
		service: service.NewProvinceService(),
	}
}
func (pc ProvinceController) Get(c *gin.Context) {
	province, err := pc.service.GetAll()
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, province)
}
