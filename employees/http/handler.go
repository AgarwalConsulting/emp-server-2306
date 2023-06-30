package http

import (
	"algogrit.com/emp-server/employees/service"
	"github.com/go-chi/chi/v5"
)

type EmployeeHandler struct {
	*chi.Mux
	svcV1 service.EmployeeService
}

func (h *EmployeeHandler) SetupRoutes(r *chi.Mux) {
	r.Get("/v1/employees", h.IndexV1)
	r.Post("/v1/employees", h.CreateV1)
	r.HandleFunc("/v1/employees/{id}", h.ShowV1)

	h.Mux = r
}

func NewHandler(svcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{svcV1: svcV1}
	r := chi.NewRouter()

	h.SetupRoutes(r)

	return h
}
