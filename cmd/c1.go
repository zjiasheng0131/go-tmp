package main

import (
	"fmt"
	"time"
)

// 父将子的cancel
// func main() {   // 父将子的cancel
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// 启动一个 goroutine
// 	go func(ctx context.Context) {
// 		//	c2 := context.WithValue(ctx, "myvalues", 111)
// 		c2, _ := context.WithCancel(ctx)
// 		fmt.Println("go in")
// 		select {
// 		case <-time.After(5 * time.Second):
// 			doWork() // 任务执行
// 			fmt.Println("Task completed")
// 		case <-c2.Done():
// 			fmt.Println("Task cancelled or timed out")
// 		}
// 	}(ctx)

// 	time.Sleep(2 * time.Second)
// 	fmt.Println("try to cancel")
// 	cancel()
// 	fmt.Println("cancel over")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Main function finished")
// }

// func main() { // 子能cancel父吗？ 不行
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// 启动一个 goroutine

// 	go func(ctx context.Context) {
// 		//	c2 := context.WithValue(ctx, "myvalues", 111)
// 		c2, can2 := context.WithCancel(ctx)
// 		fmt.Println("go in")
// 		select {
// 		case <-time.After(1 * time.Second):
// 			can2()
// 			fmt.Println("can2 close")

// 			// doWork() // 任务执行
// 			// fmt.Println("Task completed")
// 		case <-c2.Done():
// 			fmt.Println("sub Task cancelled or timed out")
// 		}
// 	}(ctx)

// 	select {

// 	case <-ctx.Done():
// 		fmt.Println("main Task cancelled or timed out")
// 	}

// 	fmt.Println("Main function finished")
// }

func doWork() {
	fmt.Println("doWork start")
	time.Sleep(5 * time.Second)
	fmt.Println("doWork end")
}
