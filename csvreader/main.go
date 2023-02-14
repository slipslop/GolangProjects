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

func main() {
	fileName := flag.String("file", "source.csv", "a string")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	stdinReader := bufio.NewReader(os.Stdin)

	var answersCorrect int = 0
	var answersWrong int = 0

	for {
		record, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var question string = record[0]
		var answer string = record[1]

		fmt.Print(question, " ?")
		input, err := stdinReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.ReplaceAll(input, "\n", "")

		fmt.Println((input))
		fmt.Println((answer))
		if input == answer {
			answersCorrect++
		} else {
			answersWrong++
		}

	}

	fmt.Println("game is now over... your stats:")
	fmt.Println(answersCorrect, answersWrong)
}
