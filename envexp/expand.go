// Package envexp provides environment variable expansion using sh syntax
package envexp

import (
	"os"
)

// A VarGetter is a function that for a given paramenter
// returns a string representing its value
type VarGetter func(key string) string

// A Expander is an object that you can use to Expand strings
type Expander struct {
	get VarGetter
}

// NewExpander allocates a new Expander using a
// given VarGetter, or os.Getenv if none is given
func NewExpander(f VarGetter) *Expander {
	if f == nil {
		f = os.Getenv
	}

	return &Expander{
		get: f,
	}
}

// Get resolves a variable name as the given Expander
// would do
func (exp *Expander) Get(key string) string {
	return exp.get(key)
}

// Expand expands a string using the VarGetter
// defined for the given Expander
func (exp *Expander) Expand(s string) string {
	return s
}
