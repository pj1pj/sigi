package repository

import (
	"errors"

	"sigi/interfaces"
	"sigi/models"
)

type TransporteRepositoryMemoria struct {
	transportes []*models.Transporte
}

func NuevoTransporteRepository() interfaces.TransporteRepository {
	return &TransporteRepositoryMemoria{
		transportes: make([]*models.Transporte, 0),
	}
}

func (r *TransporteRepositoryMemoria) Agregar(transporte *models.Transporte) error {
	if transporte == nil {
		return errors.New("el transporte no puede ser nil")
	}

	if _, err := r.BuscarPorCodigo(transporte.Codigo()); err == nil {
		return errors.New("ya existe un transporte con ese código")
	}

	r.transportes = append(r.transportes, transporte)

	return nil
}

func (r *TransporteRepositoryMemoria) BuscarPorCodigo(codigo string) (*models.Transporte, error) {
	for _, transporte := range r.transportes {
		if transporte.Codigo() == codigo {
			return transporte, nil
		}
	}

	return nil, errors.New("transporte no encontrado")
}

func (r *TransporteRepositoryMemoria) ObtenerTodos() []*models.Transporte {
	return r.transportes
}

func (r *TransporteRepositoryMemoria) Actualizar(transporte *models.Transporte) error {
	if transporte == nil {
		return errors.New("el transporte no puede ser nil")
	}

	for indice, existente := range r.transportes {
		if existente.Codigo() == transporte.Codigo() {
			r.transportes[indice] = transporte
			return nil
		}
	}

	return errors.New("transporte no encontrado")
}

func (r *TransporteRepositoryMemoria) Eliminar(codigo string) error {
	for indice, transporte := range r.transportes {
		if transporte.Codigo() == codigo {
			r.transportes = append(
				r.transportes[:indice],
				r.transportes[indice+1:]...,
			)
			return nil
		}
	}

	return errors.New("transporte no encontrado")
}
