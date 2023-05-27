package cli

import (
	"bufio"
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/helpers"
	"fmt"
	"os"
)

func ListOrder() {
	helpers.CleanScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var order []entity.Order

	err := config.DB.Preload("Product").Find(&order).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("---List oder---")
	for _, order := range order {
		order.PrintDetail()
	}

	var input string
	fmt.Println("Tekan (key apapun) untuk kembali ke halaman utama")
	fmt.Println("tekan (q) untuk keluar dari aplikasi")

	input, _ = consoleReader.ReadString('\n')

	switch input {
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan Mini Ecommerce")
		os.Exit(1)
	default:
		MainMenu()
	}
}
