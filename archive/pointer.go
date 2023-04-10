package main

import "fmt"

func main() {
	message := "viperviperr" // 0xc000048250
	fmt.Println(message)
	changeMsg(message)

	fmt.Println(&message) // 0xc000048270
}

func changeMsg(message string) {
	message += " (из функции changeMsg)"
}