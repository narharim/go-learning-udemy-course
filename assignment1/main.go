package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Printf("Number is even %d\n", n)
		} else {
			fmt.Printf("Number is odd %d\n", n)
		}
	}

}
