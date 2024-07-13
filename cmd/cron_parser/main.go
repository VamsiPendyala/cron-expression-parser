package main

import (
	"flag"
	"fmt"
	"strings"

	"cron_parser/pkg/parser"
)

func main() {
	cronString := flag.String("cron", "", "Cron string to parse")
	flag.Parse()

	if *cronString == "" {
		fmt.Println("Please provide a cron string using the -cron flag.")
		return
	}

	p := parser.NewParser()
	schedule, err := p.Parse(*cronString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("minute        %s\n", strings.Join(schedule.Minute, " "))
	fmt.Printf("hour          %s\n", strings.Join(schedule.Hour, " "))
	fmt.Printf("day of month  %s\n", strings.Join(schedule.DayOfMonth, " "))
	fmt.Printf("month         %s\n", strings.Join(schedule.Month, " "))
	fmt.Printf("day of week   %s\n", strings.Join(schedule.DayOfWeek, " "))
	fmt.Printf("command       %s\n", schedule.Command)
}
