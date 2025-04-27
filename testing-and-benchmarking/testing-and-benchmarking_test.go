package main

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	got := IntMin(2, -2)
	want := -2
	if got != want {
		t.Errorf("IntMin(2, -2) = %d; want %d", got, want)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{

		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		description := fmt.Sprintf("IntMin(%d,%d)=%d", tt.a, tt.b, tt.want)
		t.Run(description, func(t *testing.T) {
			got := IntMin(tt.a, tt.b)
			if got != tt.want {
				// t.Errorf reports the error but continue executing the test
				// t.Fatal* will report the test failures and stop the test immediately
				t.Errorf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for b.Loop() {
		IntMin(1, 2)
	}
}
