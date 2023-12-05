package main

import (
	"fmt"
	"os"
	"unicode"

	"github.com/wfrodriguez/console/input"
	"github.com/wfrodriguez/console/output"
	"gitlab.com/tozd/go/errors"
)

var food = []string{
	"🌭 - Hot dog",
	"🌮 - Taco",
	"🌯 - Burrito",
	"🍔 - Hamburger",
	"🍕 - Pizza",
	"🍖 - Meat on bone",
	"🍗 - Poultry leg",
	"🍘 - Rice cracker",
}

func perror(err errors.E) {
	if err != nil {
		fmt.Printf("Ha ocurrido un error: % #+-.1v", err)
		os.Exit(1)
	}
}

// La contraseña debe tener las siguientes reglas:
//   - Al menos 7 letras
//   - al menos 1 número
//   - Al menos 1 mayúscula
//   - Al menos 1 carácter especial
func validPassword(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func main() {
	output.AddStyle("response", output.NewStyle(
		output.WithForeground(output.New256Color(14, output.Foreground)),
		output.WithItalic(),
	))
	output.AddStyle("ok", output.NewStyle(output.WithForeground(output.New256Color(10, output.Foreground))))
	output.AddStyle("error", output.NewStyle(output.WithForeground(output.New256Color(9, output.Foreground))))

	name, err := input.Ask(
		input.WithLabel("¿Cual es tu nombre?"),
		input.WithValidate(func(s string) error {
			if s == "" {
				return errors.New("El nombre no puede estar vacío")
			}
			return nil
		}),
	)
	perror(err)

	passwd, err := input.Ask(
		input.WithLabel("Tu contraseña:"),
		input.WithValidate(func(s string) error {
			if !validPassword(s) {
				return errors.New("La contraseña debe tener al menos 7 caracteres, 1 mayúscula, 1 minúscula, 1 número y 1 carácter especial")
			}
			return nil
		}),
		input.WithMask('?'),
	)
	perror(err)

	idx, sel, err := input.Select(
		input.WithSelectLabel("¿Cual es tu comida favorita?"),
		input.WithItems(food),
	)
	perror(err)

	desc, err := input.TextArea("Describe tu dia:", 50)
	perror(err)

	fmt.Println(output.Sprintf("<response>Hola</> %s!", name))
	fmt.Println(output.Sprintf("<response>Tu contraseña es</> %q", passwd))
	fmt.Println(output.Sprintf("<response>Tu comida favorita es</> %s", sel))
	fmt.Println(output.Sprintf("<response>Tu id de comida favorita es</> %d", idx))
	fmt.Println(output.Sprintf("<response>Tu descripción es</> %q", desc))
	if input.Confirm("¿Es correcto?", 3) {
		fmt.Println(output.Sprintf("<ok>Muchas gracias!!!</>"))
	} else {
		fmt.Println(output.Sprintf("<error>Adios!</>"))
	}

}
