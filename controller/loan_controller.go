package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLoan(c *gin.Context) {

	var req models.LoanRequestCreate

	if err := c.ShouldBindJSON(&req); err != nil {

		share.RespondError(c, http.StatusBadRequest, err.Error())

		return
	}
	var employee models.Employee

	if err := config.DB.First(&employee, req.EmployeeID).Error; err != nil {

		share.RespondError(c, http.StatusInternalServerError, "Branch Not Found")

	}
	// 🔹 Step 1: Check if employee already has an active loan
	var existingLoan models.Loan

	err := config.DB.Where("employee_id = ? AND status = 1", req.EmployeeID).First(&existingLoan).Error

	if err == nil {
		// 👉 Already has active loan → Add more loan amount
		var currencypair models.CurrencyPair

		if err := config.DB.Where("base_currency_id = ? AND target_currency_id = ?", req.CurrencyID, existingLoan.CurrencyID).First(&currencypair).Error; err != nil {

			share.RespondError(c, http.StatusNotFound, "Record Not found")

			return

		}

		var exchangerate models.ExchangeRate

		if err := config.DB.Where("pair_id =?", currencypair.ID).First(&exchangerate).Error; err != nil {

			share.RespondError(c, http.StatusNotFound, "Record Not found")

			return
		}

		newvalue := float64(req.LoanAmount) * exchangerate.Rate

		existingLoan.LoanAmount += newvalue

		existingLoan.RemainingAmount += newvalue

		if err := config.DB.Save(&existingLoan).Error; err != nil {

			share.RespondError(c, http.StatusInternalServerError, err.Error())

			return
		}

		share.ResponeSuccess(c, http.StatusOK, "បន្ថែមប្រាក់កម្ចីបានជោគជ័យ")

		return
	}

	// 🔹 Step 2: No existing loan → create new one
	newLoan := models.Loan{

		EmployeeID: req.EmployeeID,

		BranchID: employee.BranchID,

		CurrencyID: req.CurrencyID,

		LoanAmount: req.LoanAmount,

		RemainingAmount: req.LoanAmount,

		Status: 1, // active
	}

	if err := config.DB.Create(&newLoan).Error; err != nil {

		share.RespondError(c, http.StatusInternalServerError, err.Error())

		return
	}

	share.ResponeSuccess(c, http.StatusOK, "បង្កើតប្រាក់កម្ចីបានជោគជ័យ")
}

func GetLoan(c *gin.Context) {

	var loans []models.LoanResponse
	var employee models.Employee
	var user models.User

	employeeid := c.Query("employee_id")
	branchid := c.Query("branch_id")

	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}

	if err := config.DB.First(&user, userID).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	if err := config.DB.First(&employee, user.EmployeeID).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	db := config.DB.Table("loans").Select(`
		loans.id AS id,
		employees.id AS employee_id,
		employees.name_kh AS employee_name,
		branches.id AS branch_id,
		branches.name AS branch_name,
		loans.loan_amount AS loan_amount,
		loans.remaining_balance AS remaining_balance,
		loans.status AS status,
		c.id AS currency_id,
		c.name AS currency_name,
		c.code AS currency_code,
		c.symbol AS currency_symbol
	`).
		Joins("INNER JOIN employees ON employees.id = loans.employee_id").
		Joins("INNER JOIN branches ON branches.id = loans.branch_id").
		Joins("INNER JOIN currencies c ON c.id = loans.currency_id")

	// Permission logic
	if user.RoleID == 1 || user.RoleID == 4 || user.RoleID == 7 || user.RoleID == 8 {
		// admin / hr
		if branchid != "" {
			db = db.Where("loans.branch_id = ?", branchid)
		}
		if employeeid != "" {
			db = db.Where("loans.employee_id = ?", employeeid)
		}
	} else {
		// normal user
		db = db.Where("loans.employee_id = ?", employee.ID)
	}

	if err := db.Order("loans.id DESC").Scan(&loans).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	share.RespondDate(c, 200, loans)
}

func UpdateLoan(c *gin.Context) {

	id := c.Param("id")

	var loanupdate models.LoanRequestupdate

	if err := c.ShouldBindJSON(&loanupdate); err != nil {

		share.RespondError(c, http.StatusBadRequest, err.Error())

		return
	}
	result := config.DB.Model(&models.Loan{}).Where("id =?", id).Updates(&loanupdate)

	if result.RowsAffected == 0 {

		share.RespondError(c, http.StatusInternalServerError, "មិនមានបុគ្គលិកនេះទេ ឬក៏មិនមានការផ្លាស់ប្តូរ")

	}

	share.ResponeSuccess(c, http.StatusOK, "Loan has update")
}
