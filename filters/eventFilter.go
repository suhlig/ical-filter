package filters

import (
	"fmt"
	"strings"
)

// EventFilter filters a line-oriented iCal stream
type EventFilter struct {
	lines          []string
	SkipIfContains []string
}

// OnLine handles a single line in a line-oriented iCal stream
func (f *EventFilter) OnLine(line string) {
	l := strings.TrimSpace(line)

	// BUG If the first event is skipped, we'll loose the lines preceding this event
	if l == "END:VEVENT" {
		if !f.isSkippedEvent() {
			f.Dump()
		}

		f.lines = nil
	}

	f.lines = append(f.lines, l)
}

// Dump prints all lines and clears the internal buffer
func (f *EventFilter) Dump() {
	for _, el := range f.lines {
		fmt.Println(el)
	}
}

func (f *EventFilter) isSkippedEvent() bool {
	for _, el := range f.lines {
		for _, skip := range f.SkipIfContains {
			if strings.Contains(el, skip) {
				return true
			}
		}
	}

	return false
}
