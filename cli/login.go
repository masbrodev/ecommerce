package cli

import (
	"bufio"
	"ecommerce/config"
	"ecommerce/entity"
	"ecommerce/helpers"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func MenuLogin() {

	var input int

	fmt.Println("---Main Menu---")
	fmt.Println("1. Register User \n2. Login \nKetik 1 atau 2")

	fmt.Scanln(&input)
	switch input {
	case 1:
		Register()
	case 2:
		Login()
	}
}
func Login() {
	var username, password string
	helpers.CleanScreen()

	consoleReeader := bufio.NewReader(os.Stdin)
	fmt.Println("Login")
	fmt.Println("____________________")
	fmt.Println("Masukkan username: ")
	username, _ = consoleReeader.ReadString('\n')

	user := entity.User{}
	checkU := config.DB.Where("username = ? ", username).First(&user).Error
	if checkU != nil {
		helpers.CleanScreen()

		fmt.Println("Username tidak ada yang cocok")
		MenuLogin()
	}

	fmt.Println("Masukkan Password: ")
	password, _ = consoleReeader.ReadString('\n')

	checkP := config.DB.Where("username = ? ", username).Where("password = ?", password).First(&user).Error
	if checkP != nil {
		helpers.CleanScreen()

		fmt.Println("Password tidak cocok")
		MenuLogin()
	}

	// var ids entity.User
	// if err := config.DB.Last(&ids).Error; err != nil {
	// 	ErrorHandler(err.Error())
	// 	return
	// }

	randNum := rand.New(rand.NewSource(10))

	token := entity.Token{
		UserID:      int(user.ID),
		Name:        user.Username,
		TokenNumber: randNum.Int(),
		TokenDate:   time.Now(),
	}

	errT := config.DB.Create(&token).Error

	if errT != nil {
		ErrorHandler(errT.Error())
		return

	}
	MainMenu()

}

func Register() {
	var username, password string
	helpers.CleanScreen()
	consoleReeader := bufio.NewReader(os.Stdin)
	fmt.Println("Register")
	fmt.Println("____________________")
	fmt.Println("Masukkan username: ")
	username, _ = consoleReeader.ReadString('\n')

	fmt.Println("Masukkan Password: ")
	password, _ = consoleReeader.ReadString('\n')

	regist := entity.User{
		Username: username,
		Password: password,
	}
	sv := config.DB.Create(&regist).Error

	if sv != nil {
		ErrorHandler(sv.Error())
	}

	var ids entity.User
	if err := config.DB.Last(&ids).Error; err != nil {
		ErrorHandler(err.Error())
		return
	}

	randNum := rand.New(rand.NewSource(10))

	token := entity.Token{
		UserID:      int(ids.ID),
		Name:        ids.Username,
		TokenNumber: randNum.Int(),
		TokenDate:   time.Now(),
	}

	errT := config.DB.Create(&token).Error

	if errT != nil {
		ErrorHandler(errT.Error())
		return

	}
	MainMenu()

}

func ShowName() string {
	var user entity.Token
	if name := config.DB.Last(&user).Error; name != nil {
		ErrorHandler(name.Error())

	}
	return user.Name

}
