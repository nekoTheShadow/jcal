package main

import (
	"strconv"
	"testing"
	"time"
)

func Test_CreateCalendarは指定した年月のカレンダを所定の形式で作成する(t *testing.T) {
	t.Run("国民の祝日が含まれる場合(2020/09)", func(t *testing.T) {
		expected := []string{
			"     2020-09",
			"\033[31m日\033[m 月 火 水 木 金 \033[34m土\033[m",
			"      01 02 03 04 \033[34m05\033[m",
			"\033[31m06\033[m 07 08 09 10 11 \033[34m12\033[m",
			"\033[31m13\033[m 14 15 16 17 18 \033[34m19\033[m",
			"\033[31m20\033[m \033[31m21\033[m \033[31m22\033[m 23 24 25 \033[34m26\033[m",
			"\033[31m27\033[m 28 29 30",
			"--------------------",
			"2020-09-21 敬老の日",
			"2020-09-22 秋分の日",
		}
		actual := CreateCalendar(2020, time.September)
		assert(t, expected, actual)
	})

	t.Run("国民の祝日が含まれない場合(2020/06)", func(t *testing.T) {
		expected := []string{
			"     2020-06",
			"\033[31m日\033[m 月 火 水 木 金 \033[34m土\033[m",
			"   01 02 03 04 05 \033[34m06\033[m",
			"\033[31m07\033[m 08 09 10 11 12 \033[34m13\033[m",
			"\033[31m14\033[m 15 16 17 18 19 \033[34m20\033[m",
			"\033[31m21\033[m 22 23 24 25 26 \033[34m27\033[m",
			"\033[31m28\033[m 29 30",
		}
		actual := CreateCalendar(2020, time.June)
		assert(t, expected, actual)
	})
}

func assert(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Fatalf("len(expected)=%v, but len(actual)=%v", len(expected), len(actual))
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Fatalf("Line No.%d - expected: %v, actual: %v", i+1, expected[i], actual[i])
		}
	}
}

func Test_ParseArgsは実行時引数から対象の年月を取得する(t *testing.T) {
	today := time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC)

	t.Run("実行時引数が与えられない場合、コマンド実行時を対象の年月とする", func(t *testing.T) {
		year, month, err := ParseArgs([]string{}, today)
		if !(year == 2020 && month == time.April && err == nil) {
			t.Errorf("ParseArgs([], Date(2020, April)) ... expected (2020, April, nil), but actual (%v, %v, %v)", year, month, err)
		}
	})

	t.Run("実行時引数が2つの場合は、第1引数を年、第2引数を月とする", func(t *testing.T) {
		mid := (MinYear() + MaxYear()) / 2

		t.Run("年はMinYear以上MaxYear以下の整数のみ、月は1以上12以下の整数のみそれぞれ有効", func(t *testing.T) {
			okCases := []struct {
				args  []string
				year  int
				month time.Month
			}{
				{args: []string{strconv.Itoa(MinYear()), "6"}, year: MinYear(), month: time.June},
				{args: []string{strconv.Itoa(mid), "6"}, year: mid, month: time.June},
				{args: []string{strconv.Itoa(MaxYear()), "6"}, year: MaxYear(), month: time.June},
				{args: []string{strconv.Itoa(mid), "1"}, year: mid, month: time.January},
				{args: []string{strconv.Itoa(mid), "12"}, year: mid, month: time.December},
			}
			for _, testcase := range okCases {
				year, month, err := ParseArgs(testcase.args, today)
				if year != testcase.year || month != testcase.month {
					t.Errorf("ParseArgs(%v) ... expected (%v, %v, nil), but actual (%v, %v, %v)", testcase.args, testcase.year, testcase.month, year, month, err)
				}
			}
		})

		t.Run("無効な年月日が与えられた場合はエラーとする", func(t *testing.T) {
			ngArgsSlice := [][]string{
				{strconv.Itoa(MinYear() - 1), "6"},
				{strconv.Itoa(MaxYear() + 1), "6"},
				{strconv.Itoa(mid), "0"},
				{strconv.Itoa(mid), "13"},
				{"A", "6"},
				{strconv.Itoa(mid), "A"},
			}
			for _, args := range ngArgsSlice {
				year, month, err := ParseArgs(args, today)
				if err == nil {
					t.Errorf("ParseArgs(%v) ... expected to throw error, but actual (%v, %v, %v)", args, year, month, err)
				}
			}
		})
	})

	t.Run("実行時引数が1つだけ、もしくは、3つ以上与えられた場合、エラーとする", func(t *testing.T) {
		argsSlice := [][]string{{"1"}, {"1", "2", "3"}}
		for _, args := range argsSlice {
			year, month, err := ParseArgs(args, today)
			if err == nil {
				t.Errorf("ParseArgs(%v) ... expected to throw error, but actual (%v, %v, %v)", args, year, month, err)
			}
		}
	})
}
