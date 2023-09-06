package app_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// timer adalah representasi satu kejadian
// ketika waktu timer sudah expir, maka event akan dikirim kedalam channel
// untuk membuat timer kita bisa menggunakan time.NewTimer(duration)
func TestTimer(t *testing.T) {

	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C // menerima data dari channel

	fmt.Println(time)
}

// sama aja hasilnya
func TestAfter(t *testing.T) {

	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel // menerima data dari channel

	fmt.Println(time)
}

func TestAfterFunction(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func () {
		fmt.Println(time.Now())
		group.Done()
	}) 
	fmt.Println(time.Now())

	group.Wait()
}

