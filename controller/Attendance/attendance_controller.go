package attendance

import (
	"HRbackend/constant/share"
	"HRbackend/helper"
	attendancerequest "HRbackend/request/Attendance"
	attendance "HRbackend/service/Attendance"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceController struct {
	attendance attendance.AttendanceService
}

func NewAttendanceController() AttendanceController {
	return AttendanceController{
		attendance: attendance.NewAttendanceService(),
	}
}

func (as AttendanceController) CheckIn(c *gin.Context) {
	var input attendancerequest.AttendanceLogRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	if err := as.attendance.CheckIn(input, uint(userID)); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Check In Success")
}

func (as AttendanceController) CheckOut(c *gin.Context) {
	var input attendancerequest.AttendanceLogRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	if err := as.attendance.CheckOut(input, uint(userID)); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Check Out Success")
}

func (as AttendanceController) GetAttendanceLog(c *gin.Context) {
	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}
	filters := map[string]string{
		"branch_id":  c.Query("branch_id"),
		"islate":     c.Query("islate"),
		"name":       c.Query("name"),
		"start_date": c.Query("start_date"),
		"end_date":   c.Query("end_date"),
	}
	attendancelog, err := as.attendance.GetAttendanceLog(filters, uint(userID))
	if err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, attendancelog)
}
