package controller

import (
	"HRbackend/config"
	"HRbackend/constant/share"
	"HRbackend/helper"
	models "HRbackend/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePayroll(c *gin.Context) {
	var inputs []models.PayrollRequestCreate

	userid, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "login please")
		return
	}

	// Bind array input
	if err := c.ShouldBindJSON(&inputs); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	if len(inputs) == 0 {
		share.RespondError(c, http.StatusBadRequest, "គ្មានប្រាក់ខែផ្ដល់")
		return
	}

	currentYear := time.Now().Year()

	for _, input := range inputs {

		// Create payroll
		p := models.Payroll{
			SalaryID:                      input.SalaryID,
			PayrollYear:                   currentYear,
			PayrollMonth:                  input.PayrollMonth,
			GrossSalary:                   input.GrossSalary,
			TotalLate:                     input.TotalLate,
			LatePenalty:                   input.LatePenalty,
			ToalEarlyexit:                 input.ToalEarlyexit,
			TotalExitpenalty:              input.TotalExitpenalty,
			LeaveWithPermission:           input.LeaveWithPermission,
			PeanaltyLeaveWithPermission:   input.PeanaltyLeaveWithPermission,
			LeaveWithoutPermission:        input.LeaveWithoutPermission,
			PenaltyLeaveWithoutPermission: input.PenaltyLeaveWithoutPermission,
			LeaveWeekend:                  input.LeaveWeekend,
			PenaltyLeaveWeekend:           input.PenaltyLeaveWeekend,
			LoanDeduction:                 input.LoanDeduction,
			IsAttendanceBonus:             input.IsAttendanceBonus,
			BonusType:                     input.BonusType,
			BonusAmount:                   input.BonusAmount,
			TotalDeduction:                input.TotalDeduction,
			NetSalary:                     input.NetSalary,
			Status:                        1,
			BranchId:                      input.BranchId,
			CurrencyID:                    input.CurrencyID,
		}

		if err := config.DB.Create(&p).Error; err != nil {
			share.RespondError(c, http.StatusInternalServerError, "មិនអាចបង្កើតប្រាក់ខែបានទេ")
			return
		}

		// Loan deduction
		if input.LoanID != 0 && input.LoanDeduction > 0 {

			var loan models.Loan
			if err := config.DB.First(&loan, input.LoanID).Error; err != nil {
				continue // loan not found → skip
			}

			// Find currency pair
			var currencypair models.CurrencyPair
			if err := config.DB.
				Where("base_currency_id = ? AND target_currency_id = ?", input.CurrencyID, loan.CurrencyID).
				// Where("base_currency_id = ? AND target_currency_id = ?", loan.CurrencyID, input.CurrencyID).
				First(&currencypair).Error; err != nil {

				share.RespondError(c, http.StatusNotFound, "ទិន្នន័យរូបិយប័ណ្ណរកមិនឃើញ")
				return
			}

			// Find exchange rate
			var exchangerate models.ExchangeRate
			if err := config.DB.
				Where("pair_id = ?", currencypair.ID).
				First(&exchangerate).Error; err != nil {

				share.RespondError(c, http.StatusNotFound, "អត្រាប្តូរប្រាក់រកមិនឃើញ")
				return
			}

			// Deduct loan
			newRemaining := loan.RemainingAmount - (float64(input.LoanDeduction) * exchangerate.Rate)
			if newRemaining < 0 {
				newRemaining = 0
			}
			loan.RemainingAmount = newRemaining

			if err := config.DB.Save(&loan).Error; err != nil {
				share.RespondError(c, http.StatusInternalServerError, "មិនអាចធ្វើបច្ចុប្បន្នភាពប្រាក់កម្ចីបានទេ")
				return
			}

			// Insert receive record
			newReceive := models.Recieve{
				LoanID:       input.LoanID,
				BranchID:     input.BranchId,
				RecieveDate:  time.Now(),
				TotalRecieve: float64(input.LoanDeduction),
				PayrollID:    p.ID,
				RecieveByID:  userid,
			}

			if err := config.DB.Create(&newReceive).Error; err != nil {
				share.RespondError(c, http.StatusInternalServerError, "មិនអាចរក្សាទុកទិន្នន័យទទួលប្រាក់បានទេ")
				return
			}
		}
	}

	share.ResponeSuccess(c, 200, "Payroll Created Successfully")
}

