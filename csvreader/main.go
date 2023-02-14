package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	fileName := flag.String("file", "source.csv", "a string");
	flag.Parse();

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close();
}
