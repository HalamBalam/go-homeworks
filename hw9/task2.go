package hw9

type Employee2 struct {
	Age int
}

type Customer2 struct {
	Age int
}

func OldestPerson(persons ...any) any {
	var res any
	if len(persons) == 0 {
		return res
	}

	maxAge := 0
	for _, p := range persons {
		age := 0
		switch p1 := p.(type) {
		case Employee2:
			age = p1.Age
		case Customer2:
			age = p1.Age
		}
		if age > maxAge {
			res = p
			maxAge = age
		}
	}
	return res
}
