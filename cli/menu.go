package cli

import (
	"ecommerce/helpers"
	"fmt"
	"os"
)

func MainMenu() {
	helpers.CleanScreen()

	fmt.Println("Selamat datang di Mini Ecommerce App")
	fmt.Println("____________________________________")

	var input string
	fmt.Println("Tekan (1) untuk melanjutkan ke list product")
	fmt.Println("Tekan (2) untuk melanjutkan ke list order")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "1":
		ListProduct()
	case "2":
		ListOrder()
	case "3", "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		MainMenu()
	}

}
