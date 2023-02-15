package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type stats struct {
	totalQuestions int
	answersCorrect int
}

type questionAnswerPair struct {
	question string
	answer   string
}

func main() {
	fileName := flag.String("file", "source.csv", "a string")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	statsRecorder := stats{totalQuestions: 0, answersCorrect: 0}
	stdinReader := bufio.NewReader(os.Stdin)
	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		pair := questionAnswerPair{question: record[0], answer: record[1]}

		fmt.Print(pair.question, ": ")

		input, err := stdinReader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.ReplaceAll(input, "\n", "")

		if input == pair.answer {
			statsRecorder.answersCorrect++
		}

		statsRecorder.totalQuestions++
	}

	fmt.Println("game is now over... your stats:")
	fmt.Println("You got ", statsRecorder.answersCorrect, "/", statsRecorder.totalQuestions, " right!")
}
