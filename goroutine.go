package main

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

func main() {

	var numSize int
	var userInput int
	var waitGroup sync.WaitGroup
	sliceOfInts := make([]int, 0, 3)

	fmt.Println("How many numbers would you like to add to the array?")
	fmt.Scan(&numSize)

	for i := 0; i < numSize; i++ {

		fmt.Println("Enter a number to fill inside array")
		_, err := fmt.Scan(&userInput)
		if err != nil {
			log.Fatal(err)
			fmt.Println("Invalid user input")
		}

		sliceOfInts = append(sliceOfInts, userInput)

	}

	sliceSize := numSize / 4
	slice1 := sliceOfInts[:sliceSize]
	slice2 := sliceOfInts[sliceSize : 2*(sliceSize)]
	slice3 := sliceOfInts[2*(sliceSize) : 3*(sliceSize)]
	slice4 := sliceOfInts[3*(sliceSize):]

	fmt.Println("Arrays after partitioned", slice1, slice2, slice3, slice4)

	waitGroup.Add(4)
	go sortList(slice1)
	waitGroup.Done()
	go sortList(slice2)
	waitGroup.Done()
	go sortList(slice3)
	waitGroup.Done()
	go sortList(slice4)
	waitGroup.Done()
	waitGroup.Wait()

	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)

	fmt.Println("Your Slice merged, sorted and printed is as follows:", newSlice)

}

func sortList(unsortedSlice []int) []int {
	sort.Ints(unsortedSlice)
	return unsortedSlice

}

func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
	newSlice := []int{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Ints(newSlice)
	return newSlice

}
