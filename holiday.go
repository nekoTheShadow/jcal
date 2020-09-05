package main

import (
	"bufio"
	"errors"
	"strings"
	"time"

	"github.com/rakyll/statik/fs"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	_ "github.com/nekoTheShadow/jcal/statik"
)

type HolidayList struct {
	MaxYear  int
	MinYear  int
	Holidays []*Holiday
}

type Holiday struct {
	Date        time.Time
	Description string
}

func NewHolidayList() (*HolidayList, error) {
	fs, err := fs.New()
	if err != nil {
		return nil, err
	}

	file, err := fs.Open("/syukujitsu.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))
	scanner.Scan()
	scanner.Text() // 1行目はヘッダなので読み飛ばしてしまう

	holidays := []*Holiday{}
	maxYear := 0
	minYear := 9999
	for scanner.Scan() {
		cells := strings.Split(scanner.Text(), ",")
		date, err := time.Parse("2006/1/2", cells[0])
		if err != nil {
			return nil, err
		}
		holidays = append(holidays, &Holiday{Date: date, Description: cells[1]})

		if maxYear < date.Year() {
			maxYear = date.Year()
		}
		if date.Year() < minYear {
			minYear = date.Year()
		}
	}

	return &HolidayList{Holidays: holidays, MaxYear: maxYear, MinYear: minYear}, nil
}

func (h *HolidayList) GetHolidayDescription(date time.Time) (string, error) {
	for _, holiday := range h.Holidays {
		if holiday.Date.Year() == date.Year() &&
			holiday.Date.Month() == date.Month() &&
			holiday.Date.Day() == date.Day() {
			return holiday.Description, nil
		}
	}
	return "", errors.New(date.Format("2006-01-02") + " is not holiday")
}
