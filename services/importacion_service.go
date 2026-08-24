package services

import (
	"fmt"
	"strings"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type ImportacionService struct {
	repository        interfaces.ImportacionRepository
	inventarioService *InventarioService
	generador         *utils.GeneradorCodigo
}

func NuevaImportacionService(
	repository interfaces.ImportacionRepository,
	inventarioService *InventarioService,
	generador *utils.GeneradorCodigo,
) *ImportacionService {
	return &ImportacionService{
		repository:        repository,
		inventarioService: inventarioService,
		generador:         generador,
	}
}

func (s *ImportacionService) Registrar(
	orden *models.OrdenCompra,
	transporte *models.Transporte,
	ciudadOrigen string,
	ciudadDestino string,
) (*models.Importacion, error) {
	if orden == nil {
		return nil, fmt.Errorf("%w: la orden de compra es obligatoria", utils.ErrDatoObligatorio)
	}

	if transporte == nil {
		return nil, fmt.Errorf("%w: el transporte es obligatorio", utils.ErrDatoObligatorio)
	}

	if !orden.EstaConfirmada() {
		return nil, fmt.Errorf("%w: la orden debe estar confirmada para crear una importacion", utils.ErrOperacionNoPermitida)
	}

	if !transporte.EstaActivo() {
		return nil, fmt.Errorf("%w: no se puede utilizar un transporte inactivo", utils.ErrOperacionNoPermitida)
	}

	if !utils.TextoValido(ciudadOrigen) {
		return nil, fmt.Errorf("%w: la ciudad de origen es obligatoria", utils.ErrDatoObligatorio)
	}

	if !utils.TextoValido(ciudadDestino) {
		return nil, fmt.Errorf("%w: la ciudad de destino es obligatoria", utils.ErrDatoObligatorio)
	}

	codigo := s.generador.Generar("IMP")

	importacion := models.NuevaImportacion(
		codigo,
		orden,
		transporte,
		ciudadOrigen,
		ciudadDestino,
	)

	if err := s.repository.Agregar(importacion); err != nil {
		return nil, err
	}

	return importacion, nil
}

func (s *ImportacionService) Buscar(codigo string) (*models.Importacion, error) {
	if strings.TrimSpace(codigo) == "" {
		return nil, fmt.Errorf("%w: el codigo de la importacion es obligatorio", utils.ErrDatoObligatorio)
	}

	return s.repository.BuscarPorCodigo(codigo)
}

func (s *ImportacionService) ObtenerTodos() []*models.Importacion {
	return s.repository.ObtenerTodos()
}

func (s *ImportacionService) ActualizarTracking(
	codigo string,
	nuevoEstado models.EstadoImportacion,
	ubicacion string,
) error {
	importacion, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	if !estadoValido(nuevoEstado) {
		return fmt.Errorf("%w: el estado de importacion no es válido", utils.ErrDatoInvalido)
	}

	if !transicionPermitida(importacion.Estado(), nuevoEstado) {
		return fmt.Errorf("%w: la transicion de estado no esta permitida", utils.ErrOperacionNoPermitida)
	}

	if nuevoEstado == models.ImportacionLlegadaBodega &&
		!utils.TextoValido(ubicacion) {
		return fmt.Errorf("%w: la ubicacion de bodega es obligatoria", utils.ErrDatoObligatorio)
	}

	estadoAnterior := importacion.Estado()
	importacion.ActualizarEstado(nuevoEstado)

	if nuevoEstado == models.ImportacionLlegadaBodega {
		if s.inventarioService == nil {
			importacion.ActualizarEstado(estadoAnterior)
			return fmt.Errorf("%w: el servicio de inventario no esta configurado", utils.ErrOperacionNoPermitida)
		}

		if err := s.inventarioService.ProcesarLlegada(importacion, ubicacion); err != nil {
			importacion.ActualizarEstado(estadoAnterior)
			return err
		}
	}

	if err := s.repository.Actualizar(importacion); err != nil {
		importacion.ActualizarEstado(estadoAnterior)
		return err
	}

	return nil
}

func estadoValido(estado models.EstadoImportacion) bool {
	return estado == models.ImportacionEnPreparacion ||
		estado == models.ImportacionEnTransito ||
		estado == models.ImportacionEnAduana ||
		estado == models.ImportacionLlegadaBodega
}

func transicionPermitida(
	estadoActual models.EstadoImportacion,
	nuevoEstado models.EstadoImportacion,
) bool {
	switch estadoActual {
	case models.ImportacionEnPreparacion:
		return nuevoEstado == models.ImportacionEnTransito

	case models.ImportacionEnTransito:
		return nuevoEstado == models.ImportacionEnAduana

	case models.ImportacionEnAduana:
		return nuevoEstado == models.ImportacionLlegadaBodega

	case models.ImportacionLlegadaBodega:
		return false

	default:
		return false
	}
}
