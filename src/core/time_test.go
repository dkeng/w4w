package core_test

import (
	"fmt"
	"testing"

	"github.com/dkeng/w4w/src/core"
)

func TestTodayStartEndTime(t *testing.T) {
	t1, t2 := core.TodayStartEndTime()
	fmt.Println(t1, t2)
}
