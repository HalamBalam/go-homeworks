package hw9

type Employee2 struct {
	Age int
}

type Customer2 struct {
	Age int
}

type Person2 interface {
	Employee2 | Customer2
}

func OldestPerson[T Person2](persons ...T) T {
	var res T
	if len(persons) == 0 {
		return res
	}

	maxAge := 0
	for _, p := range persons {
		age := 0
		switch p1 := any(p).(type) {
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
