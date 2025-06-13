package play

import (
	"fmt"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	now := time.Now()
	fmt.Println(now)

	t1 := now.Unix()
	fmt.Println(t1)
	now1 := time.Unix(t1, 0)
	fmt.Println(now1)

	z, x := now.Zone()
	fmt.Println(z, x)
}
