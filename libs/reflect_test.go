package libs

import (
	"fmt"
	"testing"
	"time"
)

func TestReflectStruct(t *testing.T) {

	type User struct {
		Name   string
		Age    int
		Gender bool
		Date   time.Time
	}
	type People struct {
		Name   string
		Age    int
		Gender bool
		Date   time.Time
	}
	sourceStruct := User{
		Name:   "Ma",
		Age:    20,
		Gender: false,
		Date:   time.Now(),
	}
	targetStruct := People{}
	fmt.Println("assign user before: ", sourceStruct)
	fmt.Println("assign people before: ", targetStruct)
	Assign(&sourceStruct, &targetStruct)
	fmt.Println("assign user after: ", sourceStruct)
	fmt.Println("assign people after: ", targetStruct)
}
