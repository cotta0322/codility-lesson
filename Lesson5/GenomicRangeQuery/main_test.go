package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		S string
		P []int
		Q []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{
				"CAGCCTA",
				[]int{2, 5, 0},
				[]int{4, 5, 6},
			},
			[]int{2, 4, 1},
		}, {
			"",
			args{
				"ACGTACGT",
				[]int{0},
				[]int{1},
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.S, tt.args.P, tt.args.Q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
