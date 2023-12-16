package main

import (
	//	"log/slog"
	//	"sync"

	"fmt"
	"log/slog"
	"math"
	"runtime/debug"
	"time"

	"github.com/timohahaa/substring-search/internal/tests"
)

//func TestWrapper(filePath, outputName string, maxPatternLength int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	tests.RunTestForVaryintPatternLength(filePath, outputName, maxPatternLength)
//}

func main() {
	// wg := &sync.WaitGroup{}
	//
	// wg.Add(1)
	// go TestWrapper("resources/lorem.txt", "resources/lorem_pattern_length.csv", 445, wg)
	//
	// wg.Add(1)
	// go TestWrapper("resources/dna.txt", "resources/dna_pattern_length.csv", 1000, wg)
	//
	// wg.Add(1)
	// go TestWrapper("resources/alice.txt", "resources/alice_pattern_length.csv", 1000, wg)
	//
	// wg.Wait()
	// slog.Info("All tests are done!")
	memlimit := debug.SetMemoryLimit(math.MaxInt64)

	start := time.Now()

	// lorem
	//tests.RunTestForVaryintPatternLength("resources/lorem.txt", "resources/lorem_pattern_length.csv", 445)
	//tests.RunTestForVaryintTextLength("resources/lorem.txt", "resources/lorem_text_length_10.csv", 10, 445)
	//tests.RunTestForVaryintTextLength("resources/lorem.txt", "resources/lorem_text_length_256.csv", 256, 445)
	//
	//	// dna
	//tests.RunTestForVaryintPatternLength("resources/dna.txt", "resources/dna_pattern_length.csv", 1000)
	//tests.RunTestForVaryintTextLength("resources/dna.txt", "resources/dna_text_length_10.csv", 10, 5010)
	//tests.RunTestForVaryintTextLength("resources/dna.txt", "resources/dna_text_length_256.csv", 256, 5256)
	//tests.RunTestForVaryintTextLength("resources/dna.txt", "resources/dna_text_length_1000.csv", 1000, 5871)
	//
	//	// alice
	//tests.RunTestForVaryintPatternLength("resources/alice.txt", "resources/alice_pattern_length.csv", 1000)
	//tests.RunTestForVaryintTextLength("resources/alice.txt", "resources/alice_text_length_10.csv", 10, 5010)
	//tests.RunTestForVaryintTextLength("resources/alice.txt", "resources/alice_text_length_256.csv", 256, 5256)
	tests.RunTestForVaryintTextLength("resources/alice.txt", "resources/alice_text_length_1000.csv", 1000, 6000)
	//
	//	// Marcus Aurelius (244067 symbols)
	//	tests.RunTestForVaryintPatternLength("resources/aurelius.txt", "resources/aurelius_pattern_length.csv", 1000)
	//	tests.RunTestForVaryintTextLength("resources/aurelius.txt", "resources/aurelius_text_length_10.csv", 10, 5010)
	//	tests.RunTestForVaryintTextLength("resources/aurelius.txt", "resources/aurelius_text_length_256.csv", 256, 5256)
	//	tests.RunTestForVaryintTextLength("resources/aurelius.txt", "resources/aurelius_text_length_1000.csv", 1000, 6000)

	end := time.Now()

	elapsed := end.Sub(start)
	slog.Info(fmt.Sprintf("All tests finished! Elapsed time: %v", elapsed.String()))

	debug.SetMemoryLimit(memlimit)
}
