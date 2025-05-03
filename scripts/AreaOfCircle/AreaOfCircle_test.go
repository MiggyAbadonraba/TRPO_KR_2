package main

import (
	"math"
	"testing"
)

func TestAreaOfCircle(t *testing.T) {
	tests := []struct {
		radius float64
		want   float64
	}{
		{1.0, math.Pi},
		{2.0, 4 * math.Pi},
		{0.0, 0.0},
	}
	for _, tt := range tests {
		got := AreaOfCircle(tt.radius)
		if got != tt.want {
			t.Errorf("AreaOfCircle(%v) = %v, Want %v", tt.radius, got, tt.want)
		}
	}
}
