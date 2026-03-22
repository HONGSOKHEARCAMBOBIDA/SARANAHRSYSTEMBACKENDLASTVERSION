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

func GetSummaryForPayroll(c *gin.Context) {

	var summaries []models.SummaryforPayroll

	branchID := c.Query("branch_id")

	month := c.Query("month")

	currencyID := c.Query("currency_id")

	currentYear := time.Now().Year()

	staffName := c.Query("staff_name")

	db := config.DB.Table("employees").Select(`

				salaries.id AS salary_id,

				employee_shifts.id AS employee_shift_id,

				exchange_rates.rate AS exchange_rate,

				salaries.base_salary / er.rate * exchange_rates.rate AS base_salary,

				salaries.worked_day,

				salaries.daily_rate / er.rate  * exchange_rates.rate AS daily_rate,

				currencies.id AS currency_id,

				currencies.code AS currency_code,

				currencies.symbol AS currency_symbol,

				currencies.name AS currency_name,

				currency_pairs.id AS currency_pair_id,

				shifts.id AS shift_id,

				shifts.name AS shift_name,

				shifts.start_time,

				shifts.end_time,

				employees.id AS employee_id,

				branches.id AS branch_id,

				branches.name AS branch_name,

				employees.name_en,

				employees.name_kh,

				employees.gender,	

				employees.contact,

				roles.id AS role_id,

				roles.name AS role_name,

				employees.type,

				employees.hire_date,

				employee_profiles.date_of_birth AS dob,

				COALESCE(att.total_late,0) AS total_late,

				COALESCE(att.penalty_late  * exchange_rates.rate ,0) AS penaltylate,

				COALESCE(att.total_earlyexit,0) AS total_earlyexit, 

				COALESCE(att.penalty_earlyexit  * exchange_rates.rate ,0) AS totalexitpenalty,

				COALESCE(att.attendancecount,0) AS attendancecount,

				COALESCE(lv.leave_with_permission,0) AS leave_with_permission,

				COALESCE(lv.leave_with_permission * salaries.daily_rate / er.rate  * exchange_rates.rate,0) AS penalty_leave_with_permission,

				COALESCE(lv.leave_without_permission,0) AS leave_without_permission,

				COALESCE(
						CASE 
							WHEN lv.leave_without_permission * salaries.daily_rate * exchange_rates.rate = 0.0 

								THEN 0

							ELSE 
									CEIL(salaries.daily_rate / er.rate + 0.000001) * exchange_rates.rate * lv.leave_without_permission
									
						END,
				0) AS penalty_leave_without_permission,

				COALESCE(lv.leave_weekend,0) AS leave_weekend,

				COALESCE(lv.leave_weekend * 10  * exchange_rates.rate ,0) AS penalty_leave_weekend,

				COALESCE(loans.id,0) AS loan_id,

				COALESCE(loans.loan_amount * erl.rate,0) AS loan_amount,

				COALESCE(loans.remaining_balance * erl.rate,0) AS remaining_balance,

				(COALESCE(att.penalty_late  * exchange_rates.rate, 0)

  					+ COALESCE(att.penalty_earlyexit  * exchange_rates.rate, 0)

  					+ (COALESCE(lv.leave_with_permission, 0) * COALESCE(salaries.daily_rate, 0) / COALESCE(er.rate,0)* COALESCE(exchange_rates.rate, 0))

  					+ COALESCE(

        				CASE 

            				WHEN lv.leave_without_permission * salaries.daily_rate * exchange_rates.rate = 0.0 

                				THEN 0

            				ELSE 

							CEIL(salaries.daily_rate / er.rate + 0.000001) * exchange_rates.rate * lv.leave_without_permission
              
        				END,0)

  					+ (COALESCE(lv.leave_weekend, 0) * 10  * exchange_rates.rate)

				) AS totalDeductions,

				COALESCE(salaries.daily_rate * att.attendancecount / er.rate * exchange_rates.rate) AS notdeduction,

				COALESCE( CASE WHEN att.attendancecount >= salaries.worked_day THEN 1 ELSE 0  END, 0) AS is_bonus_attendance,

				(COALESCE(salaries.daily_rate,0) * COALESCE(att.attendancecount,0) / COALESCE(er.rate,0) * COALESCE(exchange_rates.rate,0) - 

				(COALESCE(att.penalty_late  * exchange_rates.rate,0) 

					+ COALESCE(att.penalty_earlyexit  * exchange_rates.rate,0) 

					+ (COALESCE(lv.leave_with_permission,0) * COALESCE(salaries.daily_rate,0) / COALESCE(er.rate,0) * COALESCE(exchange_rates.rate,0))+ 

				COALESCE(
				
						CASE 

							WHEN lv.leave_without_permission * salaries.daily_rate * exchange_rates.rate = 0.0 

							THEN 0

						ELSE 

							CEIL(salaries.daily_rate / er.rate + 0.000001) * exchange_rates.rate * lv.leave_without_permission
							
						END,0)
					+ 
				(COALESCE(lv.leave_weekend,0) * 10 * exchange_rates.rate ))) AS netsalary`, currencyID, currencyID, currencyID).
		Joins("INNER JOIN employee_profiles ON employee_profiles.employee_id = employees.id").
		Joins("INNER JOIN branches ON branches.id = employees.branch_id").
		Joins("INNER JOIN roles ON roles.id = employees.role_id").
		Joins("INNER JOIN employee_shifts ON employee_shifts.employee_id = employees.id AND employee_shifts.is_active = 1").
		Joins("INNER JOIN shifts ON shifts.id = employee_shifts.shift_id").
		Joins("INNER JOIN salaries ON salaries.employee_shift_id = employee_shifts.id").
		Joins("INNER JOIN currencies ON currencies.id = salaries.currency_id").
		Joins("INNER JOIN currency_pairs AS cp ON cp.base_currency_id = 2 AND cp.target_currency_id = salaries.currency_id").
		Joins("INNER JOIN exchange_rates AS er ON er.pair_id = cp.id").
		Joins("INNER JOIN currency_pairs ON currency_pairs.base_currency_id = 2 AND currency_pairs.target_currency_id =?", currencyID).
		Joins("INNER JOIN exchange_rates ON exchange_rates.pair_id = currency_pairs.id").
		Joins("LEFT JOIN loans ON loans.employee_id = employees.id").
		Joins("LEFT JOIN currency_pairs cpl ON cpl.base_currency_id = loans.currency_id AND cpl.target_currency_id =?", currencyID).
		Joins("LEFT JOIN exchange_rates AS erl ON erl.pair_id = cpl.id").
		Joins(`LEFT JOIN (

							SELECT employee_shift_id,

								SUM(is_late) AS total_late,

								SUM(is_late * 1) AS penalty_late,

								SUM(is_left_early) AS total_earlyexit,

								SUM(is_left_early * 10) AS penalty_earlyexit,

								COUNT(id) AS attendancecount

							FROM attendance_logs

							WHERE YEAR(check_date) = ?`+func() string {

			// func() string add SQL condition if month have value
			if month != "" {

				return " AND MONTH(check_date) = ?"

			}
			return ""

		}()+`
						GROUP BY employee_shift_id

						) AS att ON att.employee_shift_id = employee_shifts.id`, func() []interface{} {
			// func() []interface ផ្តល់ parameter values ទៅក្នុង query builder
			if month != "" {

				return []interface{}{currentYear, month}
			}
			return []interface{}{currentYear}

		}()...).
		Joins(`LEFT JOIN (

									SELECT employee_shift_id,

										SUM(CASE WHEN is_permission = 1 THEN leave_days ELSE 0 END) AS leave_with_permission,

										SUM(CASE WHEN is_without_permission = 1 THEN leave_days ELSE 0 END) AS leave_without_permission,

										SUM(CASE WHEN is_weekend = 1 THEN leave_days ELSE 0 END) AS leave_weekend

									FROM leaves

									WHERE leaves.status = 0 AND YEAR(start_date) = ?`+func() string {
			if month != "" {

				return " AND MONTH(start_date) = ?"

			}
			return ""

		}()+`
								GROUP BY employee_shift_id

								) AS lv ON lv.employee_shift_id = employee_shifts.id`, func() []interface{} {

			if month != "" {

				return []interface{}{currentYear, month}
			}
			return []interface{}{currentYear}

		}()...)

	// Branch filter
	if branchID != "" {

		db = db.Where("employees.branch_id = ?", branchID)

	}
	if staffName != "" {

		db = db.Where("employees.name_kh LIKE ?", "%"+staffName+"%")

	}

	// Execute query
	if err := db.Scan(&summaries).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	for i := range summaries {

		summaries[i].Dob = helper.FormatDate(summaries[i].Dob)

		summaries[i].HireDate = helper.FormatDate(summaries[i].HireDate)

		if summaries[i].Gender == 1 {

			summaries[i].GenderText = "ប្រុស"

		} else {

			summaries[i].GenderText = "ស្រី"

		}
		if summaries[i].Type == 1 {

			summaries[i].TypeText = "Full Time"

		} else {

			summaries[i].TypeText = "Part Time"

		}

	}

	share.RespondDate(c, 200, summaries)
}
