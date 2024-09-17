package main

import "fmt"

func main() {

	// deklarasi variabel
	var nama_customer, nama_barang string
	var harga float64 = 25000
	var kuantitas int64
	var hasil_discount, total_harga float64

	// input nama customer
	fmt.Print("Masukkan nama_customer: ")
	fmt.Scanf("%s\n", &nama_customer)

	// input nama barang
	fmt.Print("Masukkan nama_barang: ")
	fmt.Scanf("%s\n", &nama_barang)

	// input kuantitas barang
	fmt.Print("Masukkan kuantitas_barang: ")
	fmt.Scanf("%d\n", &kuantitas_barang)

	// kondisi diskon, lebih dar 5 dapat 10%, selain itu 2%
	switch diskon {

	case "kuantitas>5":
		fmt.Println("Hasil Discount adalah : 0.1") // 10%

	default:
		fmt.Println("kuantitas tidak valid") // 2%

	}

	// proses
	sub_total := float64(kuantitas) * harga
	total_harga = sub_total - (hasil_discount * sub_total)

	// tampilkan hasil harga
	fmt.Println("Hasil Discount adalah : ", hasil_discount)
	fmt.Println("Kuantitas adalah : ", kuantitas)
	fmt.Println("Harga adalah : ", harga)
	fmt.Println("Total Harga adalah : ", total_harga)

}
