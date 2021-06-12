package main

import "fmt"

func Cat() string {
	//return "Wang...."
	word := "Miao"
	return fmt.Sprintf("%s~~~~", word)
}

func main() {
	saying := Cat()
	fmt.Print(saying)
}
