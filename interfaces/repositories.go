package interfaces

import "sigi/models"

// ProveedorRepository define las operaciones que el servicio necesita
// para administrar proveedores sin depender directamente de cómo
// se almacenan los datos.
type ProveedorRepository interface {
	Agregar(proveedor *models.Proveedor) error
	BuscarPorCodigo(codigo string) (*models.Proveedor, error)
	ObtenerTodos() []*models.Proveedor
	Actualizar(proveedor *models.Proveedor) error
	Eliminar(codigo string) error
}

// OrdenRepository define las operaciones de almacenamiento necesarias
// para trabajar con las órdenes de compra desde la capa de servicios.
type OrdenRepository interface {
	Agregar(orden *models.OrdenCompra) error
	BuscarPorCodigo(codigo string) (*models.OrdenCompra, error)
	ObtenerTodos() []*models.OrdenCompra
	Actualizar(orden *models.OrdenCompra) error
	Eliminar(codigo string) error
}

// TransporteRepository abstrae el almacenamiento de los transportes.
// La implementación concreta se encontrará posteriormente en repository.
type TransporteRepository interface {
	Agregar(transporte *models.Transporte) error
	BuscarPorCodigo(codigo string) (*models.Transporte, error)
	ObtenerTodos() []*models.Transporte
	Actualizar(transporte *models.Transporte) error
	Eliminar(codigo string) error
}

// ImportacionRepository define las operaciones necesarias para
// consultar y administrar las importaciones registradas.
type ImportacionRepository interface {
	Agregar(importacion *models.Importacion) error
	BuscarPorCodigo(codigo string) (*models.Importacion, error)
	ObtenerTodos() []*models.Importacion
	Actualizar(importacion *models.Importacion) error
	Eliminar(codigo string) error
}

// InventarioRepository permite desacoplar la lógica de inventario
// de la forma concreta en que los registros se almacenan.
type InventarioRepository interface {
	Agregar(inventario *models.Inventario) error
	ObtenerTodos() []*models.Inventario
	Actualizar(inventario *models.Inventario) error
}
