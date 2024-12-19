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

func TestVector_Simplify(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		want   Vector
	}{
		{
			name: "simplify to 1/2",
			fields: fields{
				x: 2,
				y: 4,
			},
			want: Vector{
				x: 1,
				y: 2,
			},
		},
		{
			name: "simplify to 2/1",
			fields: fields{
				x: 16,
				y: 8,
			},
			want: Vector{
				x: 2,
				y: 1,
			},
		},
		{
			name: "simplify to 16/1",
			fields: fields{
				x: 16,
				y: 1,
			},
			want: Vector{
				x: 16,
				y: 1,
			},
		},
		{
			name: "simplify to -4/1",
			fields: fields{
				x: -16,
				y: 4,
			},
			want: Vector{
				x: -4,
				y: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := v.Simplify(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simplify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factors(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "4",
			args: args{a: 4},
			want: []int{4, 2, 1},
		},
		{
			name: "15",
			args: args{a: 15},
			want: []int{15, 5, 3, 1},
		},
		{
			name: "31",
			args: args{a: 31},
			want: []int{31, 1},
		},
		{
			name: "44",
			args: args{a: 44},
			want: []int{44, 22, 11, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factors(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcf(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,11",
			args: args{a: 1, b: 11},
			want: 1,
		},
		{
			name: "5,15",
			args: args{a: 5, b: 15},
			want: 5,
		},
		{
			name: "200, 112",
			args: args{a: 200, b: 112},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcf(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcf() = %v, want %v", got, tt.want)
			}
		})
	}
}
