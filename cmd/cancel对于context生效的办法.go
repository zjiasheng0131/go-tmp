package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		go func() {
			defer close(dst)
			n := 1
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					fmt.Println(1111, n)
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}
