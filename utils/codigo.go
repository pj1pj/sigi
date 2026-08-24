package utils

import (
	"fmt"
	"sync"
)

type GeneradorCodigo struct {
	mu         sync.Mutex
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
	g.mu.Lock()
	defer g.mu.Unlock()

	g.contadores[prefijo]++

	return fmt.Sprintf("%s-%04d", prefijo, g.contadores[prefijo])
}
