package utils

import (
	"strings"
)

type Str string

func (s *Str) Len() int {
	return len(*s)
}
func (s *Str) Upper() string {
	return strings.ToUpper(string(*s))
}
func (s *Str) Lower() string {
	return strings.ToLower(string(*s))
}
func (s *Str) Reverse() (result string) {
	for _, v := range *s {
		result = string(v) + result
	}
	return
}

func (s *Str) Split(split string) []string {

	return strings.Split(string(*s), split)
}
// testing fix