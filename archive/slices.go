package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := [5]string{"1", "2", "3","4","5"}

	messages := make([]string, 10000)
	messages = append(messages, "10001") 

	
	printMessages(messages)
	fmt.Println(len(messages))
	fmt.Println(cap(messages))
}

func printMessages(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}
	messages[1] = "5"

	fmt.Println(messages)
	

	return nil
}