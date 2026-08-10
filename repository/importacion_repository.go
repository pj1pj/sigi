package repository

import (
	"errors"

	"sigi/interfaces"
	"sigi/models"
)

type ImportacionRepositoryMemoria struct {
	importaciones []*models.Importacion
}

func NuevaImportacionRepository() interfaces.ImportacionRepository {
	return &ImportacionRepositoryMemoria{
		importaciones: make([]*models.Importacion, 0),
	}
}

func (r *ImportacionRepositoryMemoria) Agregar(importacion *models.Importacion) error {
	if importacion == nil {
		return errors.New("la importación no puede ser nil")
	}

	if _, err := r.BuscarPorCodigo(importacion.Codigo()); err == nil {
		return errors.New("ya existe una importación con ese código")
	}

	r.importaciones = append(r.importaciones, importacion)

	return nil
}

func (r *ImportacionRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Importacion, error) {
	for _, importacion := range r.importaciones {
		if importacion.Codigo() == codigo {
			return importacion, nil
		}
	}

	return nil, errors.New("importación no encontrada")
}

func (r *ImportacionRepositoryMemoria) ObtenerTodos() []*models.Importacion {
	return r.importaciones
}

func (r *ImportacionRepositoryMemoria) Actualizar(importacion *models.Importacion) error {
	if importacion == nil {
		return errors.New("la importación no puede ser nil")
	}

	for indice, existente := range r.importaciones {
		if existente.Codigo() == importacion.Codigo() {
			r.importaciones[indice] = importacion
			return nil
		}
	}

	return errors.New("importación no encontrada")
}

func (r *ImportacionRepositoryMemoria) Eliminar(codigo string) error {
	for indice, importacion := range r.importaciones {
		if importacion.Codigo() == codigo {
			r.importaciones = append(
				r.importaciones[:indice],
				r.importaciones[indice+1:]...,
			)
			return nil
		}
	}

	return errors.New("importación no encontrada")
}
