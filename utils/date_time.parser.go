package utils

import (
	"fmt"
	"time"
)

func GetFormattedNow(now time.Time) string {
	year := now.Year()
	month := parseDigit(int(now.UTC().Month()))
	day := parseDigit(now.Day())
	hour := parseDigit(now.Hour())
	minute := parseDigit(now.Minute())
	second := parseDigit(now.Second())

	result := fmt.Sprintf("%d%s%s%s%s%s%d", year, month, day, hour, minute, second, now.UnixMilli())
	return result
}

func parseDigit(num int) string {
	if num < 10 {
		return fmt.Sprintf("0%d", num)
	}
	return fmt.Sprintf("%d", num)
}
