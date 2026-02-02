package domain_test

import (
	"errors"
	"testing"

	"edm/domain"
)

func TestAddEmployee(t *testing.T) {
	manager := domain.Manager{}

	_ = manager.AddEmployee(domain.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	_ = manager.AddEmployee(domain.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})

	if len(manager.Employees) != 2 {
		t.Errorf("Expected 2 employees, got %d", len(manager.Employees))
	}

	err := manager.AddEmployee(domain.Employee{ID: 1, Name: "Derek", Age: 40, Salary: 80000})
	if !errors.As(err, &domain.ErrEmployeeAlreadyExist{}) {
		t.Errorf("wrong error: %v", err)
	}
}
