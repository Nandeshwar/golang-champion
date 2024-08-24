package dto

import (
	"fmt"
)

type EmployeeDTO struct {
	ID     int
	Name   string
	Salary float32
}

func (e *EmployeeDTO) Info() {
	str := fmt.Sprintf("ID=%d, Name=%s, Salary=%.2f", e.ID, e.Name, e.Salary)
	fmt.Println(str)
}
