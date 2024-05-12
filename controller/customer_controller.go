package controller

import (
	"net/http"
	"next-learn-go/usecase"

	"github.com/labstack/echo/v4"
)

type ICustomerController interface {
	GetAllCustomers(c echo.Context) error
	GetFilteredCustomers(c echo.Context) error
	GetCustomerCount(c echo.Context) error
}

type customerController struct {
	cu usecase.ICustomerUsecase
}

func NewCustomerController(cu usecase.ICustomerUsecase) ICustomerController {
	return &customerController{cu}
}

func (cc *customerController) GetAllCustomers(c echo.Context) error {
	customers, err := cc.cu.GetAllCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

func (cc *customerController) GetFilteredCustomers(c echo.Context) error {
	query := c.QueryParams().Get("query")
	customers, err := cc.cu.GetFilteredCustomers(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, customers)
}

func (cc *customerController) GetCustomerCount(c echo.Context) error {
	count, err := cc.cu.GetCustomerCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, count)
}
