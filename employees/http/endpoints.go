package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"algogrit.com/emp-server/entities"
	"github.com/go-chi/chi/v5"
)

func (h *EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := h.svcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.svcV1.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

func (h *EmployeeHandler) ShowV1(w http.ResponseWriter, req *http.Request) {
	empID, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	emp, err := h.svcV1.Show(empID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
