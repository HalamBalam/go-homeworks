package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestInts(t *testing.T) {
	got := []int{4, 3, 9, 11, 0, 4, -2}
	sort.Ints(got)
	want := []int{-2, 0, 3, 4, 4, 9, 11}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("sort.Ints() = %v, want %v", got, want)
	}
}

func TestStrings(t *testing.T) {
	type args struct {
		x []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test #1",
			args: args{x: []string{"Go", "1C", "C#", "Ruby", ""}},
			want: []string{"", "1C", "C#", "Go", "Ruby"},
		},
		{
			name: "Test #2",
			args: args{x: []string{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.args.x)
			if !reflect.DeepEqual(tt.want, tt.args.x) {
				t.Errorf("sort.Strings() = %v, want %v", tt.args.x, tt.want)
			}
		})
	}
}

func sampleDataInt() []int {
	rand.Seed(time.Now().UnixNano())
	var data []int
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Intn(1000))
	}

	return data
}

func sampleDataFloat64() []float64 {
	rand.Seed(time.Now().UnixNano())
	var data []float64
	for i := 0; i < 1_000_000; i++ {
		data = append(data, rand.Float64())
	}

	return data
}

func BenchmarkInts(b *testing.B) {
	data := sampleDataInt()
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkFloat64s(b *testing.B) {
	data := sampleDataFloat64()
	for i := 0; i < b.N; i++ {
		sort.Float64s(data)
	}
}
