package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhlig/ical-filter/filters"
)

var rootCmd = &cobra.Command{
	Use:   "ical-filter",
	Short: "Reads iCal events, filters them and prints them out again",
	Run:   run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceP("skip", "s", nil, "Help message for toggle")
}

func run(cmd *cobra.Command, args []string) {
	scheduleFile, err := stdinOrFileArg(args)

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(scheduleFile)

	skips, err := cmd.Flags().GetStringSlice("skip")

	if err != nil {
		panic(err)
	}

	filter := filters.EventFilter{
		SkipIfContains: skips,
	}

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

func stdinOrFileArg(args []string) (io.Reader, error) {
	var err error
	reader := os.Stdin

	if len(args) >= 1 {
		reader, err = os.Open(args[0])

		if err != nil {
			return nil, err
		}
	}

	return reader, nil
}
