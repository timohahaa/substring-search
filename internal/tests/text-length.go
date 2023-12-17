package tests

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"strconv"
)

// vary the text length from 0 + patternLength to 5000 + patternLength
// (or from 0 + patternLength to 445 for Lorem)
// try patternLength as 10, then 256, then 1000

func testTextLength(text string, runs, patternLength, maxTextLength int, writer *csv.Writer) {
	writer.Write([]string{"No", "length", "Go", "Naive", "RK", "BM", "KMP"})

	results := &[]Result{}

	for i := patternLength + 1; i < maxTextLength; i++ {
		// creating less goroutines makes test data way more clear!
		/*go*/
		testSearchers(text[0:i], patternLength, runs, results)
	}

	No := 1
	for _, result := range *results {
		record := make([]string, 7)
		record[0] = strconv.Itoa(No)
		record[1] = strconv.Itoa(result.textLength)
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

func RunTestForVaryintTextLength(pathToFile, outputName string, patternLength, maxTextLength int) {
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

	slog.Info(fmt.Sprintf("Starting text length test with pattern length %v", patternLength), "file", pathToFile)
	testTextLength(fileText, RUNS, patternLength, maxTextLength, outCSV)
	slog.Info(fmt.Sprintf("Finished text length test with pattern length %v", patternLength), "file", pathToFile)

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
