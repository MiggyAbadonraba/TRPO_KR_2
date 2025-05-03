package main

import (
	"math"
	"testing"
)

func TestVolumeOfSphere(t *testing.T) {
	tests := []struct {
		radius float64
		want   float64
	}{
		{1.0, (4.0 / 3.0) * math.Pi},
		{2.0, (32.0 / 3.0) * math.Pi},
		{0.0, 0.0},
	}

	for _, tt := range tests {
		got := volumeOfSphere(tt.radius)
		if got != tt.want {
			t.Errorf("VolumeOfSphere(%v) = %v, want %v", tt.radius, got, tt.want)
		}
	}
}

func TestAreaOfCircle(t*testing.T){
	test := []struct{
		radius float64
		want float64
	}{
	{1.0, math.Pi},
	{2.0, 4 * math.Pi},
	{0.0, 0.0},
}
for_, tt := range test {
	got := AreaOfCircle(tt.radius)
	if got != tt.want {
		t.Errorf("AreaOfCircle(%v) = %v, Want %v", tt.radius, got, tt.want)
	}
}
}
