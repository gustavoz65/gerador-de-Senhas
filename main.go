package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

// generatePassword generates a random password of the given length.
func generatePassword(length int) string {
	var password strings.Builder
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ""
		}
		password.WriteByte(charset[index.Int64()])
	}
	return password.String()
}

func main() {
	// Create a new Fyne app
	myApp := app.New()
	myWindow := myApp.NewWindow("Gerador de Senhas")

	// Input field for password length
	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("Digite o comprimento da senha (ex: 12)")

	// Output field for the generated password
	passwordOutput := widget.NewEntry()
	passwordOutput.SetPlaceHolder("Senha gerada aparecerá aqui")
	passwordOutput.Disable()

	// Button to generate password
	generateButton := widget.NewButton("Gerar Senha", func() {
		length := 12 // Default length
		if lenText := lengthEntry.Text; len(lenText) > 0 {
			fmt.Sscanf(lenText, "%d", &length)
			if length <= 0 {
				length = 12
			}
		}
		password := generatePassword(length)
		passwordOutput.SetText(password)
	})

	// Layout the widgets
	content := container.NewVBox(
		widget.NewLabel("Gerador de Senhas"),
		lengthEntry,
		generateButton,
		passwordOutput,
	)

	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

}
