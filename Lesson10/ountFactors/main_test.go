package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{24},
			8,
		}, {
			"",
			args{1},
			1,
		}, {
			"",
			args{2},
			2,
		}, {
			"",
			args{3},
			2,
		}, {
			"",
			args{4},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
