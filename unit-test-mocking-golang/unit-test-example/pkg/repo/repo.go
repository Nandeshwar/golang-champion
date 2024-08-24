package repo

import (
	"unit-test-example/pkg/model/dto"
)

type Repo interface {
	EmployeeInfo() []dto.EmployeeDTO
}

type MySql struct {
}

func New() Repo {
	return &MySql{}
}

func (m *MySql) EmployeeInfo() []dto.EmployeeDTO {
	employeeList := []dto.EmployeeDTO{
		{ID: 1, Name: "Hari", Salary: 1000},
		{ID: 2, Name: "Shyam", Salary: 2000},
		{ID: 3, Name: "Ram", Salary: 1500},
	}

	return employeeList
}
