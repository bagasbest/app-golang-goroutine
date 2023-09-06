package app_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// pool adalah implementasi design pattern bernama object pool pattern
// sederhananya, design pattern pool ini digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil dari pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan ke poolnya
// biasa digunakan untuk membuat koneksi ke database
func TestPool(t *testing.T) {
	pool := sync.Pool{
		// set default pool value
		New: func() interface{} {
			return "New"
		},
	}
	group := sync.WaitGroup{}

	pool.Put("Bagas")
	pool.Put("Ganteng")
	pool.Put("Pangestu")

	for i := 0; i < 10; i++ {
		go func() {
			group.Add(1)
			data := pool.Get()
			fmt.Println(data)
			//time.Sleep(1 * time.Second)
			pool.Put(data)
			defer group.Done()
		}()
	}

	group.Wait()
	// time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
