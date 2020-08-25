package main

import (
	"testing"
)

func Test_numAlphabet(t *testing.T) {
	type args struct {
		searchStr string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{"a"},
			1,
		}, {
			"",
			args{"b"},
			2,
		}, {
			"",
			args{"c"},
			3,
		}, {
			"",
			args{"d"},
			4,
		}, {
			"",
			args{"e"},
			5,
		}, {
			"",
			args{"f"},
			6,
		}, {
			"",
			args{"g"},
			7,
		}, {
			"",
			args{"h"},
			8,
		}, {
			"",
			args{"i"},
			9,
		}, {
			"",
			args{"j"},
			10,
		}, {
			"",
			args{"k"},
			11,
		}, {
			"",
			args{"l"},
			12,
		}, {
			"",
			args{"m"},
			13,
		}, {
			"",
			args{"n"},
			14,
		}, {
			"",
			args{"o"},
			15,
		}, {
			"",
			args{"p"},
			16,
		}, {
			"",
			args{"q"},
			17,
		}, {
			"",
			args{"r"},
			18,
		}, {
			"",
			args{"s"},
			19,
		}, {
			"",
			args{"t"},
			20,
		}, {
			"",
			args{"u"},
			21,
		}, {
			"",
			args{"v"},
			22,
		}, {
			"",
			args{"w"},
			23,
		}, {
			"",
			args{"x"},
			24,
		}, {
			"",
			args{"y"},
			25,
		}, {
			"",
			args{"z"},
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numAlphabet(tt.args.searchStr); got != tt.want {
				t.Errorf("numAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
