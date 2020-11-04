package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var eventLines []string

func main() {
	term := "PD-Slackbot Test"

	scheduleFile, err := openStdinOrFile()

	if err != nil {
		panic(err)
	}

	rd := bufio.NewReader(scheduleFile)

	for {
		line, err := rd.ReadString('\n')

		if err != nil {
			if err == io.EOF { // last line may not have a newline
				onLine(line, term)
				break
			}

			panic(err)
		}

		onLine(line, term) // any but last line
	}

	printLines() // remainder
}

func onLine(line, term string) {
	l := strings.TrimSpace(line)

	if l == "END:VEVENT" {
		if !isSkippedEvent(term) {
			printLines()
		}

		eventLines = nil
	}

	eventLines = append(eventLines, l)
}

func isSkippedEvent(term string) bool {
	for _, el := range eventLines {
		if strings.Contains(el, term) {
			return true
		}
	}

	return false
}

func openStdinOrFile() (io.Reader, error) {
	var err error
	reader := os.Stdin

	if len(os.Args) > 1 {
		reader, err = os.Open(os.Args[1])

		if err != nil {
			return nil, err
		}
	}

	return reader, nil
}

func printLines() {
	for _, el := range eventLines {
		fmt.Println(el)
	}
}
