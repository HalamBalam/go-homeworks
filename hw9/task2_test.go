package hw9

import (
	"reflect"
	"testing"
)

func TestOldestPerson(t *testing.T) {
	type args struct {
		persons []any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test #1",
			args: args{persons: []any{
				Employee2{23},
				Employee2{40},
				Employee2{40},
				Employee2{66},
				Employee2{18},
				Employee2{65},
			}},
			want: Employee2{66},
		},
		{
			name: "Test #2",
			args: args{persons: []any{
				Customer2{20},
				Customer2{32},
				Customer2{18},
				Customer2{55},
				Customer2{88},
				Customer2{18},
			}},
			want: Customer2{88},
		},
		{
			name: "Test #3",
			args: args{persons: []any{
				Employee2{23},
			}},
			want: Employee2{23},
		},
		{
			name: "Test #4",
			args: args{persons: []any{
				Customer2{46},
			}},
			want: Customer2{46},
		},
		{
			name: "Test #5",
			args: args{persons: []any{}},
			want: nil,
		},
		{
			name: "Test #6",
			args: args{persons: []any{
				Customer2{43},
				Customer2{22},
				Employee2{18},
				Customer2{66},
				Employee2{56},
				Employee2{18},
			}},
			want: Customer2{66},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OldestPerson(tt.args.persons...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OldestPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
