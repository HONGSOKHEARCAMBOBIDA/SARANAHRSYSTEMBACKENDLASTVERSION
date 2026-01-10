package controller

import (
	"HRbackend/constant/share"
	service "HRbackend/service/District"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DistrictController struct {
	service service.DistrictService
}

func NewDistrictController() *DistrictController {
	return &DistrictController{
		service: service.NewDistrictService(),
	}
}

func (dc *DistrictController) GetDistrict(c *gin.Context) {
	provinceidstr := c.Param("id")
	provinceid, err := strconv.Atoi(provinceidstr)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	district, err := dc.service.GetByProvinceId(provinceid)
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, district)

}
