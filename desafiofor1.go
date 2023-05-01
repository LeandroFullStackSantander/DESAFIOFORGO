package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Este número é divisível por 3:", i)
		}
	}
}
