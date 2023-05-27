package cli

import (
	"fmt"
	"os"
)

func ErrorHandler(msg string) {
	fmt.Println("terjadi kesalahan dalam aplikasi")
	fmt.Println("msg")

	var input string
	fmt.Println("tekan (m) untuk kembali ke menu utama")
	fmt.Println("tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "m":
		MainMenu()
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		ErrorHandler(msg)
	}
}
