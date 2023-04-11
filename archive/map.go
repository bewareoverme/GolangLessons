package main

import "fmt"

func main() {
	users := map[string]int{
		"Vasya": 15,
		"Petya": 23,
		"Kostya": 48,
	}

	age, exist := users["serega"]
	// fmt.Println(age, exist)
	if exist {
		fmt.Println("Kostya", age)
	} else {
		fmt.Println("нет в списке")
	}
	
	// for key, value := range users {
	// 	fmt.Println(key, value)
	// }
}
