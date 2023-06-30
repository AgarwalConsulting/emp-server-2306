package repository

import "algogrit.com/emp-server/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	GetByID(id int) (*entities.Employee, error)
	Save(newEmp entities.Employee) (*entities.Employee, error)
}
