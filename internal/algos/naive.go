package algos

type NaiveSearcher struct{}

func (s *NaiveSearcher) Equals(pattern, text string, from, to int) bool {
	for i := from; i < to; i++ {
		if pattern[i-from] != text[i] {
			return false
		}
	}
	return true
}

func (s *NaiveSearcher) Find(pattern, text string) int {
	for i := 0; i <= len(text)-len(pattern); i++ {
		if s.Equals(pattern, text, i, i+len(pattern)) {
			return i
		}
	}
	return -1
}

func NewNaiveSearcher() *NaiveSearcher {
	return &NaiveSearcher{}
}
