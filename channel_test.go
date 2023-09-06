package app_golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// channel <- "Bagas" /// mengirim data ke channel
	// data := <- channel /// menerima data dari channel
	// fmt.Println(<-channel) /// atau secara langsung

	defer close(channel) /// close channel  untuk menghindari memory leak

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Bagas Pangestu Ganteng" /// mengirim data ke channel
		fmt.Println("Selesai mengirim Data ke Channel")
	}()

	data := <-channel /// menerima data dari channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

/// vchannel bisa di gunakan untuk mengirim dan menerima data

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Bagas Pangestu"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Bagas Pangestu"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

/// Buffer channel, untuk menampung data antrean di channel
/// buffer = penyimpanan di dalam channel
/// buffer capacity = 5; jika kita mengirim  data ke 6 maka harus nunggu sampai buffer ada yang kosong

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Bagas"
		channel <- "Pangestu"
		channel <- "Ganteng"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

/// terkadang ada kasus channel di kirim secara terus menerus oelh pengirim
/// dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
/// salah satu yang bisa kita lakukan adalah denghan menggunakan perulangan range ketika menerima data dari channel
/// ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
/// ini lebih sederhana dari pada melakukan pengecekan channel secara manual

func TestRangeChanne(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

/// kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
/// lalu kita ingin mendapatkan data dari semua channel tersebut
/// untuk melakuikan hal tersebut, kita bisa menggunak select channel  di golang
/// dengan selech channel kita bisa memilih data tercepat dari beberapa channel, jika datange di beberapa channel, maka akan dipilih secara random

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}

func TestDefaelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}
	}

}

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

