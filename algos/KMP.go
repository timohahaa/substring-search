package algos

type KnuthMorrisPrattSearcher struct{}

func (s *KnuthMorrisPrattSearcher) Equals(pattern, text string, from, to int) bool {
	for i := from; i < to; i++ {
		if pattern[i-from] != text[i] {
			return false
		}
	}
	return true
}

func (s *KnuthMorrisPrattSearcher) computeLPSTable(pattern string) []int {
	table := make([]int, len(pattern)+1)
	k := 0
	for q := 1; q < len(pattern); q++ {
		for k > 0 && pattern[k] != pattern[q] {
			k = table[k]
		}
		if pattern[k] == pattern[q] {
			k += 1
		}
		table[q+1] = k
	}
	return table
}

func (s *KnuthMorrisPrattSearcher) Find(pattern, text string) int {
	LPSTable := s.computeLPSTable(pattern)
	q := 0
	for i := 0; i < len(text); i++ {
		for q > 0 && pattern[q] != text[i] {
			q = LPSTable[q]
		}
		if pattern[q] == text[i] {
			q += 1
		}
		if q == len(pattern) {
			return i - len(pattern) + 1
		}
	}
	return -1
}

func NewKnuthMorrisPrattSearcher() *KnuthMorrisPrattSearcher {
	return &KnuthMorrisPrattSearcher{}
}
