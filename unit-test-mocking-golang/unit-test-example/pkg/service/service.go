package service

import (
	"unit-test-example/pkg/model/dto"
	"unit-test-example/pkg/repo"
)

type EmployeeService interface {
	ListIncreasedSalaryEmpAboveThreshold(maxAmount float32) []dto.EmployeeDTO
	IncreaseSalary(salary float32, percent int) float32
}

type Employee struct {
	Repo repo.Repo
}

func New(repo repo.Repo) EmployeeService {
	return &Employee{Repo: repo}

}

func (e *Employee) IncreaseSalary(salary float32, percent int) float32 {
	return salary + salary*float32(percent)/100
}

func (e *Employee) ListIncreasedSalaryEmpAboveThreshold(maxAmount float32) []dto.EmployeeDTO {
	employeList := e.Repo.EmployeeInfo()
	var increaseSalaryEmployeeList []dto.EmployeeDTO

	for _, emp := range employeList {
		increaseSalary := e.IncreaseSalary(emp.Salary, 10)

		if increaseSalary > maxAmount {
			newEmp := dto.EmployeeDTO{
				ID:     emp.ID,
				Name:   emp.Name,
				Salary: increaseSalary,
			}

			increaseSalaryEmployeeList = append(increaseSalaryEmployeeList, newEmp)
		}

	}
	return increaseSalaryEmployeeList
}
