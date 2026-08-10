package utils

import (
	"regexp"
	"strings"
)

func TextoValido(valor string) bool {
	return strings.TrimSpace(valor) != ""
}

func CorreoValido(correo string) bool {
	correo = strings.TrimSpace(correo)

	if correo == "" {
		return false
	}

	patron := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	return regexp.MustCompile(patron).MatchString(correo)
}

func TelefonoValido(telefono string) bool {
	telefono = strings.TrimSpace(telefono)

	if telefono == "" {
		return false
	}

	patron := `^[0-9+\-() ]{7,20}$`

	return regexp.MustCompile(patron).MatchString(telefono)
}

func CantidadValida(cantidad int) bool {
	return cantidad > 0
}

func PrecioValido(precio float64) bool {
	return precio > 0
}
