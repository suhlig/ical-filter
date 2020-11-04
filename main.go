package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	scheduleFile, err := openStdinOrFile()

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(scheduleFile)
	filter := EventFilter{SkipIfContains: "PD-Slackbot Test"}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF { // last line may not have a newline
				filter.OnLine(line)
				break
			}

			panic(err)
		}

		filter.OnLine(line) // any but last line
	}

	filter.Dump() // print the remainder
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
