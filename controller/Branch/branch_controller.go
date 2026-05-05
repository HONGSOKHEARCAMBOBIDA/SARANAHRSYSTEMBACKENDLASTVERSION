package controller

import (
	"HRbackend/constant/share"
	branch "HRbackend/request/Branch"
	service "HRbackend/service/Branch"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BranchController struct {
	service service.BranchService
}

func NewBranchController() BranchController {
	return BranchController{
		service: service.NewBranchService(),
	}
}
func (bc BranchController) Create(c *gin.Context) {
	var input branch.BranchRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	defaultLat := 13.1168673
	defaultLng := 103.1970699
	if input.Latitude == nil {
		input.Latitude = &defaultLat
	}

	if input.Longitude == nil {
		input.Longitude = &defaultLng
	}
	if err := bc.service.Create(input); err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Branch Created")
}
func (bc BranchController) Get(c *gin.Context) {
	data, err := bc.service.GetAll()
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}
func (bc BranchController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var input branch.BranchRequestUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := bc.service.Update(id, input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "សាខាត្រូវបានកែប្រែ")
}
func (bc BranchController) ChangeStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := bc.service.ChangeStatus(id); err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "បានប្ដូស្ថានភាព")
}
