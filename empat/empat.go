package main

import "fmt"

func perulangan(printSegitiga int) {
	for i := printSegitiga; i >= 1; i-- {
		hasil := ""
		for j := 1; j <= i; j++ {
			hasil += fmt.Sprint(" ",j)
		}
		fmt.Println(hasil)
	}
}

func main() {
	perulangan(6)
}