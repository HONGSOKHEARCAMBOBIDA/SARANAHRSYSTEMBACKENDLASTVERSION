package controller

import (
	"HRbackend/constant/share"
	service "HRbackend/service/Communce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommunceController struct {
	communceService service.CommunceService
}

func NewCommunceController() *CommunceController {
	return &CommunceController{
		communceService: service.NewCommunceService(),
	}
}

func (cc *CommunceController) GetCommunes(c *gin.Context) {
	districtIDStr := c.Param("id")

	districtID, err := strconv.Atoi(districtIDStr)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid district ID")
		return
	}

	communces, err := cc.communceService.GetByDistrictId(districtID)
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Always return array (good API practice)
	if len(communces) == 0 {
		share.RespondDate(c, http.StatusOK, []interface{}{})
		return
	}

	share.RespondDate(c, http.StatusOK, communces)
}
