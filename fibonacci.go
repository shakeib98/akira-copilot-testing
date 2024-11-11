// fibonacci.go

package main

import "fmt"

// Fibonacci function returns a slice containing the Fibonacci sequence up to n terms
func Fibonacci(n int) []int {
    if n <= 0 {
        return []int{}
    }
    fibSeq := make([]int, n)
    fibSeq[0] = 0
    if n > 1 {
        fibSeq[1] = 1
        for i := 2; i < n; i++ {
            fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
        }
    }
    return fibSeq
}

func main() {
    n := 10 // Example: Get the first 10 Fibonacci numbers
    fibSeq := Fibonacci(n)
    fmt.Println(fibSeq)



asdasdasd {} {P{ }P{}
