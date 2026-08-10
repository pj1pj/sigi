package utils

import "fmt"

type GeneradorCodigo struct {
	contadores map[string]int
}

func NuevoGeneradorCodigo() *GeneradorCodigo {
	return &GeneradorCodigo{
		contadores: make(map[string]int),
	}
}

// Cada prefijo mantiene su propio contador para generar códigos
// consecutivos e independientes para cada tipo de entidad.
func (g *GeneradorCodigo) Generar(prefijo string) string {
	g.contadores[prefijo]++

	return fmt.Sprintf("%s-%04d", prefijo, g.contadores[prefijo])
}
