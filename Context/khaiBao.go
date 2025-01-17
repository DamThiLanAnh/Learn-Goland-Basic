package main

import (
	"context"
	"fmt"
	"time"
)

// 1. Tạo một contex với time out = 3s
// 2. Tạo một hàm sd context trên, hàm này sẽ trả về kq

func task(ctx context.Context) (int, error) {
	chResult := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("task running #", i)
			time.Sleep(time.Second)
		}
		chResult <- 10
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-ch:
		return &result, nil
	}
}
func main() {
	var ctx, cancel = context.WithCancel(context.Background(), time.Second*12)
	defer cancel()

	result, err := task(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*result)
}
