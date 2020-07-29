package main

import (
	"log"
	"sync"
	"time"
)

func answer() {
	// 解答篇

	var wg sync.WaitGroup

	// channelLock 數量限制鎖
	channelLock := make(chan bool, 20)
	defer close(channelLock)

	inputNumberChan := make(chan int, 20)
	defer close(inputNumberChan)

	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	for _, num := range numList {
		log.Println("go:", num)

		wg.Add(1)
		channelLock <- true
		go func() {
			answerDo(inputNumberChan)
			// 當 function 執行完畢時務必讀取 chan，才會釋放 chan 空間
			<-channelLock
			wg.Done()
		}()
		inputNumberChan <- num
		log.Println("start:", num)
	}
	wg.Wait()
	log.Println("end")
}

func answerDo(inputNumberChan <-chan int) {
	num := <-inputNumberChan
	if num <= 20 {
		return
	}

	time.Sleep(1 * time.Second)

	log.Println("do:", num)
}
