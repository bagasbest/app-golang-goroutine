package app_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {

	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

// Sync maps lebih aman untuk menghindari race condition
// sama seperti map pada umumnya ada method .Store(key, value), .Load(key), .Delete(key), dan .Range(key, value) untuk iterasi
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
