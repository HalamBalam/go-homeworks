package hw9

type Person interface {
	Age() int
}

type Employee struct {
	age int
}

type Customer struct {
	age int
}

func (e *Employee) Age() int {
	return e.age
}

func (c *Customer) Age() int {
	return c.age
}

func MaximumAge(persons ...Person) int {
	var res int
	for _, p := range persons {
		if p.Age() > res {
			res = p.Age()
		}
	}
	return res
}
