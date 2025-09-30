package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	return append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, v := range slice {
		slice[i] = f(v)
	}
	fmt.Println(slice)
}

func mapArray(f func(a int) int, array [3]int) {
	for i, v := range array {
		array[i] = f(v)
	}
	fmt.Println(array)
}

func main() {

	intsSlice := []int{2, 3, 4, 5, 6}
	newSlice := intsSlice[1:3]
	mapSlice(addOne, intsSlice)
	fmt.Println("initsSlice =", intsSlice)
	mapSlice(square, newSlice)
	fmt.Println("initsSlice =", intsSlice)
	fmt.Println("newSlice =", newSlice)

	initArray := [3]int{1, 2, 3}
	mapArray(addOne, initArray)
	fmt.Println("initArray =", initArray)

	intsSlice = double(intsSlice)
	fmt.Println(intsSlice)
}
