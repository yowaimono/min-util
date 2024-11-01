package 

import (
	"fmt"
	"time"
)

// 返回年份切片，例如 ["2024", "2025"]
func GetYears(times []time.Time) []string {
	years := make([]string, len(times))
	for i, t := range times {
		years[i] = fmt.Sprintf("%d", t.Year())
	}
	return years
}

// 返回年份和月份切片，例如 ["2024-10", "2025-11"]
func GetYearMonths(times []time.Time) []string {
	yearMonths := make([]string, len(times))
	for i, t := range times {
		yearMonths[i] = t.Format("2006-01")
	}
	return yearMonths
}

// 返回年份、月份和日期切片，例如 ["2024-10-01", "2025-11-02"]
func GetYearMonthDays(times []time.Time) []string {
	yearMonthDays := make([]string, len(times))
	for i, t := range times {
		yearMonthDays[i] = t.Format("2006-01-02")
	}
	return yearMonthDays
}

// 返回年份和月份切片（中文），例如 ["2024年10月", "2025年11月"]
func GetYearMonthsChinese(times []time.Time) []string {
	yearMonthsChinese := make([]string, len(times))
	for i, t := range times {
		yearMonthsChinese[i] = fmt.Sprintf("%d年%d月", t.Year(), t.Month())
	}
	return yearMonthsChinese
}

// 返回年份切片（中文），例如 ["2024年", "2025年"]
func GetYearsChinese(times []time.Time) []string {
	yearsChinese := make([]string, len(times))
	for i, t := range times {
		yearsChinese[i] = fmt.Sprintf("%d年", t.Year())
	}
	return yearsChinese
}

// 返回年份、月份和日期切片（中文），例如 ["2024年10月10日", "2025年11月11日"]
func GetYearMonthDaysChinese(times []time.Time) []string {
	yearMonthDaysChinese := make([]string, len(times))
	for i, t := range times {
		yearMonthDaysChinese[i] = fmt.Sprintf("%d年%d月%d日", t.Year(), t.Month(), t.Day())
	}
	return yearMonthDaysChinese
}
