package main

import (
	"next-learn-go/controller"
	"next-learn-go/db"
	"next-learn-go/repository"
	"next-learn-go/router"
	"next-learn-go/usecase"
	"next-learn-go/validator"
	"os"
)

func main() {

	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	invoiceValidator := validator.NewInvoiceValidator()

	userRepository := repository.NewUserRepository(db)
	invoiceRepository := repository.NewInvoiceRepository(db)
	revenueRepository := repository.NewRevenueRepository(db)
	customerRepository := repository.NewCustomerRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	invoiceUsecase := usecase.NewInvoiceUsecase(invoiceRepository, invoiceValidator)
	revenueUsecase := usecase.NewRevenueUsecase(revenueRepository)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)

	userController := controller.NewUserController(userUsecase)
	invoiceController := controller.NewInvoiceController(invoiceUsecase)
	revenueController := controller.NewRevenueController(revenueUsecase)
	customerController := controller.NewCustomerController(customerUsecase)

	e := router.NewRouter(userController, invoiceController, revenueController, customerController)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
