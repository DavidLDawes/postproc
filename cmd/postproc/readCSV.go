package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readData() *csv.Reader {
	csvfile, err := os.Open(fn)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Read the first record from csv
	header, err := r.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if len(header) != 3 || header[1] != issue || header[2] != message {
		log.Fatal("Failed to recognize column names from the csv file first line")
		return nil
	}
	return r
}
