package routes

import (
	"HRbackend/constant/permission"
	"HRbackend/controller"
	AttendanceController "HRbackend/controller/Attendance"
	AuthController "HRbackend/controller/Auth"
	BranchController "HRbackend/controller/Branch"
	CommunceController "HRbackend/controller/Communce"
	CurrencyController "HRbackend/controller/Currency"
	CurrencyPairController "HRbackend/controller/Currency_Pair"
	DistrictController "HRbackend/controller/District"
	EmployeeController "HRbackend/controller/Employee"
	ProvinceController "HRbackend/controller/Province"
	VillageController "HRbackend/controller/Village"
	"HRbackend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	branchcontroller := BranchController.NewBranchController()
	currencycontroller := CurrencyController.NewCurrencyController()
	authcontroller := AuthController.NewAuthController()
	CommunceController := CommunceController.NewCommunceController()
	DistrictController := DistrictController.NewDistrictController()
	ProvinceController := ProvinceController.NewProvinceController()
	villageController := VillageController.NewVillageController()
	attendanceController := AttendanceController.NewAttendanceController()
	currencypairController := CurrencyPairController.NewCurrencyPairController()
	employeeController := EmployeeController.NewEmployeeController()
	// Public routes
	r.POST("/login", middleware.RateLimiterMiddleware(), authcontroller.Login)

	r.Static("/profileimage", "./public/profileimage")
	r.Static("/qrcodeimage", "./public/qrcodeimage")
	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// user
		auth.POST("//register", middleware.PermissionMiddleware("add-user"), controller.Register)

		auth.PUT("/chnagestatususer/:id", middleware.PermissionMiddleware("change-status-user"), controller.ChangeStatusUser)

		auth.PUT("/updateuser/:id", middleware.PermissionMiddleware("update-user"), controller.UpdateUser)

		auth.GET("/viewuser", middleware.PermissionMiddleware("view-user"), controller.Getuserv2)

		auth.GET("/viewuser1", middleware.PermissionMiddleware("view-user"), controller.GetUser)

		auth.PUT("/changepassword/:id", middleware.PermissionMiddleware("change-password"), controller.ChangePassword)
		// shift
		auth.POST("shift", middleware.PermissionMiddleware("add-shift"), controller.CreateShift)

		auth.PUT("shift/:id", middleware.PermissionMiddleware("edit-shift"), controller.UpdateShift)

		auth.GET("viewshift", middleware.PermissionMiddleware("view-shift"), controller.GetShift)

		auth.GET("viewshiftbybranchid/:id", middleware.PermissionMiddleware("view-shift"), controller.GetShiftByBranchID)

		auth.PUT("changestatusshift/:id", middleware.PermissionMiddleware("edit-shift"), controller.ChangeStatusShift)

		// employee
		auth.GET("viewemployee", middleware.PermissionMiddleware(permission.ViewEmployee), employeeController.GetEmployees)

		auth.PUT("editemployee/:id", middleware.PermissionMiddleware(permission.EditEmployee), employeeController.UpdateEmployee)

		auth.PUT("changestatusemployee/:id", middleware.PermissionMiddleware(permission.ChangeStatusEmployee), employeeController.ChangeStatusEmployee)

		auth.PUT("promoteemployee/:id", middleware.PermissionMiddleware(permission.EditEmployee), employeeController.PromoteEmployee)

		// province
		auth.GET("viewprovince", middleware.PermissionMiddleware(permission.ViewProvince), ProvinceController.Get)

		// district
		auth.GET("viewdistrict/:id", middleware.PermissionMiddleware(permission.ViewDistrict), DistrictController.GetDistrict)

		// communce
		auth.GET("viewcommunce/:id", middleware.PermissionMiddleware(permission.ViewCommunce), CommunceController.GetCommunes)

		// village
		auth.GET("viewvillage/:id", middleware.PermissionMiddleware(permission.ViewVillage), villageController.GetVillage)

		// branch
		auth.POST("addbranch", middleware.PermissionMiddleware(permission.AddBranch), branchcontroller.Create)

		auth.GET("viewbranch", middleware.PermissionMiddleware(permission.ViewBranch), branchcontroller.Get)

		auth.PUT("editbranch/:id", middleware.PermissionMiddleware(permission.EditBranch), branchcontroller.Update)

		auth.PUT("/changestatusbranch/:id", middleware.PermissionMiddleware(permission.ChangeStatusBranch), branchcontroller.ChangeStatus)

		// role
		auth.GET("viewrole", middleware.PermissionMiddleware("view-role"), controller.GetRole)

		auth.POST("addrole", middleware.PermissionMiddleware("add-role"), controller.CreateRole)

		auth.PUT("changestatusrole/:id", middleware.PermissionMiddleware("change-status-role"), controller.ChangeStatusRole)

		auth.PUT("editrole/:id", middleware.PermissionMiddleware("edit-role"), controller.UpdateRole)

		// rolepermission
		auth.POST("addpermissiontorole", middleware.PermissionMiddleware("add-role"), controller.CreateRolePermissions)

		auth.DELETE("deletepermissionrole", middleware.PermissionMiddleware("add-role"), controller.DeleteRolePermission)

		auth.GET("viewrolehaspermission/:id", middleware.PermissionMiddleware("view-role-permission"), controller.GetRolePermission)

		// employeeshift
		auth.PUT("editemployeeshift/:id", middleware.PermissionMiddleware("edit-employee-shift"), controller.UpdateEmployeeShift)

		auth.POST("employee-shift/:employeeshiftid/salary/:salaryid", middleware.PermissionMiddleware("add-employee-shift"), controller.CreateEmployeeShift)

		auth.GET("viewemployeeshift", middleware.PermissionMiddleware("view-shift"), controller.GetEmployeeShiftByuserLogin)

		// salary
		auth.PUT("editsalary/:id", middleware.PermissionMiddleware("edit-salary"), controller.UpdateSalary)

		// attendance
		auth.POST("checkin", middleware.PermissionMiddleware(permission.CheckIN), attendanceController.CheckIn)

		auth.POST("/checkout", middleware.PermissionMiddleware(permission.CheckOut), attendanceController.CheckOut)

		auth.GET("viewattendance", middleware.PermissionMiddleware("view-attendance"), controller.GetAttendanceLog)

		// loan
		auth.POST("addloan", middleware.PermissionMiddleware("add-loan"), controller.CreateLoan)

		auth.GET("viewloan", middleware.PermissionMiddleware("view-loan"), controller.GetLoan)

		auth.PUT("editloan/:id", middleware.PermissionMiddleware("edit-loan"), controller.UpdateLoan)

		// leave
		auth.POST("addleave", middleware.PermissionMiddleware("add-leave"), controller.CreateLeave)

		auth.GET("viewleave", middleware.PermissionMiddleware("view-leave"), controller.GetLeave)

		auth.PUT("changestatusleave/:id", middleware.PermissionMiddleware("edit-leave"), controller.ChangeStatusLeave)

		auth.PUT("updateleave/:id", middleware.PermissionMiddleware("edit-leave"), controller.UpdateLeave)

		// Payroll
		auth.GET("viewpayroll", middleware.PermissionMiddleware("view-payroll"), controller.GetSummaryForPayroll)

		auth.POST("addpayroll", middleware.PermissionMiddleware("add-payroll"), controller.CreatePayroll)

		auth.GET("getpayroll", middleware.PermissionMiddleware("view-payroll"), controller.GetPayRoll)

		auth.DELETE("deletepayroll/:id", middleware.PermissionMiddleware("delete-payroll"), controller.DeletePayroll)

		// Currency
		auth.POST("addcurrency", middleware.PermissionMiddleware(permission.AddCurrency), currencycontroller.Create)

		auth.GET("viewcurrency", middleware.PermissionMiddleware(permission.ViewCurrency), currencycontroller.Get)

		auth.PUT("updatecurrency/:id", middleware.PermissionMiddleware(permission.EditCurrency), currencycontroller.Update)

		auth.PUT("changestatuscurrency/:id", middleware.PermissionMiddleware(permission.ChangeStatusCurrency), currencycontroller.ChangeStatus)

		// CurrencyPair
		auth.POST("addcurrencypair", middleware.PermissionMiddleware(permission.AddCurrencyPair), currencypairController.CreateCurrencyPair)

		auth.GET("viewcurrencypair", middleware.PermissionMiddleware(permission.ViewCurrencyPair), currencypairController.GetCurrencypair)

		auth.PUT("updatecurrencypaire/:id", middleware.PermissionMiddleware(permission.EditCurrencyPair), currencypairController.UpdateCurrencyPaire)

		auth.PUT("changestatuscurrencypaire/:id", middleware.PermissionMiddleware(permission.ChangeStatusCurrencyPair), currencypairController.ChangeStatusCurrencyPair)

		// Exchange Rate
		auth.POST("addexchangerate", middleware.PermissionMiddleware("add-exchange-rate"), controller.CreateExchangeRate)

		auth.GET("viewexchangerate", middleware.PermissionMiddleware("view-exchange-rate"), controller.GetExchangeRate)

		auth.PUT("updatexchangerate/:id", middleware.PermissionMiddleware("edit-exchange-rate"), controller.UpdateExchageRate)

		auth.PUT("changestatusexchangerate/:id", middleware.PermissionMiddleware("change-status-exchange-rate"), controller.ChangeStatusExchangeRate)

		// DeductType
		auth.POST("adddeductype", middleware.PermissionMiddleware("add-deducttype"), controller.CreateDeductType)

		auth.GET("viewdeductype", middleware.PermissionMiddleware("view-deducttype"), controller.GetDeductType)

		auth.PUT("updatedeductype/:id", middleware.PermissionMiddleware("edit-deducttype"), controller.UpdateDeductType)

		auth.PUT("changestatusdedustype/:id", middleware.PermissionMiddleware("change-staus-deducttype"), controller.ChnageStatusDeductType)

		// LeaveType
		auth.POST("addleavetype", middleware.PermissionMiddleware("add-leave-type"), controller.CreateLeaveType)

		auth.GET("viewleavetype", middleware.PermissionMiddleware("view-leave-type"), controller.GetLeaveType)

		auth.PUT("updateleavetype/:id", middleware.PermissionMiddleware("update-leave-type"), controller.UpdateLeaveType)

		auth.PUT("changestatusleavetype/:id", middleware.PermissionMiddleware("change-status-leave-type"), controller.ChangeStatusLeaveType)

		// Part
		auth.POST("addpart", middleware.PermissionMiddleware("add-part"), controller.CreatePart)

		auth.GET("viewpart", middleware.PermissionMiddleware("view-part"), controller.Getpart)

	}
}