func GetPayRoll(c *gin.Context) {
	var payroll []models.PayResponse
	var user models.User
	year := c.Query("year")
	month := c.Query("month")
	BranchID := c.Query("branch_id")
	CurrencyID := c.Query("currency_id")

	userID, ok := helper.GetUserID(c)
	if !ok {
		share.RespondError(c, http.StatusUnauthorized, "Please Login")
		return
	}

	if err := config.DB.First(&user, userID).Error; err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	db := config.DB.Table("payrolls").Select(`
		payrolls.id AS id,
		s.id AS salary_id,
		es.id AS employee_shift_id,
		s.base_salary * exchange_rates.rate AS base_salary,
		s.worked_day AS worked_day,
		s.daily_rate * exchange_rates.rate AS daily_rate,
		e.id AS employee_id,
		e.name_en AS name_en,
		e.name_kh AS name_kh,
		e.gender AS gender,
		r.id AS role_id,
		r.display_name AS role_name,
		sf.id AS shift_id,
		sf.name AS shift_name,
		sf.start_time AS start_time,
		sf.end_time AS end_time,
		payrolls.payroll_year AS payroll_year,
		payrolls.payrollmonth AS payrollmonth,
		payrolls.grosssalary AS grosssalary,
		payrolls.total_late AS total_late,
		payrolls.latepenalty AS latepenalty,
		payrolls.total_earlyexit AS total_earlyexit,
		payrolls.totalexitpenalty AS totalexitpenalty,
		payrolls.leave_with_permission AS leave_with_permission,
		payrolls.penalty_leave_with_permission AS penalty_leave_with_permission,
		payrolls.leave_without_permission AS leave_without_permission,
		payrolls.penalty_leave_without_permission AS penalty_leave_without_permission,
		payrolls.leave_weekend AS leave_weekend,
		payrolls.penalty_leave_weekend AS penalty_leave_weekend,
		payrolls.loanDeduction AS loanDeduction,
		payrolls.is_attendance_bonus AS is_attendance_bonus,
		payrolls.bonus_type AS bonus_type,
		payrolls.bonus_amount AS bonus_amount,
		payrolls.totalDeductions AS totalDeductions,
		payrolls.netsalary AS netsalary,
		payrolls.status AS status,
		b.id AS branch_id,
		b.name AS branch_name,
		c.id AS currency_id,
		c.code AS currency_code,
		c.symbol AS currency_symbol,
		c.name AS currency_name,
		t.id AS base_currency_id,
		t.code AS base_currency_code,
		t.symbol AS base_currency_symbol,
		payrolls.exchange_rates AS exchange_rate
	`).
		Joins("INNER JOIN salaries s ON s.id = payrolls.salary_id").
		Joins("INNER JOIN employee_shifts es ON es.id = s.employee_shift_id").
		Joins("INNER JOIN employees e ON e.id = es.employee_id AND es.is_active = 1").
		Joins("INNER JOIN roles r ON r.id = e.role_id").
		Joins("INNER JOIN shifts sf ON sf.id = es.shift_id").
		Joins("INNER JOIN branches b ON b.id = payrolls.branch_id").
		Joins("INNER JOIN currencies c ON c.id = payrolls.currency_id").
		Joins("INNER JOIN currencies AS t ON t.id = s.currency_id").
		Joins("INNER JOIN currency_pairs ON currency_pairs.base_currency_id = s.currency_id AND currency_pairs.target_currency_id = payrolls.currency_id").
		Joins("INNER JOIN exchange_rates ON exchange_rates.pair_id = currency_pairs.id")

	if year != "" {
		db = db.Where("payrolls.year =?", year)
	}
	if month != "" {
		db = db.Where("payrolls.payrollmonth =?", month)
	}
	if BranchID != "" {
		db = db.Where("payrolls.branch_id =?", BranchID)
	}
	if CurrencyID != "" {
		db = db.Where("payrolls.currency_id =?", CurrencyID)
	}
	db = db.Where("( ? IN (1,4,7)) OR e.id =?", user.RoleID, user.EmployeeID)
	db = db.Order("id desc")
	if err := db.Scan(&payroll).Error; err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, payroll)
}

// func DeletePayroll(c *gin.Context) {
// 	ID := c.Param("id")

// 	var payroll models.Payroll
// 	var currencypair models.CurrencyPair
// 	var exchangerate models.ExchangeRate
// 	tx := config.DB.Begin()

// 	// 1. Find payroll
// 	if err := tx.First(&payroll, ID).Error; err != nil {
// 		tx.Rollback()
// 		share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យមិនឃើញ")
// 		return
// 	}

// 	// 2. If payroll has loan deduction
// 	if payroll.LoanDeduction > 0 {
// 		var recieve models.Recieve

