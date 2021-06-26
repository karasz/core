package runes

import (
	"unicode"
)

// Probe was borrowed from https://github.com/JamesOwenHall/json2.Scanner
//

// Probe is a func that returns a subset of the input and a success bool.
type Probe func([]rune) ([]rune, bool)

// If returns a probe that accepts the a rune if it satisfies the condition.
func If(condition func(rune) bool) Probe {
	return func(input []rune) ([]rune, bool) {
		if len(input) > 0 && condition(input[0]) {
			return input[0:1], true
		}

		return nil, false
	}
}

// Rune returns a probe that accepts r.
func Rune(r rune) Probe {
	return If(func(b rune) bool {
		return r == b
	})
}

// Space returns a probe that accepts whitespace as defined in the unicode
// package.
func Space() Probe {
	return func(input []rune) ([]rune, bool) {
		if len(input) > 0 && unicode.IsSpace(input[0]) {
			return input[0:1], true
		}

		return nil, false
	}
}

// And returns a probe that accepts all probes in sequence.
func And(probes ...Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		remaining := input
		accumulated := []rune{}

		for _, s := range probes {
			if read, ok := s(remaining); !ok {
				return nil, false
			} else {
				accumulated = append(accumulated, read...)
				remaining = remaining[len(read):]
			}
		}

		return accumulated, true
	}
}

// Or returns a probe that accepts the first successful probe in probes.
func Or(probes ...Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		for _, s := range probes {
			if read, ok := s(input); ok {
				return read, true
			}
		}

		return nil, false
	}
}

// Maybe runs probe and returns true regardless of the output.
func Maybe(probe Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		read, _ := probe(input)
		return read, true
	}
}

// Any returns a probe that accepts any number of occurrences of probe,
// including zero.
func Any(probe Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		remaining := input
		accumulated := []rune{}

		for {
			if read, ok := probe(remaining); !ok {
				return accumulated, true
			} else {
				accumulated = append(accumulated, read...)
				remaining = remaining[len(read):]
			}
		}
	}
}

// N returns a probe that accepts probe exactly n times.
func N(n int, probe Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		probes := make([]Probe, n)
		for i := 0; i < n; i++ {
			probes[i] = probe
		}

		return And(probes...)(input)
	}
}

// AtLeast returns a probe that accepts probe at least n times.
func AtLeast(n int, probe Probe) Probe {
	return func(input []rune) ([]rune, bool) {
		probes := make([]Probe, n, n+1)
		for i := range probes {
			probes[i] = probe
		}

		probes = append(probes, Any(probe))
		return And(probes...)(input)
	}
}
