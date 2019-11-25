package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file of problems 'question, answer' ")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		error(fmt.Sprintf("Failed to open csv file: %s \n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		error(fmt.Sprintf("Failed to parse csv file: %s \n", *csvFilename))
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You have %d answers correct out of %d.\n", correct, len(problems))
}

func error(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
