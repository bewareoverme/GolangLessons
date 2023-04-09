// findMinNumber --------------------------
package main

import "fmt"

func main() {
	fmt.Println(findMin(3, 6, 1, 7, 9, 3, -521, 521512))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}
