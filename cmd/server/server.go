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
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"projectID"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "projectID": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 10001},
	{2, "Vignesh", "SRE", 10002},
	{3, "Kavitha", "Cloud", 20001},
}

func EmployeeIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintln(w, employees)

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1
	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func EmployeeShowHandler(w http.ResponseWriter, req *http.Request) {
	empID := chi.URLParam(req, "id")   // Type: string
	empIdx, err := strconv.Atoi(empID) // Type: int

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	emp := employees[empIdx-1]

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
