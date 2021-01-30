package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(20, 30, 3, 4)

	fmt.Println(total)

	slice := []int{20,10,4,60}
	total = sumAll(slice...)

	fmt.Println(total)
}
