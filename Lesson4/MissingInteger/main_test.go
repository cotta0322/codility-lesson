package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{0, 1, 2, 100}},
			3,
		},
		{
			"",
			args{[]int{-1, -1, -1}},
			1,
		},
		{
			"",
			args{[]int{-1, -1, -1, 0}},
			1,
		},
		{
			"",
			args{[]int{-1, -1, -1, 0, 1}},
			2,
		},
		{
			"",
			args{[]int{0}},
			1,
		},
		{
			"",
			args{[]int{0, 1}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
