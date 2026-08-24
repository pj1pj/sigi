package repository

import (
	"errors"
	"fmt"
	"sync"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type OrdenRepositoryMemoria struct {
	mu      sync.RWMutex
	ordenes []*models.OrdenCompra
}

func NuevaOrdenRepository() interfaces.OrdenRepository {
	return &OrdenRepositoryMemoria{
		ordenes: make([]*models.OrdenCompra, 0),
	}
}

func (r *OrdenRepositoryMemoria) Agregar(orden *models.OrdenCompra) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if orden == nil {
		return errors.New("la orden de compra no puede ser nil")
	}

	for _, existente := range r.ordenes {
		if existente.Codigo() == orden.Codigo() {
			return fmt.Errorf("%w: ya existe una orden con ese código", utils.ErrRegistroDuplicado)
		}
	}

	r.ordenes = append(r.ordenes, orden)

	return nil
}

func (r *OrdenRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.OrdenCompra, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, orden := range r.ordenes {
		if orden.Codigo() == codigo {
			return orden, nil
		}
	}

	return nil, fmt.Errorf("%w: orden de compra no encontrada", utils.ErrRegistroNoEncontrado)
}

func (r *OrdenRepositoryMemoria) ObtenerTodos() []*models.OrdenCompra {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]*models.OrdenCompra(nil), r.ordenes...)
}

func (r *OrdenRepositoryMemoria) Actualizar(orden *models.OrdenCompra) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if orden == nil {
		return errors.New("la orden de compra no puede ser nil")
	}

	for indice, existente := range r.ordenes {
		if existente.Codigo() == orden.Codigo() {
			r.ordenes[indice] = orden
			return nil
		}
	}

	return fmt.Errorf("%w: orden de compra no encontrada", utils.ErrRegistroNoEncontrado)
}

func (r *OrdenRepositoryMemoria) Eliminar(codigo string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for indice, orden := range r.ordenes {
		if orden.Codigo() == codigo {
			r.ordenes = append(
				r.ordenes[:indice],
				r.ordenes[indice+1:]...,
			)
			return nil
		}
	}

	return fmt.Errorf("%w: orden de compra no encontrada", utils.ErrRegistroNoEncontrado)
}
