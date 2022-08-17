package main

import (
	"sync"
	"time"
	"fmt"
	"context"
)

var wg sync.WaitGroup

type Option struct {
	Interval time.Duration
}

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func mockReq(ctx context.Context, name string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			op := ctx.Value("myOptions").(*Option)
			fmt.Printf("\n%v", *op)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	vCtx := context.WithValue(ctx, "myOptions", &Option{2})
	
	go mockReq(vCtx, "worker1")
	go mockReq(vCtx, "worker2")
	
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)

}