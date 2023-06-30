package repository

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"algogrit.com/emp-server/entities"
)

type sqlRepo struct {
	*sql.DB
}

func (repo *sqlRepo) ListAll() ([]entities.Employee, error) {
	rows, err := repo.DB.Query("SELECT * FROM employees;")

	if err != nil {
		log.Println("Unable to fetch:", err)
		return nil, err
	}

	var output []entities.Employee

	for rows.Next() {
		var retrievedEmp entities.Employee
		err = rows.Scan(&retrievedEmp.ID, &retrievedEmp.Name, &retrievedEmp.Department, &retrievedEmp.ProjectID)

		if err != nil {
			log.Println("Unable to scan:", err)
			return nil, err
		}

		output = append(output, retrievedEmp)
	}

	return output, nil
}

func (repo *sqlRepo) GetByID(id int) (*entities.Employee, error) {
	rows, err := repo.DB.Query("SELECT * FROM employees WHERE id = ?;", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var retrievedEmp entities.Employee
		err = rows.Scan(&retrievedEmp.ID, &retrievedEmp.Name, &retrievedEmp.Department, &retrievedEmp.ProjectID)

		if err != nil {
			log.Println("Unable to scan:", err)
			return nil, err
		}

		return &retrievedEmp, nil
	}

	return nil, errors.New("record not found!")
}

func (repo *sqlRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	return nil, nil
	// newEmp.ID = len(repo.emps) + 1
	// repo.emps = append(repo.emps, newEmp)

	// return &newEmp, nil
}

func NewSQL(db *sql.DB) EmployeeRepository {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS employees (id integer PRIMARY KEY, name text, department text, project_id integer);")

	if err != nil {
		log.Fatalln("Unable to create table:", err)
	}

	return &sqlRepo{db}
}
