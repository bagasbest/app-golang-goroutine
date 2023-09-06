package app_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

// ticker adalah representasi kejadian yang berulang
// ketika waktu tickersudah expire, maka event akan di kirim ke dalam  kedalam channl
// untukl membuat ticker, bisa dnegan menggunakan time.NewTicker(duratioin)
// untuk menghentikjan ticker bisa diengan menggunakan ticker.Stop()
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()


	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}

}

