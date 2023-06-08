//channels: goroutines can communicates with each another
// we can use channels to get data from 1 goroutine to another

//2 operations ; SEND and RECIEVE  AND (close)
// both should be of same data type

//TRYING FIBONACCI WITH CHANNELS

package main

import "fmt"

func fibonacci(max int, ch chan int) {
	fib := make([]int, max)
	fib[0] = 0
	fib[1] = 1
	ch <- fib[0]
	ch <- fib[1]

	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
		ch <- fib[i]
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibonacci(20, ch)
	for msg := range ch {
		fmt.Println(msg)

	}
	// councurency fibonacci
	// func fibonacci(n int) int {
	// 	if n == 0 {
	// 		return 0
	// 	} else if n == 1 {
	// 		return 1
	// 	}
	// 	return fibonacci(n-1) + fibonacci(n-2)
	// }

}
