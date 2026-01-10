package controller

import (
	"HRbackend/constant/share"
	currency "HRbackend/request/Currency"
	service "HRbackend/service/Currency"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	service service.CurrencyService
}

func NewCurrencyController() CurrencyController {
	return CurrencyController{
		service: service.NewCurrencyService(),
	}
}

func (cc CurrencyController) Create(c *gin.Context) {
	var input currency.CurrencyRequestCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := cc.service.Create(input); err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Currency Created")
}

func (cc CurrencyController) Get(c *gin.Context) {
	data, err := cc.service.GetAll()
	if err != nil {
		share.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	share.RespondDate(c, http.StatusOK, data)
}

func (cc CurrencyController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		share.RespondError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	var input currency.CurrencyRequestUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := cc.service.Update(id, input); err != nil {
		share.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Currency Updated")
}

func (cc CurrencyController) ChangeStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := cc.service.ChangeStatus(id); err != nil {
		share.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	share.ResponeSuccess(c, http.StatusOK, "Change status success")
}
