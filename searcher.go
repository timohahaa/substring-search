package substringsearch

type StringSearcher interface {
	Find(pattern, text string) int
	Equals(pattern, text string, from, to int) bool
}
