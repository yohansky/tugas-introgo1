package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var name string = "Yohanes Hubert"
	var age int = 23
	var height float64 = 170.1

	//2 Konversi
	// int > String
	str := strconv.Itoa(age)
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(str)
	// String > int
	num, err := strconv.Atoi(name)
	fmt.Println(reflect.TypeOf(num))
	fmt.Println(num)
	_ = err
	// float > String
	str1 := fmt.Sprintln(height)
	fmt.Println(reflect.TypeOf(str1))
	fmt.Println(str1)
	// String > float
	float, err := strconv.ParseFloat(name, 64)
	fmt.Println(reflect.TypeOf(float))
	fmt.Println(float)

}