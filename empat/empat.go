package main

import "fmt"

func main() {
	const printSegitiga = 6

	for i := printSegitiga; i >= 1; i-- {
		hasil := ""
		for j := 1; j <= i; j++ {
			hasil += fmt.Sprint(j)
		}
		fmt.Println(hasil)
	}
}