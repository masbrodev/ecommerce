package cli

import (
	"bufio"
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/helpers"
	"fmt"
	"os"
	"time"
)

func ListProduct() {
	helpers.CleanScreen()

	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("--- List produk ---")
	for _, product := range products {
		product.PrintDetail()
	}
	var input string
	fmt.Println("Masukkan Id product untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	fmt.Scanln(&input)
	switch input {
	case "m":
		MainMenu()
	case "q":
		fmt.Println("Terima kasih telah menggunakaan apliaksi Mini Ecommerce")
		os.Exit(1)
	default:
		OrderProduct(input)

	}
}
func OrderProduct(id string) {
	helpers.CleanScreen()

	var product entity.Product
	err := config.DB.Where("ID = ?", id).First(&product).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	product.PrintDetail()

	var input string
	fmt.Println("Tekan (y) untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	fmt.Scanln(&input)

	switch input {
	case "y":
		CreateOrder(product)
	case "m":
		MainMenu()
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		OrderProduct(id)
	}

}

func CreateOrder(product entity.Product) {
	helpers.CleanScreen()
	consoleReeader := bufio.NewReader(os.Stdin)
	var email string
	var address string

	fmt.Println("untuk melakukan order silahkan melengkapi data berikut")
	fmt.Println("Masukkan email: ")
	email, _ = consoleReeader.ReadString('\n')

	fmt.Println("Masukkan alamat: ")
	address, _ = consoleReeader.ReadString('\n')

	order := entity.Order{
		ProductID:    int(product.ID),
		BuyerEmail:   email,
		BuyerAddress: address,
		OrderDate:    time.Now(),
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Order berhasil")
	var input string
	fmt.Println("Tekan (key apapun) untuk kembali ke halaman utama")
	fmt.Println("tekan (q) untuk keluar dari aplikasi")

	_, err = fmt.Scanln(&input)
	if err != nil {
		MainMenu()
	}
	switch input {
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan Mini Ecommerce")
		os.Exit(1)
	default:
		MainMenu()
	}
}
