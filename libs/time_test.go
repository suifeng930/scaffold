package libs

import (
	"fmt"
	"testing"
	"time"
)

func TestGetOneDayTimeStamp(t *testing.T) {
	nowStamp := time.Now().Unix()
	start, end := OneDayBeginAndEndTimeStamp(nowStamp)
	fmt.Println("now: ", nowStamp)
	fmt.Println("start: ", start)
	fmt.Println("end: ", end)
}

func TestGetOneDayTimeString(t *testing.T) {
	nowStamp := time.Now().Unix()
	timeString := ParseDate(nowStamp)
	fmt.Println("now: ", nowStamp)
	fmt.Println("timeString: ", timeString)

	timeV2 := ParseMonth(nowStamp)
	fmt.Println("time V2:", timeV2)
}
