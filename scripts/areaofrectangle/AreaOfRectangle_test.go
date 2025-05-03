package main

import "testing"

func TestAreaOfRectangle(t *testing.T) {
	tests := []struct {
		width, height float64
		want          float64
	}{
		{2.0, 3.0, 6.0},
		{5.0, 5.0, 25.0},
		{0.0, 4.0, 0.0},
	}

	for _, tt := range tests {
		got :=
			AreaOfRectangle(tt.width, tt.height)
		if got != tt.want {
			t.Errorf("AreaOfRectangle(%v, %v) = %v, want %v", tt.width, tt.height, got, tt.want)
		}
	}
}
