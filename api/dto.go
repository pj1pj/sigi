package api

import (
	"time"

	"sigi/models"
	"sigi/services"
)

// Los DTO separan la representación JSON de los modelos de dominio, cuyos
// campos permanecen privados para conservar su encapsulamiento.
type proveedorRequest struct {
	Empresa  string `json:"empresa"`
	Pais     string `json:"pais"`
	Contacto string `json:"contacto"`
	Telefono string `json:"telefono"`
	Correo   string `json:"correo"`
}

type proveedorResponse struct {
	Codigo   string `json:"codigo"`
	Empresa  string `json:"empresa"`
	Pais     string `json:"pais"`
	Contacto string `json:"contacto"`
	Telefono string `json:"telefono"`
	Correo   string `json:"correo"`
	Estado   string `json:"estado"`
}

type estadoProveedorRequest struct {
	Activo *bool `json:"activo"`
}

type transporteRequest struct {
	Tipo     string `json:"tipo"`
	Empresa  string `json:"empresa"`
	Pais     string `json:"pais"`
	Contacto string `json:"contacto"`
	Telefono string `json:"telefono"`
	Correo   string `json:"correo"`
}

type transporteResponse struct {
	Codigo   string `json:"codigo"`
	Tipo     string `json:"tipo"`
	Empresa  string `json:"empresa"`
	Pais     string `json:"pais"`
	Contacto string `json:"contacto"`
	Telefono string `json:"telefono"`
	Correo   string `json:"correo"`
	Estado   string `json:"estado"`
}

type ordenRequest struct {
	CodigoProveedor string `json:"codigo_proveedor"`
}

type productoRequest struct {
	Nombre         string  `json:"nombre"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
}

type confirmacionOrdenRequest struct {
	Confirmada *bool `json:"confirmada"`
}

type productoResponse struct {
	Nombre         string  `json:"nombre"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
	Subtotal       float64 `json:"subtotal"`
}

type ordenResponse struct {
	Codigo        string             `json:"codigo"`
	Proveedor     proveedorResponse  `json:"proveedor"`
	Productos     []productoResponse `json:"productos"`
	Total         float64            `json:"total"`
	FechaCreacion string             `json:"fecha_creacion"`
	Estado        string             `json:"estado"`
}

type importacionRequest struct {
	CodigoOrden      string `json:"codigo_orden"`
	CodigoTransporte string `json:"codigo_transporte"`
	CiudadOrigen     string `json:"ciudad_origen"`
	CiudadDestino    string `json:"ciudad_destino"`
}

type trackingRequest struct {
	Estado    string `json:"estado"`
	Ubicacion string `json:"ubicacion"`
}

type importacionResponse struct {
	Codigo               string `json:"codigo"`
	CodigoOrden          string `json:"codigo_orden"`
	CodigoTransporte     string `json:"codigo_transporte"`
	CiudadOrigen         string `json:"ciudad_origen"`
	CiudadDestino        string `json:"ciudad_destino"`
	FechaRegistro        string `json:"fecha_registro"`
	FechaEstimadaLlegada string `json:"fecha_estimada_llegada"`
	Estado               string `json:"estado"`
}

type inventarioResponse struct {
	Producto    productoResponse  `json:"producto"`
	Cantidad    int               `json:"cantidad"`
	CodigoOrden string            `json:"codigo_orden"`
	Proveedor   proveedorResponse `json:"proveedor"`
	Importacion string            `json:"codigo_importacion"`
	Ubicacion   string            `json:"ubicacion"`
	Estado      string            `json:"estado"`
}

type reporteGeneralResponse struct {
	TotalProveedores           int     `json:"total_proveedores"`
	ProveedoresActivos         int     `json:"proveedores_activos"`
	ProveedoresInactivos       int     `json:"proveedores_inactivos"`
	TotalOrdenes               int     `json:"total_ordenes"`
	OrdenesConfirmadas         int     `json:"ordenes_confirmadas"`
	OrdenesCanceladas          int     `json:"ordenes_canceladas"`
	TotalCompras               float64 `json:"total_compras"`
	TotalTransportes           int     `json:"total_transportes"`
	TransportesActivos         int     `json:"transportes_activos"`
	TransportesInactivos       int     `json:"transportes_inactivos"`
	TotalImportaciones         int     `json:"total_importaciones"`
	ImportacionesEnPreparacion int     `json:"importaciones_en_preparacion"`
	ImportacionesEnTransito    int     `json:"importaciones_en_transito"`
	ImportacionesEnAduana      int     `json:"importaciones_en_aduana"`
	ImportacionesEnBodega      int     `json:"importaciones_en_bodega"`
	TotalRegistrosInventario   int     `json:"total_registros_inventario"`
	TotalUnidadesInventario    int     `json:"total_unidades_inventario"`
}

