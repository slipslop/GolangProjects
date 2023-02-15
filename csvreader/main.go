package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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

		var input string
		fmt.Scanf("%s\n", &input)

		if input == pair.answer {
			statsRecorder.answersCorrect++
		}

		statsRecorder.totalQuestions++
	}

	fmt.Println("game is now over... your stats:")
	fmt.Println("You got ", statsRecorder.answersCorrect, "/", statsRecorder.totalQuestions, " right!")
}
