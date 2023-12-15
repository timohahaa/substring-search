package tests

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/timohahaa/substring-search/internal/algos"
	"github.com/timohahaa/substring-search/internal/interfaces"
)

const (
	RUNS = 10_000
)

type Mean struct {
	count int
	mean  float64
}

func (m *Mean) Add(x int64) {
	m.count++
	var delta float64 = float64(x) - m.mean
	m.mean += delta / float64(m.count)
}

func (m *Mean) GetMean() float64 {
	return m.mean
}

type searcherEntry struct {
	searcher interfaces.StringSearcher
	mean     *Mean
}

type Result struct {
	means         []float64
	patternLength int
	textLength    int
}

func testSearchers(text string, patternLength, runs int, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()

	searchers := []searcherEntry{
		{algos.NewBuiltintSearcher(), &Mean{}},
		{algos.NewNaiveSearcher(), &Mean{}},
		{algos.NewRabinKarpSearcher(), &Mean{}},
		{algos.NewBoyerMooreSearcher(), &Mean{}},
		{algos.NewKnuthMorrisPrattSearcher(), &Mean{}},
	}

	for i := 0; i < runs; i++ {
		// find a random pattern with a given patternLength
		// from [0, len(text) - patternLength)
		start := rand.Intn(len(text) - patternLength)
		end := start + patternLength
		pattern := text[start:end]
		expected := strings.Index(text, pattern)

		for idx, se := range searchers {
			startTime := time.Now()
			got := se.searcher.Find(pattern, text)
			endTime := time.Now()
			if got != expected {
				panic(fmt.Sprintf("got wrong result! patternLength: %v, pattern: %v, got: %v, expected: %v, searcher: %v\n", patternLength, pattern, got, expected, idx))
			}
			elapsed := endTime.Sub(startTime)
			se.mean.Add(int64(elapsed))
		}
	}

	// now send results to the channel
	result := []float64{}
	for _, se := range searchers {
		result = append(result, se.mean.GetMean())
	}

	results <- Result{result, patternLength, len(text)}
}
