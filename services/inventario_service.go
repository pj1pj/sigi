package services

import (
	"errors"
	"fmt"
	"strings"

	"sigi/interfaces"
	"sigi/models"
	"sigi/utils"
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
		return fmt.Errorf("%w: la importación es obligatoria", utils.ErrDatoObligatorio)
	}

	if !importacion.LlegoABodega() {
		return fmt.Errorf("%w: la importación todavía no ha llegado a bodega", utils.ErrOperacionNoPermitida)
	}

	if strings.TrimSpace(ubicacion) == "" {
		return fmt.Errorf("%w: la ubicación de bodega es obligatoria", utils.ErrDatoObligatorio)
	}

	orden := importacion.Orden()
	if orden == nil {
		return fmt.Errorf("%w: la importación no tiene una orden asociada", utils.ErrDatoInvalido)
	}

	proveedor := orden.Proveedor()
	if proveedor == nil {
		return fmt.Errorf("%w: la orden no tiene un proveedor asociado", utils.ErrDatoInvalido)
	}

	productos := orden.Productos()

	if len(productos) == 0 {
		return fmt.Errorf("%w: la orden no contiene productos", utils.ErrDatoInvalido)
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
		return fmt.Errorf("%w: el registro de inventario es obligatorio", utils.ErrDatoObligatorio)
	}

	if cantidad <= 0 {
		return fmt.Errorf("%w: la cantidad debe ser mayor que cero", utils.ErrDatoInvalido)
	}

	inventario.AgregarCantidad(cantidad)

	return s.repository.Actualizar(inventario)
}

func (s *InventarioService) RetirarCantidad(
	inventario *models.Inventario,
	cantidad int,
) error {
	if inventario == nil {
		return fmt.Errorf("%w: el registro de inventario es obligatorio", utils.ErrDatoObligatorio)
	}

	if cantidad <= 0 {
		return fmt.Errorf("%w: la cantidad debe ser mayor que cero", utils.ErrDatoInvalido)
	}

	if !inventario.RetirarCantidad(cantidad) {
		return errors.New("la cantidad a retirar no está disponible")
	}

	return s.repository.Actualizar(inventario)
}
