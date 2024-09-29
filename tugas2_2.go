package main

import "fmt"

// Struct untuk merepresentasikan sebuah buku
type Book struct {
	Title  string
	Writer string
	Price  int
	Count  int
}

// Struct untuk merepresentasikan transaksi
type Transaction struct {
	Total int
	Date  string
	Books []Book
}

func main() {
	// Membuat slice dari Book yang berisi buku-buku yang dibeli
	books := []Book{
		{
			Title:  "Laut Bercerita",
			Writer: "Leila S. Chudori",
			Price:  108000,
			Count:  1,
		},
		{
			Title:  "Bumi",
			Writer: "Tere Liye",
			Price:  97000,
			Count:  1,
		},
	}

	// Menghitung total harga transaksi
	total := 0
	for _, book := range books {
		total += book.Price * book.Count
	}

	// Membuat transaksi dengan data yang sudah tersedia
	tx1 := Transaction{
		Total: total,
		Date:  "23-01-2021",
		Books: books,
	}

	// Menampilkan transaksi yang terjadi
	fmt.Println("Detail Transaksi:")
	fmt.Println("Tanggal:", tx1.Date)
	fmt.Println("Total:", tx1.Total)
	fmt.Println("\nDaftar Buku yang Dibeli:")
	for _, book := range tx1.Books {
		fmt.Printf("Judul: %s\nPenulis: %s\nHarga: %d\nJumlah: %d\n\n", book.Title, book.Writer, book.Price, book.Count)
	}
}
