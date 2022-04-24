package hw9

import (
	"testing"
)

func TestOldestPerson(t *testing.T) {
	var got interface{}
	var want interface{}

	// Test #1
	got = OldestPerson(
		Employee2{25},
		Employee2{44},
		Employee2{18},
		Employee2{45},
		Employee2{28},
		Employee2{37},
		Employee2{36},
	)
	want = Employee2{45}
	if got != want {
		t.Errorf("OldestPerson -> Test #1 = %v, want %v", got, want)
	}

	// Test #2
	got = OldestPerson(
		Customer2{25},
		Customer2{44},
		Customer2{18},
		Customer2{45},
		Customer2{28},
		Customer2{37},
		Customer2{36},
	)
	want = Customer2{45}
	if got != want {
		t.Errorf("OldestPerson -> Test #2 = %v, want %v", got, want)
	}

	// Test #3
	got = OldestPerson(Employee2{74})
	want = Employee2{74}
	if got != want {
		t.Errorf("OldestPerson -> Test #3 = %v, want %v", got, want)
	}

	// Test #4
	got = OldestPerson(Customer2{24})
	want = Customer2{24}
	if got != want {
		t.Errorf("OldestPerson -> Test #4 = %v, want %v", got, want)
	}
}
