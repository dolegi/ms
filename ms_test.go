package ms

import (
	"testing"
)

var testFmtTests = []struct {
	title    string
	input    int
	expected string
}{
	{"500 milliseconds", 500, "500ms"},
	{"1 second", 1000, "1s"},
	{"1.2 seconds", 1200, "1s"},
	{"10 seconds", 10000, "10s"},
	{"1 minute", 60 * 1000, "1m"},
	{"10 minutes", 10 * 60 * 1000, "10m"},
	{"1 hour", 60 * 60 * 1000, "1h"},
	{"10 hours", 10 * 60 * 60 * 1000, "10h"},
	{"minus 10 hours", -10 * 60 * 60 * 1000, "-10h"},
	{"1 day", 24 * 60 * 60 * 1000, "1d"},
	{"10 days", 10 * 24 * 60 * 60 * 1000, "10d"},
	{"rounds", 234234234, "3d"},
	{"minus rounds", -234234234, "-3d"},
}

func TestFmt(t *testing.T) {
	for _, test := range testFmtTests {
		t.Run(test.title, func(t *testing.T) {
			actual := Fmt(test.input)
			if actual != test.expected {
				t.Errorf("TestFmt %s\n actual: %v\n expected: %v\n", test.title, actual, test.expected)
			}
		})
	}
}

var testFmtLongTests = []struct {
	title    string
	input    int
	expected string
}{
	{"500 milliseconds", 500, "500 milliseconds"},
	{"1 second", 1000, "1 second"},
	{"1.2 seconds", 1200, "1 second"},
	{"10 seconds", 10000, "10 seconds"},
	{"1 minute", 60 * 1000, "1 minute"},
	{"10 minutes", 10 * 60 * 1000, "10 minutes"},
	{"1 hour", 60 * 60 * 1000, "1 hour"},
	{"minus 1 hour", -60 * 60 * 1000, "-1 hour"},
	{"10 hours", 10 * 60 * 60 * 1000, "10 hours"},
	{"1 day", 24 * 60 * 60 * 1000, "1 day"},
	{"10 days", 10 * 24 * 60 * 60 * 1000, "10 days"},
	{"rounds", 234234234, "3 days"},
}

func TestFmtLong(t *testing.T) {
	for _, test := range testFmtLongTests {
		t.Run(test.title, func(t *testing.T) {
			actual := FmtLong(test.input)
			if actual != test.expected {
				t.Errorf("TestFmtLong %s\n actual: %v\n expected: %v\n", test.title, actual, test.expected)
			}
		})
	}
}

var testParseTests = []struct {
	title    string
	input    string
	expected int
}{
	{"500 milliseconds", "500ms", 500},
	{"1 second", "1s", 1000},
	{"1.2 seconds", "1.2s", 1200},
	{"10 seconds", "10s", 10000},
	{"1 minute", "1m", 60 * 1000},
	{"10 minutes", "10m", 10 * 60 * 1000},
	{"1 hour", "1h", 60 * 60 * 1000},
	{"10 hours", "10h", 10 * 60 * 60 * 1000},
	{"1 day", "1d", 24 * 60 * 60 * 1000},
	{"10 days", "10d", 10 * 24 * 60 * 60 * 1000},
	{"minus 10 days", "-10d", -10 * 24 * 60 * 60 * 1000},
	{"rounds", "2.742345435d", 236938645},
}

func TestParse(t *testing.T) {
	for _, test := range testParseTests {
		t.Run(test.title, func(t *testing.T) {
			actual := Parse(test.input)
			if actual != test.expected {
				t.Errorf("TestFmtLong %s\n actual: %v\n expected: %v\n", test.title, actual, test.expected)
			}
		})
	}
}

var testParseLongTests = []struct {
	title    string
	input    string
	expected int
}{
	{"500 milliseconds", "500 milliseconds", 500},
	{"1 second", "1 second", 1000},
	{"1.2 seconds", "1.2 second", 1200},
	{"10 seconds", "10 seconds", 10000},
	{"1 minute", "1 minute", 60 * 1000},
	{"10 minutes", "10 minutes", 10 * 60 * 1000},
	{"1 hour", "1 hour", 60 * 60 * 1000},
	{"minus 1 hour", "-1 hour", -60 * 60 * 1000},
	{"10 hours", "10 hours", 10 * 60 * 60 * 1000},
	{"1 day", "1 day", 24 * 60 * 60 * 1000},
	{"10 days", "10 days", 10 * 24 * 60 * 60 * 1000},
	{"rounds", "2.742345435days", 236938645},
}

func TestParseLong(t *testing.T) {
	for _, test := range testParseLongTests {
		t.Run(test.title, func(t *testing.T) {
			actual := Parse(test.input)
			if actual != test.expected {
				t.Errorf("TestParseLong %s\n actual: %v\n expected: %v\n", test.title, actual, test.expected)
			}
		})
	}
}

var testParseInvalidTests = []struct {
	title    string
	input    string
	expected int
}{
	{"Empty string", "", 0},
	{"Unknown unit", "1t", 0},
	{"No unit", "10", 10},
}

func TestParseInvalid(t *testing.T) {
	for _, test := range testParseInvalidTests {
		t.Run(test.title, func(t *testing.T) {
			actual := Parse(test.input)
			if actual != test.expected {
				t.Errorf("TestParseInvalid %s\n actual: %v\n expected: %v\n", test.title, actual, test.expected)
			}
		})
	}
}
