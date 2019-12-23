package main

import "testing"

import "reflect"

type TestCase struct {
	input []int
	exp   []int
}

func TestPlusOne_ZeroArray(t *testing.T) {
	tc := TestCase{
		input: []int{0, 0, 0, 0},
		exp:   []int{0, 0, 0, 1},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func TestPlusOne_ExampleOne(t *testing.T) {
	tc := TestCase{
		input: []int{1, 2, 3},
		exp:   []int{1, 2, 4},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func TestPlusTwo_ExampleOne(t *testing.T) {
	tc := TestCase{
		input: []int{4, 3, 2, 1},
		exp:   []int{4, 3, 2, 2},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func TestPlusTwo_CarryEnd(t *testing.T) {
	tc := TestCase{
		input: []int{4, 3, 2, 9},
		exp:   []int{4, 3, 3, 0},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func TestPlusTwo_CarryStart(t *testing.T) {
	tc := TestCase{
		input: []int{4, 9, 9, 9},
		exp:   []int{5, 0, 0, 0},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func TestPlusTwo_CarryStartMax(t *testing.T) {
	tc := TestCase{
		input: []int{9, 9, 9, 9},
		exp:   []int{1, 0, 0, 0, 0},
	}

	got := PlusOne(tc.input)

	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp:%+v\ngot:%+v", tc.exp, got)
	}
}

func equal(exp, got []int) bool {
	return reflect.DeepEqual(exp, got)
}
