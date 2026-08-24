package repository

import (
	"errors"
	"fmt"
	"sync"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type InventarioRepositoryMemoria struct {
	mu          sync.RWMutex
	inventarios []*models.Inventario
}

func NuevoInventarioRepository() interfaces.InventarioRepository {
	return &InventarioRepositoryMemoria{
		inventarios: make([]*models.Inventario, 0),
	}
}

func (r *InventarioRepositoryMemoria) Agregar(inventario *models.Inventario) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if inventario == nil {
		return errors.New("el registro de inventario no puede ser nil")
	}

	r.inventarios = append(r.inventarios, inventario)

	return nil
}

func (r *InventarioRepositoryMemoria) ObtenerTodos() []*models.Inventario {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]*models.Inventario(nil), r.inventarios...)
}

func (r *InventarioRepositoryMemoria) Actualizar(inventario *models.Inventario) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if inventario == nil {
		return errors.New("el registro de inventario no puede ser nil")
	}

	for indice, existente := range r.inventarios {
		if mismoInventario(existente, inventario) {
			r.inventarios[indice] = inventario
			return nil
		}
	}

	return fmt.Errorf("%w: registro de inventario no encontrado", utils.ErrRegistroNoEncontrado)
}

// La comparación utiliza las relaciones del registro en lugar de agregar
// un identificador artificial, porque el inventario se genera a partir
// de productos asociados a una importación y una orden de compra.
func mismoInventario(a, b *models.Inventario) bool {
	if a == nil || b == nil {
		return false
	}

	if a.Producto() == nil || b.Producto() == nil {
		return false
	}

	if a.Importacion() == nil || b.Importacion() == nil {
		return false
	}

	return a.Producto().Nombre() == b.Producto().Nombre() &&
		a.Importacion().Codigo() == b.Importacion().Codigo()
}