// 		// Find receive record by payroll ID
// 		if err := tx.Where("payroll_id = ?", payroll.ID).First(&recieve).Error; err != nil {
// 			tx.Rollback()
// 			share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យទទួលមិនឃើញ")
// 			return
// 		}

// 		// Restore loan remaining amount
// 		var loan models.Loan
// 		if err := tx.First(&loan, recieve.LoanID).Error; err == nil {
// 			if err := tx.First(&currencypair).Where("base_currency_id =? AND target_currency_id =?", payroll.CurrencyID, loan.CurrencyID).Error; err != nil {
// 				share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យទទួលមិនឃើញ")
// 				return
// 			}
// 			if err := tx.First(&exchangerate).Where("pair_id =?", currencypair.ID).Error; err != nil {
// 				share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យទទួលមិនឃើញ")
// 				return
// 			}
// 			loan.Status = 1
// 			loan.RemainingAmount += (payroll.LoanDeduction * exchangerate.Rate)

// 			if err := tx.Save(&loan).Error; err != nil {
// 				tx.Rollback()
// 				share.RespondError(c, http.StatusInternalServerError, "មិនអាចធ្វើបច្ចុប្បន្នភាពប្រាក់កម្ចីបាន")
// 				return
// 			}
// 		}
// 	}

// 	// 3. Delete receive record
// 	if err := tx.Where("payroll_id = ?", payroll.ID).Delete(&models.Recieve{}).Error; err != nil {
// 		tx.Rollback()
// 		share.RespondError(c, http.StatusInternalServerError, "មិនអាចលុបទទួលប្រាក់បាន")
// 		return
// 	}

// 	// 4. Delete payroll
// 	if err := tx.Delete(&models.Payroll{}, ID).Error; err != nil {
// 		tx.Rollback()
// 		share.RespondError(c, http.StatusInternalServerError, "មិនអាចលុបប្រាក់ខែបាន")
// 		return
// 	}

// 	tx.Commit()
// 	share.ResponeSuccess(c, 200, "លុបបានជោគជ័យ")
// }

func DeletePayroll(c *gin.Context) {
	ID := c.Param("id")

	var payroll models.Payroll
	var currencypair models.CurrencyPair
	var exchangerate models.ExchangeRate

	tx := config.DB.Begin()

	// 1. Find payroll
	if err := tx.First(&payroll, ID).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យមិនឃើញ")
		return
	}

	// 2. Handle loan restore
	if payroll.LoanDeduction > 0 {
		var recieve models.Recieve

		if err := tx.Where("payroll_id = ?", payroll.ID).First(&recieve).Error; err != nil {
			tx.Rollback()
			share.RespondError(c, http.StatusNotFound, "រកទិន្ន័យទទួលមិនឃើញ")
			return
		}

		var loan models.Loan
		if err := tx.First(&loan, recieve.LoanID).Error; err == nil {

			// Find currency pair
			if err := tx.Where(
				"base_currency_id = ? AND target_currency_id = ?",
				payroll.CurrencyID, loan.CurrencyID,
			).First(&currencypair).Error; err != nil {
				tx.Rollback()
				share.RespondError(c, http.StatusNotFound, "រក currency pair មិនឃើញ")
				return
			}

			// Find exchange rate
			if err := tx.Where("pair_id = ?", currencypair.ID).
				First(&exchangerate).Error; err != nil {
				tx.Rollback()
				share.RespondError(c, http.StatusNotFound, "រក exchange rate មិនឃើញ")
				return
			}

			// Restore loan amount
			restoredAmount := payroll.LoanDeduction * exchangerate.Rate
			loan.RemainingAmount += restoredAmount

			// Update loan status
			if loan.RemainingAmount > 0 {
				loan.Status = 1 // active
			} else {
				loan.Status = 2 // completed
			}

			if err := tx.Save(&loan).Error; err != nil {
				tx.Rollback()
				share.RespondError(c, http.StatusInternalServerError, "មិនអាចធ្វើបច្ចុប្បន្នភាពប្រាក់កម្ចីបាន")
				return
			}
		}
	}

	// 3. Delete receive records
	if err := tx.Where("payroll_id = ?", payroll.ID).
		Delete(&models.Recieve{}).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "មិនអាចលុបទទួលប្រាក់បាន")
		return
	}

	// 4. Delete payroll
	if err := tx.Delete(&models.Payroll{}, ID).Error; err != nil {
		tx.Rollback()
		share.RespondError(c, http.StatusInternalServerError, "មិនអាចលុបប្រាក់ខែបាន")
		return
	}

	tx.Commit()
	share.ResponeSuccess(c, 200, "លុបបានជោគជ័យ")
}
