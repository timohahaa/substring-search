package algos

import "strings"

type BuiltinSearcher struct{}

func (s *BuiltinSearcher) Equals(pattern, text string, from, to int) bool {
	return true
}

func (s *BuiltinSearcher) Find(pattern, text string) int {
	return strings.Index(text, pattern)
}

func NewBuiltintSearcher() *BuiltinSearcher {
	return &BuiltinSearcher{}
}
