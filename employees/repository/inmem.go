package repository

import (
	"errors"

	"algogrit.com/emp-server/entities"
)

type inmemRepo struct {
	emps []entities.Employee
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	return repo.emps, nil
}

func (repo *inmemRepo) GetByID(id int) (*entities.Employee, error) {
	if id > len(repo.emps) {
		return nil, errors.New("invalid id")
	}

	return &repo.emps[id-1], nil
}

func (repo *inmemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.emps) + 1
	repo.emps = append(repo.emps, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	return &inmemRepo{
		emps: []entities.Employee{
			{1, "Gaurav", "LnD", 10001},
			{2, "Vignesh", "SRE", 10002},
			{3, "Kavitha", "Cloud", 20001},
		},
	}
}
