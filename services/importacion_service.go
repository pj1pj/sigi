package services

import (
	"errors"
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
		return nil, errors.New("la orden de compra es obligatoria")
	}

	if transporte == nil {
		return nil, errors.New("el transporte es obligatorio")
	}

	if !orden.EstaConfirmada() {
		return nil, errors.New("la orden debe estar confirmada para crear una importacion")
	}

	if !transporte.EstaActivo() {
		return nil, errors.New("no se puede utilizar un transporte inactivo")
	}

	if strings.TrimSpace(ciudadOrigen) == "" {
		return nil, errors.New("la ciudad de origen es obligatoria")
	}

	if strings.TrimSpace(ciudadDestino) == "" {
		return nil, errors.New("la ciudad de destino es obligatoria")
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
		return nil, errors.New("el codigo de la importacion es obligatorio")
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
		return errors.New("el estado de importacion no es valido")
	}

	if !transicionPermitida(importacion.Estado(), nuevoEstado) {
		return errors.New("la transicion de estado no esta permitida")
	}

	if nuevoEstado == models.ImportacionLlegadaBodega &&
		strings.TrimSpace(ubicacion) == "" {
		return errors.New("la ubicacion de bodega es obligatoria")
	}

	importacion.ActualizarEstado(nuevoEstado)

	if err := s.repository.Actualizar(importacion); err != nil {
		return err
	}

	if nuevoEstado == models.ImportacionLlegadaBodega {
		if s.inventarioService == nil {
			return errors.New("el servicio de inventario no esta configurado")
		}

		if err := s.inventarioService.ProcesarLlegada(importacion, ubicacion); err != nil {
			return err
		}
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
