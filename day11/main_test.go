package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToDigits(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{i: 2024},
			want: []int{2, 0, 2, 4},
		},
		{
			args: args{i: 1000},
			want: []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.args.i), func(t *testing.T) {
			if got := ToDigits(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		d []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2024",
			args: args{d: []int{2, 0, 2, 4}},
			want: 2024,
		},
		{
			name: "024",
			args: args{d: []int{0, 2, 4}},
			want: 24,
		},
		{
			name: "118999",
			args: args{d: []int{1, 1, 8, 9, 9, 9}},
			want: 118999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.d); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
