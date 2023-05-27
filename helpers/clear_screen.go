package helpers

import (
	"os"
	"os/exec"
)

func CleanScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
