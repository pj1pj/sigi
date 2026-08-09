package repository

import (
	"errors"

	"sigi/interfaces"
	"sigi/models"
)

type ProveedorRepositoryMemoria struct {
	proveedores []*models.Proveedor
}

func NuevoProveedorRepository() interfaces.ProveedorRepository {
	return &ProveedorRepositoryMemoria{
		proveedores: make([]*models.Proveedor, 0),
	}
}

func (r *ProveedorRepositoryMemoria) Agregar(proveedor *models.Proveedor) error {
	if proveedor == nil {
		return errors.New("el proveedor no puede ser nil")
	}

	if _, err := r.BuscarPorCodigo(proveedor.Codigo()); err == nil {
		return errors.New("ya existe un proveedor con ese código")
	}

	r.proveedores = append(r.proveedores, proveedor)

	return nil
}

func (r *ProveedorRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Proveedor, error) {
	for _, proveedor := range r.proveedores {
		if proveedor.Codigo() == codigo {
			return proveedor, nil
		}
	}

	return nil, errors.New("proveedor no encontrado")
}

func (r *ProveedorRepositoryMemoria) ObtenerTodos() []*models.Proveedor {
	return r.proveedores
}

func (r *ProveedorRepositoryMemoria) Actualizar(proveedor *models.Proveedor) error {
	if proveedor == nil {
		return errors.New("el proveedor no puede ser nil")
	}

	for indice, existente := range r.proveedores {
		if existente.Codigo() == proveedor.Codigo() {
			r.proveedores[indice] = proveedor
			return nil
		}
	}

	return errors.New("proveedor no encontrado")
}

func (r *ProveedorRepositoryMemoria) Eliminar(codigo string) error {
	for indice, proveedor := range r.proveedores {
		if proveedor.Codigo() == codigo {
			r.proveedores = append(
				r.proveedores[:indice],
				r.proveedores[indice+1:]...,
			)
			return nil
		}
	}

	return errors.New("proveedor no encontrado")
}
