package main

import (
	"io"
	"log"
)
var messagesByIssue = make(map[string][]string)

func main() {
	r := readData()
	if r == nil {
		return
	}

	// Iterate through the records, filling up the above map
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		checked := check(record)

		//nxtTime := checked[0]
		nxtIssue := checked[1]
		nxtMessage := checked[2]
		if nxtIssue != "issue_key" {

			//fmt.Printf("Time: %s Issue: %s Message: %s\n", nxtTime, nxtIssue, nxtMessage)

			// use it if it's NOT the value for "wrong number of columns"
			if nxtIssue != "********" {
				match := false
				issueMessages := messagesByIssue[nxtIssue]
				for i := 0; i < len(issueMessages); i++ {
					if issueMessages[i] == nxtMessage {
						match = true
						break
					}
				}
				if !match {
					messagesByIssue[nxtIssue] = append(messagesByIssue[nxtIssue], nxtMessage)
				}
			}
		}
	}
	setupUi()
}

func check(csvLine []string) ([]string) {
	result := make([]string, 3)
	if len(csvLine) != 3 {
		if len(csvLine) > 0 {
			result[0] = csvLine[0]
		} else {
			result[0] = "************************"
		}
		result[1] = "********"
		result[2] = "***** WRONG NUMBER OF COLUMNS *****"
		return result
	}
	return csvLine
}

