package controller

import (
	"net/http"
	"next-learn-go/model"
	"next-learn-go/usecase"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type IInvoiceController interface {
	GetLatestInvoices(c echo.Context) error
	GetFilteredInvoices(c echo.Context) error
	GetInvoiceCount(c echo.Context) error
	GetInvoiceStatusCount(c echo.Context) error
	GetInvoicesPages(c echo.Context) error
	GetInvoiceById(c echo.Context) error
	CreateInvoice(c echo.Context) error
	UpdateInvoice(c echo.Context) error
	DeleteInvoice(c echo.Context) error
}

type invoiceController struct {
	iu usecase.IInvoiceUsecase
}

func NewInvoiceController(iu usecase.IInvoiceUsecase) IInvoiceController {
	return &invoiceController{iu}
}

func (ic *invoiceController) GetLatestInvoices(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 6
	}

	invoiceRes, err := ic.iu.GetLatestInvoices(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetFilteredInvoices(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	query := c.QueryParams().Get("query")

	invoiceRes, err := ic.iu.GetFilteredInvoices(query, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceCount(c echo.Context) error {
	invoiceRes, err := ic.iu.GetInvoiceCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceStatusCount(c echo.Context) error {
	pending, paid, err := ic.iu.GetInvoiceStatusCount()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]int{"pending": pending, "paid": paid})
}

func (ic *invoiceController) GetInvoicesPages(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	query := c.QueryParams().Get("query")

	invoiceRes, err := ic.iu.GetInvoicesPages(query, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) GetInvoiceById(c echo.Context) error {
	invoiceId, err := uuid.Parse(c.Param("invoiceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	invoiceRes, err := ic.iu.GetInvoiceById(invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) CreateInvoice(c echo.Context) error {

	invoice := model.Invoice{}
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	invoiceRes, err := ic.iu.CreateInvoice(invoice)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, invoiceRes)
}

func (ic *invoiceController) UpdateInvoice(c echo.Context) error {
	invoiceId, err := uuid.Parse(c.Param("invoiceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	invoice := model.Invoice{}
	if err := c.Bind(&invoice); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	invoiceRes, err := ic.iu.UpdateInvoice(invoice, invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoiceRes)
}

func (ic *invoiceController) DeleteInvoice(c echo.Context) error {
	invoiceId, err := uuid.Parse(c.Param("invoiceId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = ic.iu.DeleteInvoice(invoiceId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
