package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestUniquePairs(t *testing.T) {
	type args[E any] struct {
		l []E
	}
	type testCase[E any] struct {
		name string
		args args[E]
		want [][2]E
	}
	tests := []testCase[int]{
		{
			name: "size 2",
			args: args[int]{
				l: []int{1, 2},
			},
			want: [][2]int{{1, 2}},
		},
		{
			name: "size 3",
			args: args[int]{
				l: []int{1, 2, 3},
			},
			want: [][2]int{{1, 2}, {1, 3}, {2, 3}},
		},
		{
			name: "size 5",
			args: args[int]{
				l: []int{1, 2, 3, 4, 5},
			},
			want: [][2]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := UniquePairs(tt.args.l)
			got := slices.Collect(seq)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniquePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Within(t *testing.T) {
	type fields struct {
		x int
		y int
	}

	lo := Vector{
		x: 0,
		y: 0,
	}
	hi := Vector{
		x: 10,
		y: 10,
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "within",
			fields: fields{
				x: 4,
				y: 4,
			},
			want: true,
		},
		{
			name: "lower edge",
			fields: fields{
				x: 0,
				y: 0,
			},
			want: true,
		},
		{
			name: "outer edge",
			fields: fields{
				x: 9,
				y: 9,
			},
			want: true,
		},
		{
			name: "oob",
			fields: fields{
				x: 1,
				y: 11,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := v.Within(lo, hi); got != tt.want {
				t.Errorf("Within() = %v, want %v", got, tt.want)
			}
		})
	}
}
