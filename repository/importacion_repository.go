package repository

import (
	"errors"
	"fmt"
	"sync"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type ImportacionRepositoryMemoria struct {
	mu            sync.RWMutex
	importaciones []*models.Importacion
}

func NuevaImportacionRepository() interfaces.ImportacionRepository {
	return &ImportacionRepositoryMemoria{
		importaciones: make([]*models.Importacion, 0),
	}
}

func (r *ImportacionRepositoryMemoria) Agregar(importacion *models.Importacion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if importacion == nil {
		return errors.New("la importación no puede ser nil")
	}

	for _, existente := range r.importaciones {
		if existente.Codigo() == importacion.Codigo() {
			return fmt.Errorf("%w: ya existe una importación con ese código", utils.ErrRegistroDuplicado)
		}
	}

	r.importaciones = append(r.importaciones, importacion)

	return nil
}

func (r *ImportacionRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Importacion, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, importacion := range r.importaciones {
		if importacion.Codigo() == codigo {
			return importacion, nil
		}
	}

	return nil, fmt.Errorf("%w: importación no encontrada", utils.ErrRegistroNoEncontrado)
}

func (r *ImportacionRepositoryMemoria) ObtenerTodos() []*models.Importacion {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]*models.Importacion(nil), r.importaciones...)
}

func (r *ImportacionRepositoryMemoria) Actualizar(importacion *models.Importacion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if importacion == nil {
		return errors.New("la importación no puede ser nil")
	}

	for indice, existente := range r.importaciones {
		if existente.Codigo() == importacion.Codigo() {
			r.importaciones[indice] = importacion
			return nil
		}
	}

	return fmt.Errorf("%w: importación no encontrada", utils.ErrRegistroNoEncontrado)
}

func (r *ImportacionRepositoryMemoria) Eliminar(codigo string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for indice, importacion := range r.importaciones {
		if importacion.Codigo() == codigo {
			r.importaciones = append(
				r.importaciones[:indice],
				r.importaciones[indice+1:]...,
			)
			return nil
		}
	}

	return fmt.Errorf("%w: importación no encontrada", utils.ErrRegistroNoEncontrado)
}
