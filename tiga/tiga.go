package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func integerToString(integer int) string {
	return strconv.Itoa(integer)
}

func validateInteger(integer int) bool {
	// mtkStr := integerToString(80)
    // bahasaIndonesiaStr := integerToString(90)
    // bahasaInggrisStr := integerToString(89)
    // ipaStr := integerToString(69)

	// Definisikan pola regex untuk integer
	integerRegex := regexp.MustCompile(`^[0-9]+$`)
	_ = integerRegex

	// Cek apakah integer bernilai 0
	if integer == 0 {
		fmt.Println("Nilai harus diisi")
		return false
	}

	// Jika integer tidak bernilai 0, maka kembalikan nilai true
	return true
}

func hitungTotal(mtk,bahasaIndonesia,bahasaInggris,ipa int) int{//parameter/acuannya integer,hasil akhir function juga bernilai integer
	total := int(mtk + bahasaIndonesia + bahasaInggris + ipa) //buat var total dengan tipe data int
	fmt.Println("Total = ", total)
	return total
}

func hitungRata(total int) float64{ // mengambil nilai total (yang ada di main) lalu di konversi jadi float64
	rata := float64((total) / 4) 
	fmt.Println("Rata-rata = ", rata)
	return rata
}

func pilahGrade(rata float64) string{
	if rata >= 90 {
		return "Grade = A"
	} else if rata >= 80 && rata <= 89 {
		return"Grade = B"
	} else if rata >= 70 && rata <= 79 {
		return "Grade = C"
	} else if rata >= 60 && rata <= 69 {
		return "Grade = D"
	} else {
		return "Grade = E"
	}
}

func main() {
	mtkStr := integerToString(80)
    bahasaIndonesiaStr := integerToString(90)
    bahasaInggrisStr := integerToString(89)
    ipaStr := integerToString(69)

	if !validateInteger(80) {
        fmt.Println("mtk tidak valid")
    } else {
		fmt.Println("mtk:", mtkStr)
	}
    if !validateInteger(90) {
        fmt.Println("bahasa Indonesia tidak valid")
    } else {
		fmt.Println("bahasa Indonesia:", bahasaIndonesiaStr)
	}
    if !validateInteger(89) {
        fmt.Println("bahasa Inggris tidak valid")
    } else {
		fmt.Println("bahasa Inggris:", bahasaInggrisStr)
	}
    if !validateInteger(69) {
        fmt.Println("ipa tidak valid")
    } else {
		fmt.Println("ipa:", ipaStr)
	}    

	total := hitungTotal(80,90,89,69)// buat variabel total untuk menampung func total dengan parameter/acuannya angka
	rata := hitungRata(total) // buat variabel rata untuk menampung func rata dengan parameter/acuannya var total
	grade := pilahGrade(rata)// buat variabel grade untuk menampung func grade dengan parameter/acuannya var rata
	fmt.Println(grade)
	
}