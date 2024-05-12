package controller

import (
	"net/http"
	"next-learn-go/usecase"

	"github.com/labstack/echo/v4"
)

type IRevenueController interface {
	GetAllRevenues(c echo.Context) error
}

type revenueController struct {
	ru usecase.IRevenueUsecase
}

func NewRevenueController(ru usecase.IRevenueUsecase) IRevenueController {
	return &revenueController{ru}
}

func (rc *revenueController) GetAllRevenues(c echo.Context) error {
	revenues, err := rc.ru.GetAllRevenues()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, revenues)
}
