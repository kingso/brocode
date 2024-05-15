package pkg

import (
	"strings"
)

type Str string
type Strs []string

// Len returns the length of the Str string as an int.
func (s *Str) Len() int {
	return len(*s)
}

// Upper returns the Str string as an uppercase string.
func (s *Str) Upper() Str {
	return Str(strings.ToUpper(string(*s)))
}

// Lower returns the Str string as an lowercase string.
func (s *Str) Lower() Str {
	return Str(strings.ToLower(string(*s)))
}

// Reverse returns the Str string as an uppercase string.
func (s *Str) Reverse() (result Str) {
	for _, v := range *s {
		result = Str(v) + result
	}
	return
}

// Split returns the Str string split by delim.
func (s *Str) Split(delim string) Strs {
	return Strs(strings.Split(string(*s), delim))
}
