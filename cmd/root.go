package cmd

import (
	"bufio"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/suhlig/ical-filter/filters"
)

var shortDescription = "Reads iCal events, filters them and prints the filtered events. "

var rootCmd = &cobra.Command{
	Use:   "ical-filter",
	Short: shortDescription,
	Long:  shortDescription + "\n\nA calendar (stream of iCal VEVENTs) is expected on STDIN or as a file name parameter.",
	RunE:  run,
	Example: `
  # Fetch the NASA calendar and remove all SpaceX events
  $ curl https://www.nasa.gov/templateimages/redesign/calendar/iCal/nasa_calendar.ics | ical-filter --skip SpaceX`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringSliceP("skip", "s", nil, "skip all events containing this string")
	rootCmd.Flags().BoolP("verbose", "V", false, "verbose output")
}

func run(cmd *cobra.Command, args []string) error {
	calendar, err := stdinOrFileArg(args)

	if err != nil {
		return err
	}

	reader := bufio.NewReader(calendar)

	verbose, err := cmd.Flags().GetBool("verbose")

	if err != nil {
		return err
	}

	skips, err := cmd.Flags().GetStringSlice("skip")

	if err != nil {
		return err
	}

	filter := filters.EventFilter{
		SkipIfContains: skips,
		Verbose:        verbose,
	}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF { // last line may not have a newline
				filter.OnLine(line)
				break
			}

			return err
		}

		filter.OnLine(line) // any but last line
	}

	filter.Dump() // print the remainder

	return nil
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
