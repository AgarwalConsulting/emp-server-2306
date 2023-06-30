package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

var empRepo = repository.NewInMem()
var empSvcV1 = service.NewV1(empRepo)

func EmployeeIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := empSvcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empSvcV1.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

func EmployeeShowHandler(w http.ResponseWriter, req *http.Request) {
	empID, err := strconv.Atoi(chi.URLParam(req, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	emp, err := empSvcV1.Show(empID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else {
// 		EmployeeIndexHandler(w, req)
// 	}
// }

func LoggingMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		// if //
		next.ServeHTTP(w, req)
		// else

		dur := time.Since(startTime)

		log.Infof("%s %s took %v", req.Method, req.URL, dur)
	}

	return http.HandlerFunc(h)
}

func main() {
	// r := http.NewServeMux()
	r := chi.NewRouter()

	// r.Use(LoggingMiddleware)
	r.Use(middleware.DefaultLogger)

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeeIndexHandler)
	// r.HandleFunc("/employees", EmployeeCreateHandler)
	r.Get("/employees", EmployeeIndexHandler)
	r.Post("/employees", EmployeeCreateHandler)
	r.HandleFunc("/employees/{id}", EmployeeShowHandler)

	// http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
	http.ListenAndServe("localhost:8000", r)
}
