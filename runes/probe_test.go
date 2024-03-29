package runes

import (
	"testing"
)

func TestProbe(t *testing.T) {
	type TestCase struct {
		probe Probe
		input string
	}

	tests := []TestCase{
		{Rune('a'), "a"},
		{Space(), " "},
		{Space(), "\t"},
		{Space(), "\n"},
		{And(Rune('1'), Rune('2'), Space()), "12 "},
		{Or(Rune('r'), Space(), Rune('x')), "r"},
		{Or(Rune('r'), Space(), Rune('x')), " "},
		{Or(Rune('r'), Space(), Rune('x')), "x"},
		{Any(Rune('w')), ""},
		{Any(Rune('w')), "w"},
		{Any(Rune('w')), "ww"},
		{Any(Rune('w')), "www"},
		{N(6, Rune('w')), "wwwwww"},
		{Maybe(Rune('w')), ""},
		{Maybe(Rune('w')), "w"},
	}

	for _, test := range tests {
		if read, ok := test.probe([]rune(test.input)); !ok {
			t.Errorf("Expected to read %s", string(test.input))
		} else if string(read) != test.input {
			t.Errorf("Mismatch of input %s and read %s", test.input, string(read))
		}
	}
}

func TestProbeFail(t *testing.T) {
	type TestCase struct {
		probe Probe
		input string
	}

	tests := []TestCase{
		{Rune('a'), "b"},
		{Space(), "a"},
		{And(Rune('1'), Rune('2'), Space()), "12"},
		{Or(Rune('r'), Space(), Rune('x')), "4"},
	}

	for _, test := range tests {
		if read, ok := test.probe([]rune(test.input)); ok {
			t.Errorf("Unexpectedly read %s with input %s", string(read), test.input)
		}
	}
}
