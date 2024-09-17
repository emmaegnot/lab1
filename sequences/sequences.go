package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
	return slice
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	for i := 0; i < len(array); i++ {
		array[i] = f(array[i])
		fmt.Println()
	}
	return array
}

func main() {
	var mySlice = []int{1, 2, 3}
	fmt.Println(mapSlice(addOne, mySlice))
	var myArray = [3]int{1, 2, 3}
	fmt.Println(mapArray(addOne, myArray))
}
