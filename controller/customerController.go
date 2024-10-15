package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipehfs/appointment-app/dto"
	"github.com/felipehfs/appointment-app/repository"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	CustomerRepo repository.CustomerRepository
}

func NewCustomerController(customerRepository repository.CustomerRepository) CustomerController {
	return CustomerController{
		CustomerRepo: customerRepository,
	}
}

func (cc CustomerController) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /customers", cc.Insert)
	mux.HandleFunc("GET /customers", cc.Select)
	mux.HandleFunc("GET /customers/{id}", cc.FindById)
	mux.HandleFunc("PUT /customers/{id}", cc.Update)
	mux.HandleFunc("DELETE /customers/{id}", cc.Delete)
}

func (cc CustomerController) Insert(writer http.ResponseWriter, req *http.Request) {
	var customer dto.CreateCustomer

	if err := json.NewDecoder(req.Body).Decode(&customer); err != nil {
		http.Error(writer, fmt.Sprintf("Invalid JSON: %s", err), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err := validate.Struct(customer)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(writer, fmt.Sprintf("validation errors: %s", errors), http.StatusBadRequest)
		return
	}

	err = cc.CustomerRepo.Insert(customer)

	if err != nil {
		http.Error(writer, fmt.Sprintf("A error occured during to save: %s", err), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User created successfully")
}

func (cc CustomerController) Select(w http.ResponseWriter, r *http.Request) {
	data, err := cc.CustomerRepo.Select()

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occured during to read data: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (cc CustomerController) FindById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	data, err := cc.CustomerRepo.FindById(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occured during to read data: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (cc CustomerController) Update(writer http.ResponseWriter, req *http.Request) {
	var customer dto.UpdateCustomer
	id, _ := strconv.Atoi(req.PathValue("id"))

	if err := json.NewDecoder(req.Body).Decode(&customer); err != nil {
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err := validate.Struct(customer)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(writer, fmt.Sprintf("validation errors: %s", errors), http.StatusBadRequest)
		return
	}

	err = cc.CustomerRepo.Update(id, customer)

	if err != nil {
		http.Error(writer, fmt.Sprintf("A error occured during to save: %s", err), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "User updated successfully")
}

func (cc CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	err := cc.CustomerRepo.Remove(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occured during to read data: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
