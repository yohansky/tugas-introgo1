package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1 Println
	var name string = "Yohanes Hubert"
	var age int = 23
	var height float64 = 170.1
	var isMarried bool = false
	var interestInCoding bool = false

	fmt.Println(reflect.TypeOf(name))             //tipe data name
	fmt.Println(name)                             //nilai name
	fmt.Println(reflect.TypeOf(age))              // tipe data age
	fmt.Println(age)                              // nilai age
	fmt.Println(reflect.TypeOf(height))           // tipe data height
	fmt.Println(height)                           // nilai height
	fmt.Println(reflect.TypeOf(isMarried))        // tipe data isMarried
	fmt.Println(isMarried)                        // nilai isMarried
	fmt.Println(reflect.TypeOf(interestInCoding)) // tipe data interestInCoding
	fmt.Println(interestInCoding)                 // nilai
	// _, _, _, _ =age, height,isMarried,interestInCoding
	fmt.Println("========================================")
}