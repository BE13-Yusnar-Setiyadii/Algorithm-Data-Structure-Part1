package main

import (
	"fmt"
	"math"
)

func checkPrima(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for baris := 2; baris <= sqrtNumber; baris++ {
		if number%baris == 0 {
			return false
		}
	}
	return true
}

func PrimaSegiEmpat(wide, high, start int) string {
	var result string
	var number int = start
	var sumNumber int
	for baris := 1; baris <= high; baris++ {
		for kolom := 1; kolom <= wide; kolom++ {
			number++
			for !checkPrima(number) {
				number++
			}
			result = result + fmt.Sprintf("%d", number) + "\t"
			sumNumber += number
		}
		result += "\n"
	}
	result += fmt.Sprintf("%d", sumNumber)
	return result 
}

func main() {

	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// /*
	// 	17	19
	// 	23	29
	// 	31	37
	// 	156
	// */
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// /*
	// 	2	3	5	7	11
	// 	13	17	19	23	29
	// 	129
	// */
}
