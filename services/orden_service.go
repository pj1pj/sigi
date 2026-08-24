package services

import (
	"fmt"
	"strings"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
)

type OrdenService struct {
	repository interfaces.OrdenRepository
	generador  *utils.GeneradorCodigo
}

func NuevaOrdenService(
	repository interfaces.OrdenRepository,
	generador *utils.GeneradorCodigo,
) *OrdenService {
	return &OrdenService{
		repository: repository,
		generador:  generador,
	}
}

func (s *OrdenService) Crear(
	proveedor *models.Proveedor,
) (*models.OrdenCompra, error) {
	if proveedor == nil {
		return nil, fmt.Errorf("%w: el proveedor es obligatorio", utils.ErrDatoObligatorio)
	}

	if !proveedor.EstaActivo() {
		return nil, fmt.Errorf("%w: no se puede crear una orden con un proveedor inactivo", utils.ErrOperacionNoPermitida)
	}

	codigo := s.generador.Generar("ORD")

	orden := models.NuevaOrdenCompra(codigo, proveedor)

	if err := s.repository.Agregar(orden); err != nil {
		return nil, err
	}

	return orden, nil
}

func (s *OrdenService) Buscar(codigo string) (*models.OrdenCompra, error) {
	if strings.TrimSpace(codigo) == "" {
		return nil, fmt.Errorf("%w: el codigo de la orden es obligatorio", utils.ErrDatoObligatorio)
	}

	return s.repository.BuscarPorCodigo(codigo)
}

func (s *OrdenService) ObtenerTodos() []*models.OrdenCompra {
	return s.repository.ObtenerTodos()
}

func (s *OrdenService) AgregarProducto(
	codigoOrden string,
	producto *models.Producto,
) error {
	if producto == nil {
		return fmt.Errorf("%w: el producto es obligatorio", utils.ErrDatoObligatorio)
	}

	if !utils.CantidadValida(producto.Cantidad()) {
		return fmt.Errorf("%w: la cantidad del producto debe ser mayor que cero", utils.ErrDatoInvalido)
	}

	if !utils.PrecioValido(producto.PrecioUnitario()) {
		return fmt.Errorf("%w: el precio unitario debe ser mayor que cero", utils.ErrDatoInvalido)
	}

	orden, err := s.Buscar(codigoOrden)
	if err != nil {
		return err
	}

	if orden.EstaCancelada() {
		return fmt.Errorf("%w: no se pueden agregar productos a una orden cancelada", utils.ErrOperacionNoPermitida)
	}

	orden.AgregarProducto(producto)

	return s.repository.Actualizar(orden)
}

func (s *OrdenService) Confirmar(codigo string) error {
	orden, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	if len(orden.Productos()) == 0 {
		return fmt.Errorf("%w: no se puede confirmar una orden sin productos", utils.ErrOperacionNoPermitida)
	}

	if orden.EstaCancelada() {
		return fmt.Errorf("%w: no se puede confirmar una orden cancelada", utils.ErrOperacionNoPermitida)
	}

	orden.Confirmar()

	return s.repository.Actualizar(orden)
}

func (s *OrdenService) Cancelar(codigo string) error {
	orden, err := s.Buscar(codigo)
	if err != nil {
		return err
	}

	if orden.EstaConfirmada() {
		return fmt.Errorf("%w: no se puede cancelar una orden confirmada", utils.ErrOperacionNoPermitida)
	}

	orden.Cancelar()

	return s.repository.Actualizar(orden)
}
