package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lector = bufio.NewReader(os.Stdin)

func LeerTexto(mensaje string) string {
	fmt.Print(mensaje)

	texto, _ := lector.ReadString('\n')

	return strings.TrimSpace(texto)
}

func LeerEntero(mensaje string) (int, error) {
	texto := LeerTexto(mensaje)

	return strconv.Atoi(texto)
}

func LeerDecimal(mensaje string) (float64, error) {
	texto := LeerTexto(mensaje)

	return strconv.ParseFloat(texto, 64)
}

// Solicita una confirmación al usuario y repite la pregunta
// hasta recibir una respuesta válida de sí o no.
func Confirmar(mensaje string) bool {
	for {
		respuesta := strings.ToLower(LeerTexto(mensaje + " (s/n): "))

		switch respuesta {
		case "s", "si":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Respuesta no válida. Utilice s/n.")
		}
	}
}
