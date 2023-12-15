package algos

type RabinKarpSearcher struct{}

func (s *RabinKarpSearcher) Equals(pattern, text string, from, to int) bool {
	for i := from; i < to; i++ {
		if pattern[i-from] != text[i] {
			return false
		}
	}
	return true
}

func (s *RabinKarpSearcher) stringHashCode(str string, start, end int) int {
	result := 0
	for i := start; i < end; i++ {
		result += int(str[i])
	}
	return result
}

func (s *RabinKarpSearcher) Find(pattern, text string) int {
	patternLength := len(pattern)
	hashedPattern := s.stringHashCode(pattern, 0, patternLength)

	textLength := len(text)
	hashedTextWindow := s.stringHashCode(text, 0, patternLength)

	for i := 0; i <= textLength-patternLength; i++ {
		if hashedPattern == hashedTextWindow && s.Equals(pattern, text, i, i+patternLength) {
			return i
		}
		hashedTextWindow = hashedTextWindow - int(text[i]) + int(text[i+patternLength])
	}
	return -1
}

func NewRabinKarpSearcher() *RabinKarpSearcher {
	return &RabinKarpSearcher{}
}
