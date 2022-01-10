package functions

type empDetails struct {
	id     int
	age    int
	name   string
	desig  string
	hasPan bool
}

func CreateEmployee(id, age int, name, desig string, hasPan bool) empDetails {
	em := empDetails{id: id, age: age, name: name, desig: desig, hasPan: hasPan}
	return em
}

func (e *empDetails) GetId() int {
	return e.id
}

func (e *empDetails) GetAge() int {
	return e.age
}

func (e *empDetails) GetName() string {
	return e.name
}

func (e *empDetails) GetDesignation() string {
	return e.desig
}

func (e *empDetails) CheckEmpAge() (bool, empDetails) {
	if e.age < 22 {
		return false, empDetails{}
	}
	return true, *e
}
