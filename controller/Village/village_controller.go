package controller

import (
	"HRbackend/constant/share"
	service "HRbackend/service/Village"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VillageController struct {
	service service.VillageService
}

func NewVillageController() *VillageController {
	return &VillageController{
		service: service.NewVillageService(),
	}
}
func (vs *VillageController) GetVillage(c *gin.Context) {
	communcdidstr := c.Param("id")
	communceid, err := strconv.Atoi(communcdidstr)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid id")
		return
	}
	village, err := vs.service.GetByCommunceId(communceid)
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, village)
}
