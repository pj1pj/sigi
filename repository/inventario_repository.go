package repository

import (
	"errors"

	"sigi/interfaces"
	"sigi/models"
)

type InventarioRepositoryMemoria struct {
	inventarios []*models.Inventario
}

func NuevoInventarioRepository() interfaces.InventarioRepository {
	return &InventarioRepositoryMemoria{
		inventarios: make([]*models.Inventario, 0),
	}
}

func (r *InventarioRepositoryMemoria) Agregar(inventario *models.Inventario) error {
	if inventario == nil {
		return errors.New("el registro de inventario no puede ser nil")
	}

	r.inventarios = append(r.inventarios, inventario)

	return nil
}

func (r *InventarioRepositoryMemoria) ObtenerTodos() []*models.Inventario {
	return r.inventarios
}

func (r *InventarioRepositoryMemoria) Actualizar(inventario *models.Inventario) error {
	if inventario == nil {
		return errors.New("el registro de inventario no puede ser nil")
	}

	for indice, existente := range r.inventarios {
		if mismoInventario(existente, inventario) {
			r.inventarios[indice] = inventario
			return nil
		}
	}

	return errors.New("registro de inventario no encontrado")
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
