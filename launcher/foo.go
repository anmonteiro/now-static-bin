package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x := os.Getenv("ANM")
	y, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("foo" + y)
		fmt.Println(err)
	}
	fmt.Printf("ANM: '%d'", y)
}