package hw9

import "testing"

func TestMaximumAge(t *testing.T) {
	type args struct {
		persons []Person
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{persons: []Person{
				&Employee{23},
				&Employee{40},
				&Employee{40},
				&Employee{66},
				&Employee{18},
				&Employee{65},
			}},
			want: 66,
		},
		{
			name: "Test #2",
			args: args{persons: []Person{
				&Customer{20},
				&Customer{32},
				&Customer{18},
				&Customer{55},
				&Customer{88},
				&Customer{18},
			}},
			want: 88,
		},
		{
			name: "Test #3",
			args: args{persons: []Person{
				&Employee{23},
			}},
			want: 23,
		},
		{
			name: "Test #4",
			args: args{persons: []Person{
				&Customer{46},
			}},
			want: 46,
		},
		{
			name: "Test #5",
			args: args{persons: []Person{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumAge(tt.args.persons...); got != tt.want {
				t.Errorf("MaximumAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
