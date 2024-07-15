package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// CronParser interface
type CronParser interface {
	Parse(cronString string) (CronSchedule, error)
}

// CronSchedule struct
type CronSchedule struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

// Parser struct
type Parser struct{}

// NewParser creates a new Parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse method implementation
func (p *Parser) Parse(cronString string) (CronSchedule, error) {
	fields := strings.Fields(cronString)
	if len(fields) != 6 {
		return CronSchedule{}, fmt.Errorf("invalid cron string")
	}

	minute := expandCronField(fields[0], 0, 59)
	hour := expandCronField(fields[1], 0, 23)
	dayOfMonth := expandCronField(fields[2], 1, 31)
	month := expandCronField(fields[3], 1, 12)
	dayOfWeek := expandCronField(fields[4], 0, 6)
	command := fields[5]

	return CronSchedule{
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dayOfMonth,
		Month:      month,
		DayOfWeek:  dayOfWeek,
		Command:    command,
	}, nil
}

// expandCronField expands a cron field based on its value.
func expandCronField(field string, min int, max int) []string {
	var result []string

	if field == "*" {
		for i := min; i <= max; i++ {
			result = append(result, fmt.Sprintf("%d", i))
		}
	} else if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, _ := strconv.Atoi(parts[1])
		for i := min; i <= max; i += step {
			result = append(result, fmt.Sprintf("%d", i))
		}
	} else if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		for _, part := range parts {
			if strings.Contains(part, "-") {
				rangeParts := strings.Split(part, "-")
				start, _ := strconv.Atoi(rangeParts[0])
				end, _ := strconv.Atoi(rangeParts[1])
				if start < min {
					start = min
				}
				if end > max {
					end = max
				}
				for i := start; i <= end; i++ {
					result = append(result, fmt.Sprintf("%d", i))
				}
			} else {
				num, _ := strconv.Atoi(part)
				if num < min {
					num = min
				}
				if num > max {
					num = max
				}
				result = append(result, fmt.Sprintf("%d", num))
			}
		}
	} else if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, _ := strconv.Atoi(rangeParts[0])
		end, _ := strconv.Atoi(rangeParts[1])

		if start < min {
			start = min
		}
		if end > max {
			end = max
		}

		for i := start; i <= end; i++ {
			result = append(result, fmt.Sprintf("%d", i))
		}
	} else {
		num, _ := strconv.Atoi(field)
		result = append(result, fmt.Sprintf("%d", num))
	}

	return result
}
