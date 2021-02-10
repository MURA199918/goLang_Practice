package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome bubble sort program")
	size := 0
	fmt.Println("Enter Number of inputs")
	fmt.Scanln(&size)
	fmt.Println("Enter inputs: ")
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&numbers[i])
	}
	//arrayone := []int{3, 4, 5, 2, 1}
	fmt.Println("Before sorting: ")
	fmt.Println(numbers)
	bubbleSort(numbers)
}

func bubbleSort(input []int) {
	length := len(input)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if input[j] > input[j+1] {
				swapNumber(&input[j], &input[j+1])
			}
		}
	}
	fmt.Println("Bubble sorted array: ")
	for _, val := range input {
		fmt.Println(val)
	}
}

func swapNumber(x *int, y *int) {
	*x, *y = *y, *x
}
