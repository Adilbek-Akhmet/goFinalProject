package main
type Employee struct {
	Name, Position string
}

func NewEmployeeFactory(position string) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position}
	}
}


