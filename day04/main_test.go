package main

import (
	"reflect"
	"testing"
)

func Test_rot(t *testing.T) {
	data := []string{
		"123",
		"456",
		"789",
	}

	v := rot(data)

	expected := []string{
		"741",
		"852",
		"963",
	}
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("rot() failed: expected %v, got %v", expected, v)
	}
}

func Test_rot4(t *testing.T) {
	data := []string{
		"123a",
		"456b",
		"789c",
	}

	v := rot(rot(rot(rot(data))))
	if !reflect.DeepEqual(v, data) {
		t.Errorf("rot() failed: expected %v, got %v", data, v)
	}
}
