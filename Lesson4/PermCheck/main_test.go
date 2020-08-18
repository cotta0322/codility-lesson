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
		{"", args{[]int{1}}, 1},
		{"", args{[]int{2}}, 0},
		{"", args{[]int{1, 2}}, 1},
		{"", args{[]int{2, 2}}, 0},
		{"", args{[]int{1, 3, 2}}, 1},
		{"", args{[]int{1, 3, 3}}, 0},
		{"", args{[]int{4, 1, 3, 2}}, 1},
		{"", args{[]int{4, 1, 3}}, 0},
		{"", args{[]int{1, 3, 5, 4, 2}}, 1},
		{"", args{[]int{1, 3, 6, 4, 1, 2}}, 0},
		{"", args{[]int{1, 3, 5, 4, 6, 2}}, 1},
		{"", args{[]int{1000000}}, 0},
		{"", args{[]int{1, 1000000}}, 0},
		{"", args{[]int{2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
