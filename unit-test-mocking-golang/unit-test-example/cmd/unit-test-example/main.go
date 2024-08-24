package main

import (
	"unit-test-example/pkg/repo"
	"unit-test-example/pkg/service"
)

func main() {
	repoObj := repo.New()
	empService := service.New(repoObj)
	empList := empService.ListIncreasedSalaryEmpAboveThreshold(2000)

	for _, emp := range empList {
		emp.Info()
	}

}
