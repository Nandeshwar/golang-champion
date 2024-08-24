package repo

import (
	"github.com/stretchr/testify/mock"
	"unit-test-example/pkg/model/dto"
)

type RepoTest struct {
	mock.Mock
}

func (r *RepoTest) EmployeeInfo() []dto.EmployeeDTO {
	args := r.Called()
	//return args.Get(0).([]dto.EmployeeDTO), args.Error(1)
	return args.Get(0).([]dto.EmployeeDTO)
}
