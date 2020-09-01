package main

import (
	"testing"
	"time"
)

func Test_GetHolidayDescriptionは与えられた日がどの国民の祝日にあたるのかを判定する(t *testing.T) {
	t.Run("2020-11-23は勤労感謝の日", func(t *testing.T) {
		description, err := GetHolidayDescription(time.Date(2020, time.November, 23, 0, 0, 0, 0, time.UTC))
		if description != "勤労感謝の日" {
			t.Fatalf("Reason: %v", err)
		}
	})

	t.Run("2020-11-24は国民の祝日ではない", func(t *testing.T) {
		description, err := GetHolidayDescription(time.Date(2020, time.November, 24, 0, 0, 0, 0, time.UTC))
		if err == nil {
			t.Fatalf("expected: 国民の祝日ではない, but: %s", description)
		}
	})
}
