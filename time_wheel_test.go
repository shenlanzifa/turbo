package turbo

import (
	// "fmt"
	// "sync"
	// "sync/atomic"
	"testing"
	"time"
)

func TestTimeWheel(t *testing.T) {
	tw := NewTimeWheel(100*time.Millisecond, 10, 100)

	for i := 0; i < 10; i++ {
		to := false
		ch := make(chan bool, 1)
		tw.After(3*time.Second, func() {
			ch <- true
		})
		select {
		case to = <-ch:
			to = true
		case <-time.After(5 * time.Second):
			to = false

		}
		if !to {
			t.Fail()
		} else {
			t.Logf("TestTimeWheel|SUCC|%d\n", i)
		}
	}
}

func BenchmarkTimeWheel(t *testing.B) {
	tw := NewTimeWheel(1*time.Second, 10, 100)

	for i := 0; i < t.N; i++ {

		timerId, _ := tw.After(3*time.Second, func() {

		})
		tw.Remove(timerId)
	}
}

func BenchmarkParalleTimeWheel(t *testing.B) {
	tw := NewTimeWheel(1*time.Second, 10, 100)
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			timerId, _ := tw.After(3*time.Second, func() {

			})
			tw.Remove(timerId)
		}
	})

}
