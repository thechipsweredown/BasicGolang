package main

import "fmt"

func f(n int) {
	fmt.Println(n)
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
