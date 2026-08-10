package repository

import (
	"errors"

	"sigi/interfaces"
	"sigi/models"
)

type OrdenRepositoryMemoria struct {
	ordenes []*models.OrdenCompra
}

func NuevaOrdenRepository() interfaces.OrdenRepository {
	return &OrdenRepositoryMemoria{
		ordenes: make([]*models.OrdenCompra, 0),
	}
}

func (r *OrdenRepositoryMemoria) Agregar(orden *models.OrdenCompra) error {
	if orden == nil {
		return errors.New("la orden de compra no puede ser nil")
	}

	if _, err := r.BuscarPorCodigo(orden.Codigo()); err == nil {
		return errors.New("ya existe una orden con ese código")
	}

	r.ordenes = append(r.ordenes, orden)

	return nil
}

func (r *OrdenRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.OrdenCompra, error) {
	for _, orden := range r.ordenes {
		if orden.Codigo() == codigo {
			return orden, nil
		}
	}

	return nil, errors.New("orden de compra no encontrada")
}

func (r *OrdenRepositoryMemoria) ObtenerTodos() []*models.OrdenCompra {
	return r.ordenes
}

func (r *OrdenRepositoryMemoria) Actualizar(orden *models.OrdenCompra) error {
	if orden == nil {
		return errors.New("la orden de compra no puede ser nil")
	}

	for indice, existente := range r.ordenes {
		if existente.Codigo() == orden.Codigo() {
			r.ordenes[indice] = orden
			return nil
		}
	}

	return errors.New("orden de compra no encontrada")
}

func (r *OrdenRepositoryMemoria) Eliminar(codigo string) error {
	for indice, orden := range r.ordenes {
		if orden.Codigo() == codigo {
			r.ordenes = append(
				r.ordenes[:indice],
				r.ordenes[indice+1:]...,
			)
			return nil
		}
	}

	return errors.New("orden de compra no encontrada")
}
