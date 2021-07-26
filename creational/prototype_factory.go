package creational

import (
	"fmt"
)

type RoleType int

const (
	Developer RoleType = iota
	Manager
)

type Employee struct {
	Name         string
	Role         RoleType
	AnnualIncome int
}

func NewEmployee(role RoleType) (*Employee, error) {
	switch role {
	case Developer:
		return &Employee{"", Developer, 80000}, nil
	case Manager:
		return &Employee{"", Manager, 100000}, nil
	default:
		return nil, fmt.Errorf("role %d is not a valid role", role)
	}
}
