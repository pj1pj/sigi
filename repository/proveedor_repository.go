package repository

import (
	"errors"
	"fmt"
	"sync"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type ProveedorRepositoryMemoria struct {
	mu          sync.RWMutex
	proveedores []*models.Proveedor
}

func NuevoProveedorRepository() interfaces.ProveedorRepository {
	return &ProveedorRepositoryMemoria{
		proveedores: make([]*models.Proveedor, 0),
	}
}

func (r *ProveedorRepositoryMemoria) Agregar(proveedor *models.Proveedor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if proveedor == nil {
		return errors.New("el proveedor no puede ser nil")
	}

	for _, existente := range r.proveedores {
		if existente.Codigo() == proveedor.Codigo() {
			return fmt.Errorf("%w: ya existe un proveedor con ese código", utils.ErrRegistroDuplicado)
		}
	}

	r.proveedores = append(r.proveedores, proveedor)

	return nil
}

func (r *ProveedorRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Proveedor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, proveedor := range r.proveedores {
		if proveedor.Codigo() == codigo {
			return proveedor, nil
		}
	}

	return nil, fmt.Errorf("%w: proveedor no encontrado", utils.ErrRegistroNoEncontrado)
}

func (r *ProveedorRepositoryMemoria) ObtenerTodos() []*models.Proveedor {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]*models.Proveedor(nil), r.proveedores...)
}

func (r *ProveedorRepositoryMemoria) Actualizar(proveedor *models.Proveedor) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if proveedor == nil {
		return errors.New("el proveedor no puede ser nil")
	}

	for indice, existente := range r.proveedores {
		if existente.Codigo() == proveedor.Codigo() {
			r.proveedores[indice] = proveedor
			return nil
		}
	}

	return fmt.Errorf("%w: proveedor no encontrado", utils.ErrRegistroNoEncontrado)
}

func (r *ProveedorRepositoryMemoria) Eliminar(codigo string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for indice, proveedor := range r.proveedores {
		if proveedor.Codigo() == codigo {
			r.proveedores = append(
				r.proveedores[:indice],
				r.proveedores[indice+1:]...,
			)
			return nil
		}
	}

	return fmt.Errorf("%w: proveedor no encontrado", utils.ErrRegistroNoEncontrado)
}
