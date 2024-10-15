package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipehfs/appointment-app/dto"
	"github.com/felipehfs/appointment-app/infra"
	"github.com/felipehfs/appointment-app/repository"
	"github.com/go-playground/validator/v10"
)

type AppointmentController struct {
	Repository repository.AppointmentRepository
}

func NewAppointmentController(repository repository.AppointmentRepository) AppointmentController {
	return AppointmentController{
		Repository: repository,
	}
}

func (ac AppointmentController) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /appointments", infra.SecureRoute(ac.Insert))
	mux.HandleFunc("GET /appointments", ac.Select)
	mux.HandleFunc("GET /appointments/{id}", ac.FindById)
	mux.HandleFunc("PUT /appointments/{id}", ac.Update)
	mux.HandleFunc("DELETE /appointments/{id}", ac.Remove)
}

func (ac AppointmentController) Remove(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	err := ac.Repository.Delete(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to remove data: %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ac AppointmentController) Update(w http.ResponseWriter, r *http.Request) {
	var data dto.UpdateAppointmentDto

	id, _ := strconv.Atoi(r.PathValue("id"))

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to parse JSON: %s", err), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validation errors: %s", errors), http.StatusBadRequest)
		return
	}

	err = ac.Repository.Update(id, &data)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to retrieve data: %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ac AppointmentController) FindById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	data, err := ac.Repository.FindById(id)

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to retrieve data: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func (ac AppointmentController) Select(w http.ResponseWriter, r *http.Request) {
	data, err := ac.Repository.Select()

	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to retrieve data: %s", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}

func (ac AppointmentController) Insert(w http.ResponseWriter, r *http.Request) {
	var appointment dto.CreateAppointmentDto

	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to parse JSON: %s", err), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err := validate.Struct(appointment)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("validation errors: %s", errors), http.StatusBadRequest)
		return
	}

	err = ac.Repository.Insert(&appointment)
	if err != nil {
		http.Error(w, fmt.Sprintf("A error occurred to save appointment: %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
