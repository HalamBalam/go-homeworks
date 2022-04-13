package hw

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		p1           Point
		p2           Point
		wantDistance float64
	}{
		{
			name:         "#1",
			p1:           Point{1, 1},
			p2:           Point{4, 5},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance, _ := CalculateDistance(tt.p1, tt.p2); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
