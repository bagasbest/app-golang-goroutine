package app_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1) /// running 1 proses asynchronus

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronus(group)
	}

	/// gak perlu time.Sleep() lagi

	group.Wait()
	fmt.Println("Selesai")
}

