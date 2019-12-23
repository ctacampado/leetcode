package main

import "testing"

type TestCase struct {
	input []string
	exp   string
}

func TestAddBinary_AddZero(t *testing.T) {
	tc := TestCase{
		input: []string{"0000", "0001"},
		exp:   "0001",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_AddOneOne(t *testing.T) {
	tc := TestCase{
		input: []string{"0001", "0001"},
		exp:   "0010",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_SameLength(t *testing.T) {
	tc := TestCase{
		input: []string{"1010", "1011"},
		exp:   "10101",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_AddDiffLength(t *testing.T) {
	tc := TestCase{
		input: []string{"101", "1"},
		exp:   "110",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_AddDiffTwoLengthCarry(t *testing.T) {
	tc := TestCase{
		input: []string{"11", "1"},
		exp:   "100",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_AddDiffThreeLengthCarry(t *testing.T) {
	tc := TestCase{
		input: []string{"111", "1"},
		exp:   "1000",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func TestAddBinary_Combination(t *testing.T) {
	tc := TestCase{
		input: []string{"10101", "1011"},
		exp:   "100000",
	}

	got := AddBinary(tc.input[0], tc.input[1])
	if !equal(tc.exp, got) {
		t.Fatalf("output mismatch!\nexp: %s\ngot: %s", tc.exp, got)
	}
}

func equal(exp, got string) bool {
	if exp == got {
		return true
	}
	return false
}
