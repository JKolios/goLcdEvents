package lcd

import (
	"time"
)

const (
	NO_FLASH         = -1
	BEFORE           = 0
	AFTER            = 1
	BEFORE_AND_AFTER = 2

	EVENT_DISPLAY  = 0
	EVENT_SHUTDOWN = 1
)

type LcdEvent struct {
	eventType               int
	message                 string
	duration                time.Duration
	flash, flashRepetitions int
	clearAfter              bool
}

func newLcdEvent(eventType int, message string, duration time.Duration, flash int, flashRepetitions int, clearAfter bool) *LcdEvent {
	return &LcdEvent{eventType, message, duration, flash, flashRepetitions, clearAfter}
}
