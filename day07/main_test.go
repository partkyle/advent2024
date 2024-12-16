package main

import (
	"reflect"
	"testing"
)

func Test_permutations(t *testing.T) {
	type args struct {
		a     []string
		count int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "size of 1 is easy",
			args: args{
				count: 1,
			},
			want: [][]string{{"+"}, {"*"}},
		},
		{
			name: "size of 2",
			args: args{
				count: 2,
			},
			want: [][]string{{"+", "+"}, {"*", "+"}, {"+", "*"}, {"*", "*"}},
		},
		{
			name: "size of 3",
			args: args{
				count: 3,
			},
			want: [][]string{
				{"+", "+", "+"}, {"*", "+", "+"}, {"+", "*", "+"}, {"*", "*", "+"},
				{"+", "+", "*"}, {"*", "+", "*"}, {"+", "*", "*"}, {"*", "*", "*"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutations(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_record_isPossible(t *testing.T) {
	tests := []struct {
		name string
		line string
		want bool
	}{
		{
			name: "190: 10 19",
			line: "190: 10 19",
			want: true,
		},
		{
			name: "191: 10 19",
			line: "191: 10 19",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := parseLine(tt.line)
			if got := r.isPossible(); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