func nuevoProveedorResponse(proveedor *models.Proveedor) proveedorResponse {
	return proveedorResponse{
		Codigo: proveedor.Codigo(), Empresa: proveedor.Empresa(), Pais: proveedor.Pais(),
		Contacto: proveedor.Contacto(), Telefono: proveedor.Telefono(), Correo: proveedor.Correo(),
		Estado: string(proveedor.Estado()),
	}
}

func nuevoTransporteResponse(transporte *models.Transporte) transporteResponse {
	return transporteResponse{
		Codigo: transporte.Codigo(), Tipo: string(transporte.Tipo()), Empresa: transporte.Empresa(),
		Pais: transporte.Pais(), Contacto: transporte.Contacto(), Telefono: transporte.Telefono(),
		Correo: transporte.Correo(), Estado: string(transporte.Estado()),
	}
}

func nuevoProductoResponse(producto *models.Producto) productoResponse {
	return productoResponse{
		Nombre: producto.Nombre(), Cantidad: producto.Cantidad(),
		PrecioUnitario: producto.PrecioUnitario(), Subtotal: producto.Subtotal(),
	}
}

func nuevaOrdenResponse(orden *models.OrdenCompra) ordenResponse {
	productos := make([]productoResponse, 0, len(orden.Productos()))
	for _, producto := range orden.Productos() {
		if producto != nil {
			productos = append(productos, nuevoProductoResponse(producto))
		}
	}

	return ordenResponse{
		Codigo: orden.Codigo(), Proveedor: nuevoProveedorResponse(orden.Proveedor()),
		Productos: productos, Total: orden.Total(),
		FechaCreacion: orden.FechaCreacion().Format(time.RFC3339), Estado: string(orden.Estado()),
	}
}

func nuevaImportacionResponse(importacion *models.Importacion) importacionResponse {
	return importacionResponse{
		Codigo: importacion.Codigo(), CodigoOrden: importacion.Orden().Codigo(),
		CodigoTransporte: importacion.Transporte().Codigo(), CiudadOrigen: importacion.CiudadOrigen(),
		CiudadDestino:        importacion.CiudadDestino(),
		FechaRegistro:        importacion.FechaRegistro().Format(time.RFC3339),
		FechaEstimadaLlegada: importacion.FechaEstimadaLlegada().Format(time.RFC3339),
		Estado:               string(importacion.Estado()),
	}
}

func nuevoInventarioResponse(inventario *models.Inventario) inventarioResponse {
	return inventarioResponse{
		Producto: nuevoProductoResponse(inventario.Producto()), Cantidad: inventario.Cantidad(),
		CodigoOrden: inventario.Orden().Codigo(), Proveedor: nuevoProveedorResponse(inventario.Proveedor()),
		Importacion: inventario.Importacion().Codigo(), Ubicacion: inventario.Ubicacion(),
		Estado: string(inventario.Estado()),
	}
}

func nuevoReporteGeneralResponse(reporte services.ReporteGeneral) reporteGeneralResponse {
	return reporteGeneralResponse{
		TotalProveedores: reporte.TotalProveedores, ProveedoresActivos: reporte.ProveedoresActivos,
		ProveedoresInactivos: reporte.ProveedoresInactivos, TotalOrdenes: reporte.TotalOrdenes,
		OrdenesConfirmadas: reporte.OrdenesConfirmadas, OrdenesCanceladas: reporte.OrdenesCanceladas,
		TotalCompras: reporte.TotalCompras, TotalTransportes: reporte.TotalTransportes,
		TransportesActivos: reporte.TransportesActivos, TransportesInactivos: reporte.TransportesInactivos,
		TotalImportaciones: reporte.TotalImportaciones, ImportacionesEnPreparacion: reporte.ImportacionesEnPreparacion,
		ImportacionesEnTransito: reporte.ImportacionesEnTransito,
		ImportacionesEnAduana:   reporte.ImportacionesEnAduana, ImportacionesEnBodega: reporte.ImportacionesEnBodega,
		TotalRegistrosInventario: reporte.TotalRegistrosInventario, TotalUnidadesInventario: reporte.TotalUnidadesInventario,
	}
}
