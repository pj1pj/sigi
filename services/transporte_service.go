package services

import (
	"fmt"
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
		return nil, fmt.Errorf("%w: el tipo de transporte no es válido", utils.ErrDatoInvalido)
	}

	if !utils.TextoValido(empresa) {
		return nil, fmt.Errorf("%w: la empresa de transporte es obligatoria", utils.ErrDatoObligatorio)
	}

	if !utils.TextoValido(pais) {
		return nil, fmt.Errorf("%w: el pais del transporte es obligatorio", utils.ErrDatoObligatorio)
	}

	if !utils.TextoValido(contacto) {
		return nil, fmt.Errorf("%w: el contacto del transporte es obligatorio", utils.ErrDatoObligatorio)
	}

	if !utils.TelefonoValido(telefono) {
		return nil, fmt.Errorf("%w: el telefono del transporte no es válido", utils.ErrDatoInvalido)
	}

	if strings.TrimSpace(correo) != "" && !utils.CorreoValido(correo) {
		return nil, fmt.Errorf("%w: el correo del transporte no es válido", utils.ErrDatoInvalido)
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
		return nil, fmt.Errorf("%w: el codigo del transporte es obligatorio", utils.ErrDatoObligatorio)
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
		return fmt.Errorf("%w: el contacto es obligatorio", utils.ErrDatoObligatorio)
	}

	if strings.TrimSpace(telefono) == "" {
		return fmt.Errorf("%w: el telefono no es válido", utils.ErrDatoInvalido)
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
		return fmt.Errorf("%w: no se puede eliminar un transporte activo", utils.ErrOperacionNoPermitida)
	}

	return s.repository.Eliminar(codigo)
}

func tipoValido(tipo models.TipoTransporte) bool {
	return tipo == models.TransporteMaritimo ||
		tipo == models.TransporteAereo ||
		tipo == models.TransporteTerrestre
}
