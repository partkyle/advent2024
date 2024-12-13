package main

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph[int]()

	graph.addEdge(47, 53)
	graph.addEdge(97, 13)

	if graph.hasEdge(47, 53) != true {
		t.Error("graph.hasEdge(47, 53) != true")
	}

	if graph.hasEdge(97, 11) != false {
		t.Error("graph.hasEdge(97, 11) != false")
	}
}

func Test_containsAll(t *testing.T) {
	type args[E comparable] struct {
		haystack []E
		needles  []E
	}
	type testCase[E comparable] struct {
		name string
		args args[E]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "complete subset",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: " subset",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 3},
			},
			want: true,
		},
		{
			name: " single",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1},
			},
			want: true,
		},

		{
			name: " single not int",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{11},
			},
			want: false,
		},
		{
			name: "1 missing",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 12},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAll(tt.args.haystack, tt.args.needles); got != tt.want {
				t.Errorf("containsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsAny(t *testing.T) {
	type args[E comparable] struct {
		haystack []E
		needles  []E
	}
	type testCase[E comparable] struct {
		name string
		args args[E]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "complete subset",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: " subset",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 3},
			},
			want: true,
		},
		{
			name: " single",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1},
			},
			want: true,
		},

		{
			name: " single not int",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{11},
			},
			want: false,
		},
		{
			name: "1 missing",
			args: args[int]{
				haystack: []int{1, 2, 3},
				needles:  []int{1, 12},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsAny(tt.args.haystack, tt.args.needles); got != tt.want {
				t.Errorf("containsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args[E comparable] struct {
		a []E
		b []E
	}
	type testCase[E comparable] struct {
		name string
		args args[E]
		want []E
	}
	tests := []testCase[int]{
		{
			name: "complete subset",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			want: []int{1, 2, 3},
		},
		{
			name: " subset",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{1, 3},
			},
			want: []int{1, 3},
		},
		{
			name: " single",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{1},
			},
			want: []int{1},
		},

		{
			name: " single not int",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{11},
			},
			want: nil,
		},
		{
			name: "1 missing",
			args: args[int]{
				a: []int{1, 2, 3},
				b: []int{1, 12},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
