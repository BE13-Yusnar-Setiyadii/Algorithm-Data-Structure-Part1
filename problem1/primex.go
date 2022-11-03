package main

import (
	"fmt"
	"sort"
)

func PrimeX(number int) int {
	elements := []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	findIndex := sort.SearchInts(elements, number)
	if number == elements[findIndex] {
	}
	return elements[number]
}

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29
}
