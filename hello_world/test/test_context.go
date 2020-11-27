package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	 // defer cancel()
	d := time.Now().Add(time.Second*1)
	fmt.Println("%v| %v", d, time.Now())
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	// go handle(ctx, 500*time.Millisecond) // case 1
	go handle(ctx, 200*time.Millisecond) // case 2
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
		fmt.Println("default")

	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	//default:
	//	fmt.Println("with cancelkkkkkkkkk")
	//	context.WithCancel(ctx)
	}
}


/*
context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来；
context.TODO 应该只在不确定应该使用哪种上下文时使用；
*/