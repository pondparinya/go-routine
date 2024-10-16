package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func main() {
	data := make([]interface{}, 10000)
	pool, err := ants.NewPool(200)
	if err != nil {
		panic(err)
	}
	defer pool.Release()
	var wg sync.WaitGroup
	tn := time.Now()
	for i := range data {
		wg.Add(1)
		if err := pool.Submit(func() {
			defer wg.Done()
			process(i)
		}); err != nil {
			fmt.Printf("error : %v", err)
		}
	}
	wg.Wait()
	fmt.Printf("run time : %.2fs\n", time.Since(tn).Seconds())
	fmt.Printf("running goroutines: %d\n", pool.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}

func process(data int) {
	time.Sleep(100 * time.Millisecond)
	atomic.AddInt32(&sum, int32(data))
	fmt.Printf("run with %d\n", data)
}
