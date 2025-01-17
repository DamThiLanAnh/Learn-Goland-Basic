package main

import (
	"fmt"
	"sync"
	"time"
)

func demoUnbufferedChannel() {
	var ch = make(chan int)
	go func() {
		ch <- 1
		fmt.Println("sent")
	}()
	time.Sleep(time.Second * 2)
	fmt.Printf("received value = %d \n", <-ch)
	time.Sleep(time.Second * 5)
}

func demobufferedChannel() {
	var ch = make(chan int, 2)

	go func() {
		ch <- 1
		fmt.Println("sent")
	}()
	time.Sleep(time.Second * 2)
	fmt.Printf("received value = %d \n", <-ch)
	time.Sleep(time.Second * 5)
}

func demoSelectChannel() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	ch1 <- 10
	ch2 <- 20

	select {
	case fromCh1 := <-ch1:
		fmt.Println("received value", fromCh1)
	case fromCh2 := <-ch2:
		fmt.Println("received value", fromCh2)
	}
}

func demoSelectChannelV2() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go func() {
		ch1 <- 10
	}()

	go func() {
		ch2 <- 20
	}()

	for i := range 2 {
		fmt.Println(i)
		select {
		case fromCh1 := <-ch1:
			fmt.Println("received value", fromCh1)
		case fromCh2 := <-ch2:
			fmt.Println("received value", fromCh2)
		}
	}
}

func demoSelectChannelV3() {
	var ch1 = make(chan int, 2)
	var ch2 = make(chan int, 2)

	ch1 <- 10
	ch2 <- 20

	for i := range 2 {
		fmt.Println(i)
		select {
		case fromCh1 := <-ch1:
			fmt.Println("received value", fromCh1)
		case fromCh2 := <-ch2:
			fmt.Println("received value", fromCh2)
		}
	}
}

var mutex sync.Mutex
var wg sync.WaitGroup //control so lg gorou team

func demoRaceCondition() {
	var count = 0
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(count)
	time.Sleep(time.Second * 5)
}

func main() {
	demoRaceCondition()
}
