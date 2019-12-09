package main

import (
	"fmt"
)

func main() {
	n, e := fmt.Println("hello", 42, "wack")
	result := fmt.Sprint(n, e)
	fmt.Println(result)
}
