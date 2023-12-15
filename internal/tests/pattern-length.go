package tests

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strconv"
	"sync"
)

// vary the pattern length from 0 to 1000
// (or from 0 to 445 with Lorem)

func testPatternLength(text string, runs, maxPatternLength int, writer *csv.Writer) {
	writer.Write([]string{"No", "length", "Go", "Naive", "RK", "BM", "KMP"})

	results := make(chan Result)
	wg := &sync.WaitGroup{}

	for i := 1; i < maxPatternLength; i++ {
		wg.Add(1)
		go testSearchers(text, i, runs, results, wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	No := 1
	for result := range results {
		record := make([]string, 7)
		record[0] = strconv.Itoa(No)
		record[1] = strconv.Itoa(result.patternLength)
		record[2] = strconv.FormatFloat(result.means[0], 'f', -1, 64)
		record[3] = strconv.FormatFloat(result.means[1], 'f', -1, 64)
		record[4] = strconv.FormatFloat(result.means[2], 'f', -1, 64)
		record[5] = strconv.FormatFloat(result.means[3], 'f', -1, 64)
		record[6] = strconv.FormatFloat(result.means[4], 'f', -1, 64)
		err := writer.Write(record)
		if err != nil {
			fmt.Println(err)
		}
		No += 1
	}
}

func RunTestForVaryintPatternLength(pathToFile, outputName string, maxPatternLength int) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	// remove \n char
	fileBytes = fileBytes[:len(fileBytes)-1]
	fileText := string(fileBytes)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.OpenFile(outputName, os.O_RDWR|os.O_CREATE, 0666)
	defer outFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	outCSV := csv.NewWriter(outFile)

	slog.Info("Starting pattern length test", "file", pathToFile)
	testPatternLength(fileText, RUNS, maxPatternLength, outCSV)
	slog.Info("Finished pattern length test", "file", pathToFile)

	outCSV.Flush()
	err = outCSV.Error()
	if err != nil {
		log.Fatal(err)
	}

	err = outFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
