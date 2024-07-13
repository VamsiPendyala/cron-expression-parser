package parser

import (
	"reflect"
	"testing"
)

func TestExpandCronField(t *testing.T) {
	tests := []struct {
		field    string
		min      int
		max      int
		expected []string
	}{
		{"*", 0, 59, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"}},
		{"*/15", 0, 59, []string{"0", "15", "30", "45"}},
		{"0", 0, 23, []string{"0"}},
		{"1,15", 1, 31, []string{"1", "15"}},
		{"1-5", 0, 6, []string{"1", "2", "3", "4", "5"}},
	}

	for _, test := range tests {
		result := expandCronField(test.field, test.min, test.max)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("expandCronField(%q, %d, %d) = %v; want %v", test.field, test.min, test.max, result, test.expected)
		}
	}
}

func TestParse(t *testing.T) {
	parser := NewParser()
	tests := []struct {
		cronString string
		expected   CronSchedule
		err        bool
	}{
		{"*/15 0 1,15 * 1-5 /usr/bin/find", CronSchedule{
			Minute:     []string{"0", "15", "30", "45"},
			Hour:       []string{"0"},
			DayOfMonth: []string{"1", "15"},
			Month:      []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
			DayOfWeek:  []string{"1", "2", "3", "4", "5"},
			Command:    "/usr/bin/find",
		}, false},
		{"invalid cron string", CronSchedule{}, true},
	}

	for _, test := range tests {
		result, err := parser.Parse(test.cronString)
		if (err != nil) != test.err {
			t.Errorf("Parse(%q) error = %v; want error = %v", test.cronString, err, test.err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Parse(%q) = %v; want %v", test.cronString, result, test.expected)
		}
	}
}
