package main

import (
	"net/http"

	"github.com/felipehfs/appointment-app/controller"
	"github.com/felipehfs/appointment-app/infra"
	"github.com/felipehfs/appointment-app/repository"
)

func main() {
	db, err := infra.CreateDatabase()
	if err != nil {
		panic(err)
	}

	db.Ping()

	mux := http.NewServeMux()

	customerRepository := repository.NewCustomerRepository(db)
	appointmentRepository := repository.NewAppointmentRepository(db)

	customerController := controller.NewCustomerController(customerRepository)
	appointmentController := controller.NewAppointmentController(appointmentRepository)

	withRoutes := infra.RegisterAllRoutes(mux)
	withRoutes(customerController, appointmentController)

	http.ListenAndServe(":3000", mux)
}
