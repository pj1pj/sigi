package services

import (
	"errors"
	"strings"

	"sigi/interfaces"
	"sigi/models"
)

type InventarioService struct {
	repository interfaces.InventarioRepository
}

func NuevoInventarioService(repository interfaces.InventarioRepository) *InventarioService {
	return &InventarioService{
		repository: repository,
	}
}

func (s *InventarioService) ObtenerTodos() []*models.Inventario {
	return s.repository.ObtenerTodos()
}

// ProcesarLlegada convierte los productos de una importación que llegó
// a bodega en registros de inventario. La creación se realiza a partir
// de la orden asociada, por lo que el usuario no registra existencias
// manualmente.
func (s *InventarioService) ProcesarLlegada(
	importacion *models.Importacion,
	ubicacion string,
) error {
	if importacion == nil {
		return errors.New("la importación es obligatoria")
	}

	if !importacion.LlegoABodega() {
		return errors.New("la importación todavía no ha llegado a bodega")
	}

	if strings.TrimSpace(ubicacion) == "" {
		return errors.New("la ubicación de bodega es obligatoria")
	}

	orden := importacion.Orden()
	if orden == nil {
		return errors.New("la importación no tiene una orden asociada")
	}

	proveedor := orden.Proveedor()
	if proveedor == nil {
		return errors.New("la orden no tiene un proveedor asociado")
	}

	productos := orden.Productos()

	if len(productos) == 0 {
		return errors.New("la orden no contiene productos")
	}

	for _, producto := range productos {
		if producto == nil {
			continue
		}

		inventario := models.NuevoInventario(
			producto,
			producto.Cantidad(),
			orden,
			proveedor,
			importacion,
			ubicacion,
		)

		if err := s.repository.Agregar(inventario); err != nil {
			return err
		}
	}

	return nil
}

func (s *InventarioService) AgregarCantidad(
	inventario *models.Inventario,
	cantidad int,
) error {
	if inventario == nil {
		return errors.New("el registro de inventario es obligatorio")
	}

	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}

	inventario.AgregarCantidad(cantidad)

	return s.repository.Actualizar(inventario)
}

func (s *InventarioService) RetirarCantidad(
	inventario *models.Inventario,
	cantidad int,
) error {
	if inventario == nil {
		return errors.New("el registro de inventario es obligatorio")
	}

	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor que cero")
	}

	if !inventario.RetirarCantidad(cantidad) {
		return errors.New("la cantidad a retirar no está disponible")
	}

	return s.repository.Actualizar(inventario)
}
