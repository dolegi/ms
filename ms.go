// Converts time strings to and from number of milliseconds
package ms

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type time struct {
	short string
	long  string
}

const s = 1000
const m = s * 60
const h = m * 60
const d = h * 24
const w = d * 7
const y = d * 365.25
const match = `^(-?(?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|weeks?|w|years?|yrs?|y)?$`

var suffixes = map[string]time{
	"day":         time{"d", "day"},
	"hour":        time{"h", "hour"},
	"minute":      time{"m", "minute"},
	"second":      time{"s", "second"},
	"millisecond": time{"ms", "millisecond"},
}

// Format a number of milliseconds into a readable format
// 	- ms.Fmt(60000)            			// "1m"
// 	- ms.Fmt(2 * 60000)        			// "2m"
// 	- ms.Fmt(-3 * 60000)       			// "-3m"
// 	- ms.Fmt(ms('10 hours'))   			// "10h"
func Fmt(ms int) string {
	return format(ms, true)
}

// Format a number of milliseconds into a readable long format
// 	- ms.FmtLong(60000)                 // "1 minute"
// 	- ms.FmtLong(2 * 60000)             // "2 minutes"
// 	- ms.FmtLong(-3 * 60000)            // "-3 minutes"
// 	- ms.FmtLong(ms.Parse('10 hours'))  // "10 hours"
func FmtLong(ms int) string {
	return format(ms, false)
}

func format(ms int, short bool) string {
	abs := math.Abs(float64(ms))
	if abs >= d {
		t := int(math.Round(float64(ms) / d))
		return strconv.FormatInt(int64(t), 10) + suffix("day", t, short)
	}
	if abs >= h {
		t := int(math.Round(float64(ms) / h))
		return strconv.FormatInt(int64(t), 10) + suffix("hour", t, short)
	}
	if abs >= m {
		t := int(math.Round(float64(ms) / m))
		return strconv.FormatInt(int64(t), 10) + suffix("minute", t, short)
	}
	if abs >= s {
		t := int(math.Round(float64(ms) / s))
		return strconv.FormatInt(int64(t), 10) + suffix("second", t, short)
	}
	return strconv.FormatInt(int64(ms), 10) + suffix("millisecond", ms, short)
}

func suffix(s string, t int, short bool) string {
	if short {
		return suffixes[s].short
	}
	if math.Abs(float64(t)) > 1 {
		return " " + suffixes[s].long + "s"
	}
	return " " + suffixes[s].long
}

// Parse a time string into the number of milliseconds
// 	- ms.Parse('1y')     				// 31557600000
// 	- ms.Parse('100')    				// 100
// 	- ms.Parse('-3 days') 				// -259200000
// 	- ms.Parse('-1h')     				// -3600000
// 	- ms.Parse('-200')   				// -200
func Parse(input string) int {
	if len(input) > 100 {
		return 0
	}
	re := regexp.MustCompile(match)

	found := re.FindAllStringSubmatch(input, -1)

	if len(found) == 0 {
		return 0
	}
	numStr := found[0][1]
	unit := strings.ToLower(found[0][2])

	n, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0
	}
	if unit == "years" ||
		unit == "year" ||
		unit == "yrs" ||
		unit == "yr" ||
		unit == "y" {
		return int(n * y)
	}
	if unit == "weeks" ||
		unit == "week" ||
		unit == "w" {
		return int(n * w)
	}
	if unit == "days" ||
		unit == "day" ||
		unit == "d" {
		return int(n * d)
	}
	if unit == "hours" ||
		unit == "hour" ||
		unit == "hrs" ||
		unit == "hr" ||
		unit == "h" {
		return int(n * h)
	}
	if unit == "minutes" ||
		unit == "minute" ||
		unit == "mins" ||
		unit == "min" ||
		unit == "m" {
		return int(n * m)
	}
	if unit == "seconds" ||
		unit == "second" ||
		unit == "secs" ||
		unit == "sec" ||
		unit == "s" {
		return int(n * s)
	}
	if unit == "milliseconds" ||
		unit == "millisecond" ||
		unit == "msecs" ||
		unit == "msec" ||
		unit == "ms" {
		return int(n)
	}
	return int(n)
}
