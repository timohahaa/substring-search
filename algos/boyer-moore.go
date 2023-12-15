package algos

type BoyerMooreSearcher struct{}

func (s *BoyerMooreSearcher) Equals(pattern, text string, from, to int) bool {
	for i := from; i < to; i++ {
		if pattern[i-from] != text[i] {
			return false
		}
	}
	return true
}

func (s *BoyerMooreSearcher) Find(pattern, text string) int {
	table := make(map[byte]int)
	for i := 0; i < len(pattern)-1; i++ {
		table[pattern[i]] = i
	}

OUTTER:
	for i := 0; i < len(text)-len(pattern)+1; {
		for j := len(pattern) - 1; j > -1; j-- {
			if pattern[i] != text[i+j] {
				move, ok := table[text[i+j]]
				if !ok {
					i += max(1, j-len(pattern))
				} else {
					i += max(1, move)
				}
				continue OUTTER
			}
		}
		return i
	}
	return -1
}

func NewBoyerMooreSearcher() *BoyerMooreSearcher {
	return &BoyerMooreSearcher{}
}
