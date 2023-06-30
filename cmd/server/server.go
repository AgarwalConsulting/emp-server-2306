package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	log "github.com/sirupsen/logrus"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, req)

		dur := time.Since(startTime)

		log.Infof("%s %s took %v", req.Method, req.URL, dur)
	}

	return http.HandlerFunc(h)
}

func main() {
	// var empRepo = repository.NewInMem()
	db, err := sql.Open("mysql", "root@/employees")
	if err != nil {
		log.Fatalln("Unable to connect:", err)
	}
	defer db.Close()

	var empRepo = repository.NewSQL(db)
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	r := chi.NewRouter()

	r.Use(middleware.DefaultLogger)

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!" // Type: string

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	port := "8000"
	log.Infof("Starting server on port: %v", port)
	http.ListenAndServe(":"+port, r)
}
