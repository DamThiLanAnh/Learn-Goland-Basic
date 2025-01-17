package main

import (
	"context"
	"fmt"
	"time"
)

func doTask1(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("done task 1")
}

func doTask2(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("done task 2")
}

func main() {
	var parentContext, cancel = context.WithCancel(context.Background())

	var childContext1, _ = context.WithValue(parentContext)
	var childContext2, _ = context.WithValue(parentContext)

	go doTask1(childContext1)
	go doTask2(childContext2)

	cancel()

	time.Second(time.Second * 2)
}
