package service

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"unit-test-example/pkg/model/dto"
	"unit-test-example/pkg/repo"
)

func TestService(t *testing.T) {

	convey.Convey("Test Employee Service", t, func() {

		empRepo := new(repo.RepoTest)
		empService := New(empRepo)

		convey.Convey("success testing increase salary function", func() {
			increasedSalary := empService.IncreaseSalary(100, 10)
			expectation := 110
			convey.So(increasedSalary, convey.ShouldEqual, expectation)
		})

		convey.Convey("success testing filtering emplist by max salary threshold", func() {

			empRepo.On("EmployeeInfo").Return([]dto.EmployeeDTO{
				{ID: 100, Name: "Krish", Salary: 2000},
				{ID: 200, Name: "John", Salary: 3000},
			})
			empList := empService.ListIncreasedSalaryEmpAboveThreshold(2000)
			convey.So(len(empList), convey.ShouldEqual, 2)

			empRepo.AssertExpectations(t)
		})
	})
}
