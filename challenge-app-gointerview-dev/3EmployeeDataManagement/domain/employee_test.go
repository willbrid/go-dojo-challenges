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

func TestRemoveEmployee(t *testing.T) {
	manager := domain.Manager{}

	_ = manager.AddEmployee(domain.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	_ = manager.AddEmployee(domain.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})

	manager.RemoveEmployee(1)
	if len(manager.Employees) != 1 {
		t.Errorf("Expected 1 employee after removing ID 1, got %d", len(manager.Employees))
	}

	manager.RemoveEmployee(3)
	if len(manager.Employees) != 1 {
		t.Errorf("Expected 1 employee after attempting to remove non-existing ID, got %d", len(manager.Employees))
	}
}

func TestGetAverageSalary(t *testing.T) {
	manager := domain.Manager{}

	_ = manager.AddEmployee(domain.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	_ = manager.AddEmployee(domain.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})

	var expectedAverage float64 = (70000 + 65000) / 2
	if avg := manager.GetAverageSalary(); avg != expectedAverage {
		t.Errorf("Expected average salary %f, got %f", expectedAverage, avg)
	}

	manager.Employees = []domain.Employee{}
	expectedAverage = 0
	if avg := manager.GetAverageSalary(); avg != expectedAverage {
		t.Errorf("Expected average salary %f with no employees, got %f", expectedAverage, avg)
	}
}

func TestFindEmployeeByID(t *testing.T) {
	manager := domain.Manager{}

	_ = manager.AddEmployee(domain.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	_ = manager.AddEmployee(domain.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})

	employee := manager.FindEmployeeByID(2)
	if employee == nil || employee.Name != "Bob" {
		t.Errorf("Expected to find Bob, got %+v", employee)
	}

	employee = manager.FindEmployeeByID(3)
	if employee != nil {
		t.Errorf("Expected no employee, got %+v", employee)
	}
}

func TestFindEmployeeAfterRemoval(t *testing.T) {
	manager := domain.Manager{}

	_ = manager.AddEmployee(domain.Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	_ = manager.AddEmployee(domain.Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})

	employee := manager.FindEmployeeByID(2)
	if employee == nil || employee.Name != "Bob" {
		t.Errorf("Expected to find Bob, got %+v", employee)
	}

	manager.RemoveEmployee(2)
	employee = manager.FindEmployeeByID(2)
	if employee != nil {
		t.Errorf("Expected no employee, got %+v", employee)
	}
}
