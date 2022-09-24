package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsv(filePath string) [][]string {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, ".\n", err)
	}

	r := csv.NewReader(file)

	content, err := r.ReadAll()
	if err != nil {
		log.Fatal("Unable to read supplied .csv file at "+filePath, ".\n", err)
	}

	return content
}

func readAnswer(timeLimit int) string {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text

}

func runQuiz(qs [][]string, timeLimit int) (int, int) {

	score := 0
	numQs := len(qs)

	for q := 0; q < numQs; q++ {

		question := qs[q][0]
		correctAnswer := qs[q][1]

		fmt.Printf("Problem #%s: %s = ", strconv.Itoa(q+1), question)
		answer := readAnswer(timeLimit)

		if answer == correctAnswer {
			score += 1
		}
	}

	return score, numQs
}

func main() {

	// get command line args
	csvPtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limitPtr := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// read CSV
	csv := readCsv(*csvPtr)

	// run Quiz
	score, numQs := runQuiz(csv, *limitPtr)
	fmt.Printf("\nYou scored %d out of %d.", score, numQs)
}
