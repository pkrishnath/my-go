package main

import "fmt"

func main() {
	// array
	myArray := [...]int{67, 69, 66}
	// slice of aray
	mySlice := myArray[:]

	fmt.Println(myArray)
	fmt.Println(len(myArray))
	fmt.Println(mySlice)

	// make function
	myMakeSlice := make([]float32, 100)

	myMakeSlice[0] = 12
	myMakeSlice[1] = 123

	fmt.Println(myMakeSlice)

	// map
	myMap := make(map[int]string)
	fmt.Println(myMap)

	myMap[12] = "foo"
	myMap[67] = "bar"

	fmt.Println(myMap)
}
