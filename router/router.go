package router

import (
	"net/http"
	"next-learn-go/controller"

	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	uc controller.IUserController,
	ic controller.IInvoiceController,
	rc controller.IRevenueController,
	cc controller.ICustomerController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PATCH", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)

	tv := e.Group("/token/verify")
	tv.Use(jwtMiddleware)
	tv.POST("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	tr := e.Group("/token/refresh")
	tr.POST("", uc.RefreshToken)


	i := e.Group("/invoices")
	i.Use(jwtMiddleware)
	i.GET("/latest", ic.GetLatestInvoices)
	i.GET("/filtered", ic.GetFilteredInvoices)
	i.GET("/count", ic.GetInvoiceCount)
	i.GET("/statusCount", ic.GetInvoiceStatusCount)
	i.GET("/pages", ic.GetInvoicesPages)
	i.GET("/:invoiceId", ic.GetInvoiceById)
	i.POST("", ic.CreateInvoice)
	i.PATCH("/:invoiceId", ic.UpdateInvoice)
	i.DELETE("/:invoiceId", ic.DeleteInvoice)

	r := e.Group("/revenues")
	r.Use(jwtMiddleware)
	r.GET("", rc.GetAllRevenues)

	c := e.Group("/customers")
	c.Use(jwtMiddleware)
	c.GET("", cc.GetAllCustomers)
	c.GET("/filtered", cc.GetFilteredCustomers)
	c.GET("/count", cc.GetCustomerCount)

	u := e.Group("/user")
	u.Use(jwtMiddleware)
	u.GET("", uc.GetUserById)
	u.GET("/email", uc.GetUserByEmail)
	return e
}
