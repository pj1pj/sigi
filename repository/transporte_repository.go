package repository

import (
	"errors"
	"fmt"
	"sync"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type TransporteRepositoryMemoria struct {
	mu          sync.RWMutex
	transportes []*models.Transporte
}

func NuevoTransporteRepository() interfaces.TransporteRepository {
	return &TransporteRepositoryMemoria{
		transportes: make([]*models.Transporte, 0),
	}
}

func (r *TransporteRepositoryMemoria) Agregar(transporte *models.Transporte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if transporte == nil {
		return errors.New("el transporte no puede ser nil")
	}

	for _, existente := range r.transportes {
		if existente.Codigo() == transporte.Codigo() {
			return fmt.Errorf("%w: ya existe un transporte con ese código", utils.ErrRegistroDuplicado)
		}
	}

	r.transportes = append(r.transportes, transporte)

	return nil
}

func (r *TransporteRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Transporte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, transporte := range r.transportes {
		if transporte.Codigo() == codigo {
			return transporte, nil
		}
	}

	return nil, fmt.Errorf("%w: transporte no encontrado", utils.ErrRegistroNoEncontrado)
}

func (r *TransporteRepositoryMemoria) ObtenerTodos() []*models.Transporte {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]*models.Transporte(nil), r.transportes...)
}

func (r *TransporteRepositoryMemoria) Actualizar(transporte *models.Transporte) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if transporte == nil {
		return errors.New("el transporte no puede ser nil")
	}

	for indice, existente := range r.transportes {
		if existente.Codigo() == transporte.Codigo() {
			r.transportes[indice] = transporte
			return nil
		}
	}

	return fmt.Errorf("%w: transporte no encontrado", utils.ErrRegistroNoEncontrado)
}

func (r *TransporteRepositoryMemoria) Eliminar(codigo string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for indice, transporte := range r.transportes {
		if transporte.Codigo() == codigo {
			r.transportes = append(
				r.transportes[:indice],
				r.transportes[indice+1:]...,
			)
			return nil
		}
	}

	return fmt.Errorf("%w: transporte no encontrado", utils.ErrRegistroNoEncontrado)
}
