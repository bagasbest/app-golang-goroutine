package app_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// golang memiliki package bernama sync/atomic
// atomic metupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
// contohnya jika sebelumnya menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya bisa di gunakan menggunakan Atomic package
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}


