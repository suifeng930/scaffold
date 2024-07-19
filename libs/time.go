package libs

import (
	"strconv"
	"strings"
	"time"
)

// OneDayBeginAndEndTimeStamp 获取天的起始时间
func OneDayBeginAndEndTimeStamp(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day()+1, 0, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix()
}

// OneMonthBeginAndEndTimeStamp 获取月的起始时间
func OneMonthBeginAndEndTimeStamp(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), 1, 0, 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month()+1, 1, 0, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix()
}

func GetDtByOffset(tm time.Time, offset int) (r int) {
	r, _ = strconv.Atoi(tm.AddDate(0, 0, offset).Format("20060102"))
	return
}

// ParseDate 通过秒时间戳获取日期
func ParseDate(timer int64) string {
	if timer == 0 {
		return "--"
	}
	return time.Unix(timer, 0).Format("2006-01-02")
}

// ParseMonth 通过秒时间戳获取日期
func ParseMonth(_timer int64) string {
	fmtString := time.Unix(_timer, 0).Format("2006-01-02 15:04:05")
	date := strings.Split(strings.Split(fmtString, " ")[0], "-")
	return date[0] + "-" + date[1]
}

// DayAndMonthStart 根据时间戳(秒)获取当天的起始时间和当月的起始时间
func DayAndMonthStart(t int64) (dayStart, monthStart int64) {
	oTime := time.Unix(t, 0)
	dayStart = time.Date(oTime.Year(), oTime.Month(), oTime.Day(), 0, 0, 0, 0, oTime.Location()).Unix()
	monthStart = time.Date(oTime.Year(), oTime.Month(), 1, 0, 0, 0, 0, oTime.Location()).Unix()
	return
}

func MonthDayNum(t int64) (dayNum int64) {
	oTime := time.Unix(t, 0)
	start := time.Date(oTime.Year(), oTime.Month(), 1, 0, 0, 0, 0, oTime.Location())
	end := time.Date(oTime.Year(), oTime.Month()+1, 1, 0, 0, 0, 0, oTime.Location())
	dayNum = int64(end.Sub(start).Hours() / 24)
	return
}

// GetHourStAndEt 获取小时的起始时间和结束时间
func GetHourStAndEt(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour(), 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour()+1, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix()
}

func Get2HourStAndEt(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour()-1, 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour()+1, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix()
}

func Get2HourTimeStampAfterStart(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour(), 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), beginTime.Hour()+2, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix()
}

// YesterdayBeginAndEndTimeStamp 获取前一天的起始时间
func YesterdayBeginAndEndTimeStamp(t int64) (int64, int64) {
	beginTime := time.Unix(t, 0)
	timeBegin := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day()-1, 0, 0, 0, 0, beginTime.Location())
	timeEnd := time.Date(beginTime.Year(), beginTime.Month(), beginTime.Day(), 0, 0, 0, 0, beginTime.Location())
	return timeBegin.Unix(), timeEnd.Unix() - 1
}
