package service

import (
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

type v1 struct {
	repo repository.EmployeeRepository
}

func (svc *v1) Index() ([]entities.Employee, error) {
	return svc.repo.ListAll()
}

func (svc *v1) Show(id int) (*entities.Employee, error) {
	return svc.repo.GetByID(id)
}

func (svc *v1) Create(newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return &v1{repo}
}
