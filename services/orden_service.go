package services

import (
	"errors"
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
		return nil, errors.New("el proveedor es obligatorio")
	}

	if !proveedor.EstaActivo() {
		return nil, errors.New("no se puede crear una orden con un proveedor inactivo")
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
		return nil, errors.New("el codigo de la orden es obligatorio")
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
		return errors.New("el producto es obligatorio")
	}

	if producto.Cantidad() <= 0 {
		return errors.New("la cantidad del producto debe ser mayor que cero")
	}

	if producto.PrecioUnitario() <= 0 {
		return errors.New("el precio unitario debe ser mayor que cero")
	}

	orden, err := s.Buscar(codigoOrden)
	if err != nil {
		return err
	}

	if orden.EstaCancelada() {
		return errors.New("no se pueden agregar productos a una orden cancelada")
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
		return errors.New("no se puede confirmar una orden sin productos")
	}

	if orden.EstaCancelada() {
		return errors.New("no se puede confirmar una orden cancelada")
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
		return errors.New("no se puede cancelar una orden confirmada")
	}

	orden.Cancelar()

	return s.repository.Actualizar(orden)
}
