package services

import (
	"errors"
	"strings"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type TransporteService struct {
	repository interfaces.TransporteRepository
	generador  *utils.GeneradorCodigo
}

func NuevoTransporteService(
	repository interfaces.TransporteRepository,
	generador *utils.GeneradorCodigo,
) *TransporteService {
	return &TransporteService{
		repository: repository,
		generador:  generador,
	}
}

func (s *TransporteService) Registrar(
	tipo models.TipoTransporte,
	empresa string,
	pais string,
	contacto string,
	telefono string,
	correo string,
) (*models.Transporte, error) {
	if !tipoValido(tipo) {
		return nil, errors.New("el tipo de transporte no es valido")
	}

	if strings.TrimSpace(empresa) == "" {
		return nil, errors.New("la empresa de transporte es obligatoria")
	}

	if strings.TrimSpace(pais) == "" {
		return nil, errors.New("el pais del transporte es obligatorio")
	}

	if strings.TrimSpace(contacto) == "" {
		return nil, errors.New("el contacto del transporte es obligatorio")
	}

	if strings.TrimSpace(telefono) == "" {
		return nil, errors.New("el telefono del transporte es obligatorio")
	}

	codigo := s.generador.Generar("TRN")

	transporte := models.NuevoTransporte(
		codigo,
		tipo,
		empresa,
		pais,
		contacto,
		telefono,
		correo,
	)

	if err := s.repository.Agregar(transporte); err != nil {
		return nil, err
	}

	return transporte, nil
}

func (s *TransporteService) Buscar(codigo string) (*models.Transporte, error) {
	if strings.TrimSpace(codigo) == "" {
		return nil, errors.New("el codigo del transporte es obligatorio")
	}

	return s.repository.BuscarPorCodigo(codigo)
}

func (s *TransporteService) ObtenerTodos() []*models.Transporte {
	return s.repository.ObtenerTodos()
}

func (s *TransporteService) Activar(codigo string) error {
	transporte, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	transporte.Activar()

	return s.repository.Actualizar(transporte)
}

func (s *TransporteService) Desactivar(codigo string) error {
	transporte, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	transporte.Desactivar()

	return s.repository.Actualizar(transporte)
}

func (s *TransporteService) ActualizarContacto(
	codigo string,
	contacto string,
	telefono string,
	correo string,
) error {
	if strings.TrimSpace(contacto) == "" {
		return errors.New("el contacto es obligatorio")
	}

	if strings.TrimSpace(telefono) == "" {
		return errors.New("el telefono es obligatorio")
	}

	transporte, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	transporte.ActualizarContacto(contacto, telefono, correo)

	return s.repository.Actualizar(transporte)
}

func (s *TransporteService) Eliminar(codigo string) error {
	transporte, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	if transporte.EstaActivo() {
		return errors.New("no se puede eliminar un transporte activo")
	}

	return s.repository.Eliminar(codigo)
}

func tipoValido(tipo models.TipoTransporte) bool {
	return tipo == models.TransporteMaritimo ||
		tipo == models.TransporteAereo ||
		tipo == models.TransporteTerrestre
}
