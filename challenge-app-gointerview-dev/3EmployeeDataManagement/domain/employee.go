package domain

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type ErrEmployeeAlreadyExist struct {
	EmployeeId int
}

func (e ErrEmployeeAlreadyExist) Error() string {
	return fmt.Sprintf("employee with id: %d already exist", e.EmployeeId)
}

type Manager struct {
	Employees []Employee
}

func (m *Manager) AddEmployee(employee Employee) error {
	for _, existEmployee := range m.Employees {
		if existEmployee.ID == employee.ID {
			return ErrEmployeeAlreadyExist{EmployeeId: employee.ID}
		}
	}

	m.Employees = append(m.Employees, employee)

	return nil
}
