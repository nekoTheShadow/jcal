package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Cmd struct {
	holidayList *HolidayList
}

func NewCmd() (*Cmd, error) {
	holidayList, err := NewHolidayList()
	if err != nil {
		return nil, err
	}
	return &Cmd{holidayList: holidayList}, nil
}

func (c *Cmd) ParseArgs(args []string, today time.Time) (int, time.Month, error) {
	if len(args) == 0 {
		return today.Year(), today.Month(), nil
	}

	if len(args) == 2 {
		year, err := strconv.Atoi(args[0])
		if err != nil || year < c.holidayList.MinYear || c.holidayList.MaxYear < year {
			return -1, -1, fmt.Errorf("year (%v) is not between %v and %v", args[0], c.holidayList.MinYear, c.holidayList.MaxYear)
		}

		month, err := strconv.Atoi(args[1])
		if err != nil || month < 1 || 12 < month {
			return -1, -1, fmt.Errorf("month (%v) is not between 1 and 12", args[1])
		}

		return year, time.Month(month), nil
	}

	return -1, -1, errors.New("number of arguments is 0 or 2")
}

func (c *Cmd) CreateCalendar(year int, month time.Month) []string {
	days := []string{Red("日"), "月", "火", "水", "木", "金", Blue("土")}
	descriptions := []string{}
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	for date := start; date.Weekday() != time.Sunday; date = date.AddDate(0, 0, -1) {
		days = append(days, "  ")
	}
	for date := start; date.Month() == month; date = date.AddDate(0, 0, 1) {
		day := date.Format("02")
		description, err := c.holidayList.GetHolidayDescription(date)
		if err == nil {
			// 祝日の場合
			days = append(days, Red(day))
			descriptions = append(descriptions, date.Format("2006-01-02")+" "+description)
		} else {
			if date.Weekday() == time.Saturday {
				days = append(days, Blue(day))
			} else if date.Weekday() == time.Sunday {
				days = append(days, Red(day))
			} else {
				days = append(days, day)
			}
		}
	}

	lines := []string{start.Format("     2006-01")}
	for i := 0; i < len(days); i += 7 {
		lines = append(lines, strings.TrimRight(strings.Join(days[i:i+7], " "), " "))
	}
	if len(descriptions) > 0 {
		lines = append(lines, strings.Repeat("-", 20))
		for _, desdescription := range descriptions {
			lines = append(lines, desdescription)
		}
	}
	return lines
}

func (c *Cmd) CreateUsage() []string {
	return []string{
		"Usage: jcal [year month]",
		fmt.Sprintf("  - year : %v..%v", c.holidayList.MinYear, c.holidayList.MaxYear),
		"  - month: 1..12",
	}
}

func Red(s string) string {
	return "\033[31m" + s + "\033[m"
}

func Blue(s string) string {
	return "\033[34m" + s + "\033[m"
}
