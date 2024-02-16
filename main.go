package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quicksort(type_pivot string, arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var pivot int

	switch {
	case type_pivot == "firtst":
		pivot = arr[0]
	case type_pivot == "middle":
		pivot = arr[len(arr)/2]
	case type_pivot == "random":
		pivot = arr[rand.Intn(len(arr))]
	case type_pivot == "last":
		pivot = arr[len(arr)-1]
	}

	var less, equal, greater []int

	for _, value := range arr {
		switch {
		case value < pivot:
			less = append(less, value)
		case value == pivot:
			equal = append(equal, value)
		case value > pivot:
			greater = append(greater, value)
		}
	}
	
	fmt.Println("")
	fmt.Println("|Array Sorted:", arr)

	fmt.Printf("|Pivô: %v | Menor: %v | Igual: %v | Maior: %v\n", pivot, less, equal, greater)

	return append(append(quicksort(type_pivot, less), equal...), quicksort(type_pivot, greater)...)
}

func randomArray(n int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		randomNumber := rand.Intn(100)
		arr[i] = randomNumber
	}
	return arr
}

func main() {
	type_pivot := "random" // firtst, middle, random, last
	arr := randomArray(rand.Intn(50))
	fmt.Println("Original:", arr)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Start Quicksort")
	fmt.Println("Type Pivot:", type_pivot)
	sortedArr := quicksort(type_pivot, arr)
	fmt.Println("Sorted:", sortedArr)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Start Quicksort")
	fmt.Println("Type Pivot:", "firtst")
	sortedArr = quicksort("firtst", arr)
	fmt.Println("Sorted:", sortedArr)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Start Quicksort")
	fmt.Println("Type Pivot:", "middle")
	sortedArr = quicksort("middle", arr)
	fmt.Println("Sorted:", sortedArr)
	fmt.Println("")
	fmt.Println("--------------------")
	fmt.Println("Start Quicksort")
	fmt.Println("Type Pivot:", "last")
	sortedArr = quicksort("last", arr)
	fmt.Println("Sorted:", sortedArr)
}
