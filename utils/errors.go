package utils

import "errors"

var (
	ErrDatoObligatorio      = errors.New("el dato es obligatorio")
	ErrRegistroNoEncontrado = errors.New("registro no encontrado")
	ErrRegistroDuplicado    = errors.New("el registro ya existe")
	ErrOperacionNoPermitida = errors.New("la operación no está permitida")
	ErrDatoInvalido         = errors.New("el dato ingresado no es válido")
)
