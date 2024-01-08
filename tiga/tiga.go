package main

import "fmt"

func hitungTotal(mtk,bahasaIndonesia,bahasaInggris,ipa int) int{//parameter/acuannya integer,hasil akhir function juga bernilai integer
	if mtk <= 0 || bahasaIndonesia <= 0 || bahasaInggris <= 0 || ipa <= 0 {
		fmt.Println("Nilai harus lebih dari 0.")
	}
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

	total := hitungTotal(80,90,89,69)// buat variabel total untuk menampung func total dengan parameter/acuannya angka
	rata := hitungRata(total) // buat variabel rata untuk menampung func rata dengan parameter/acuannya var total
	grade := pilahGrade(rata)// buat variabel grade untuk menampung func grade dengan parameter/acuannya var rata
	fmt.Println(grade)
}