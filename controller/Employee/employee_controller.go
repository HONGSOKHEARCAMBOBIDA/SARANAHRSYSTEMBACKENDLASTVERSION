package employeecontroller

import (
	"HRbackend/constant/share"
	employeerequstupdate "HRbackend/request/Employee"
	employeeservice "HRbackend/service/Employee"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	employeeService employeeservice.EmployeeService
}

func NewEmployeeController() *EmployeeController {
	return &EmployeeController{
		employeeService: employeeservice.NewEmployeeService(),
	}
}
func (ctl *EmployeeController) GetEmployees(c *gin.Context) {
	filters := map[string]string{
		"branch_id":  c.Query("branch_id"),
		"name":       c.Query("name"),
		"role_id":    c.Query("role_id"),
		"is_active":  c.Query("is_active"),
		"shift_id":   c.Query("shift_id"),
		"is_promote": c.Query("is_promote"),
	}

	result, err := ctl.employeeService.GetEmployee(filters)
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, http.StatusOK, result)
}
func (ctl *EmployeeController) UpdateEmployee(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "invalid employee id")
		return
	}

	var input employeerequstupdate.EmployeeRequestUpdate
	if err := c.ShouldBind(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	profileImage, _ := c.FormFile("profile_image")
	qrImage, _ := c.FormFile("qr_code_bank_account")

	if err := ctl.employeeService.UpdateEmployee(
		id,
		input,
		profileImage,
		qrImage,
		c,
	); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "employee updated successfully")
}
func (ctl *EmployeeController) ChangeStatusEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "invalid employee id")
		return
	}

	if err := ctl.employeeService.ChangeStatusEmployee(id); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "status changed successfully")
}
func (ctl *EmployeeController) PromoteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "invalid employee id")
		return
	}

	if err := ctl.employeeService.PromoteEmployee(id); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	share.ResponeSuccess(c, http.StatusOK, "employee promoted successfully")
}
