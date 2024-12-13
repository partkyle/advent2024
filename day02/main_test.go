package main

import (
	"testing"
)

func Test_safept1(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{
			line: "7 6 4 2 1",
			want: 1,
		},
		{
			line: "1 2 7 8 9",
			want: 0,
		},
		{
			line: "9 7 6 2 1",
			want: 0,
		},
		{
			line: "1 3 2 4 5",
			want: 0,
		},
		{
			line: "8 6 4 4 1",
			want: 0,
		},
		{
			line: "1 3 6 7 9",
			want: 1,
		},
	}
	for _, tt := range tests {
		name := tt.line
		t.Run(name, func(t *testing.T) {
			if got := pt1safe(transform(tt.line)); got != tt.want {
				t.Errorf("safe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_safept2(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{
			line: "7 6 4 2 1",
			want: 1,
		},
		{
			line: "1 2 7 8 9",
			want: 0,
		},
		{
			line: "9 7 6 2 1",
			want: 0,
		},
		{
			line: "1 3 2 4 5",
			want: 1,
		},
		{
			line: "8 6 4 4 1",
			want: 1,
		},
		{
			line: "1 3 6 7 9",
			want: 1,
		},
	}
	for _, tt := range tests {
		name := tt.line
		t.Run(name, func(t *testing.T) {
			if got := pt2Safe(transform(tt.line)); got != tt.want {
				t.Errorf("safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
