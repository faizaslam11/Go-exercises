// By Mohd Faiz Aslam
package main
import (
    "fmt"
    "time"
)

// Prints numbers from 1-3 along with the passed string
func foo(s string) {
    for i := 1; i <= 3; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, ": ", i)
    }
}

func main() {
    
    // Starting two goroutines
    go foo("1st goroutine (activated)")
    foo("2nd goroutine (deactivated)")
    go foo("3rd (activated)")

    // Wait for goroutines to finish before main goroutine ends
    time.Sleep(time.Second)
    fmt.Println("Main goroutine finished")
}
