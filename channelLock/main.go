package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// 執行該程式後會出現 程式卡死 或 deallock 錯誤
	// 請嘗試修正問題
	// 提示: 注意 chan 讀取的位置

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
			do(inputNumberChan, channelLock)
			wg.Done()
		}()
		inputNumberChan <- num
		log.Println("start:", num)
	}
	wg.Wait()
	log.Println("end")
}

func do(inputNumberChan <-chan int, channelLock <-chan bool) {
	num := <-inputNumberChan
	if num <= 20 {
		return
	}

	time.Sleep(1 * time.Second)

	log.Println("do:", num)
	<-channelLock
}
