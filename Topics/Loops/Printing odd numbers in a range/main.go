package main

import "fmt"

func main() {
	for i := 3; i < 1024; i += 2 {
		fmt.Println(i)
	}
}
