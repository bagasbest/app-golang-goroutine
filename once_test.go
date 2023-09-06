package app_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

// / sync.Once adalah fitur yang di gunakan untuk memastikan bahasa sebuah function di eksekusi hanya sekali
// / jadi berapa banyak pun goroutine yang memgakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi function nya
// / goroutine yang lain akan di hiraukan, artinya function tidak akan di eksekusi lagi
func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce) /// hanya sekali di jalankan
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}


