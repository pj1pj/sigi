package services

import (
	"sigi/interfaces"
	"sigi/models"
)

type ReporteGeneral struct {
	TotalProveedores         int
	ProveedoresActivos       int
	ProveedoresInactivos     int
	TotalOrdenes             int
	OrdenesConfirmadas       int
	OrdenesCanceladas        int
	TotalCompras             float64
	TotalTransportes         int
	TransportesActivos       int
	TransportesInactivos     int
	TotalImportaciones       int
	ImportacionesEnTransito  int
	ImportacionesEnAduana    int
	ImportacionesEnBodega    int
	TotalRegistrosInventario int
	TotalUnidadesInventario  int
}

type ReporteService struct {
	proveedorRepository   interfaces.ProveedorRepository
	ordenRepository       interfaces.OrdenRepository
	transporteRepository  interfaces.TransporteRepository
	importacionRepository interfaces.ImportacionRepository
	inventarioRepository  interfaces.InventarioRepository
}

func NuevoReporteService(
	proveedorRepository interfaces.ProveedorRepository,
	ordenRepository interfaces.OrdenRepository,
	transporteRepository interfaces.TransporteRepository,
	importacionRepository interfaces.ImportacionRepository,
	inventarioRepository interfaces.InventarioRepository,
) *ReporteService {
	return &ReporteService{
		proveedorRepository:   proveedorRepository,
		ordenRepository:       ordenRepository,
		transporteRepository:  transporteRepository,
		importacionRepository: importacionRepository,
		inventarioRepository:  inventarioRepository,
	}
}

func (s *ReporteService) General() ReporteGeneral {
	proveedores := s.proveedorRepository.ObtenerTodos()
	ordenes := s.ordenRepository.ObtenerTodos()
	transportes := s.transporteRepository.ObtenerTodos()
	importaciones := s.importacionRepository.ObtenerTodos()
	inventarios := s.inventarioRepository.ObtenerTodos()

	reporte := ReporteGeneral{
		TotalProveedores:         len(proveedores),
		TotalOrdenes:             len(ordenes),
		TotalTransportes:         len(transportes),
		TotalImportaciones:       len(importaciones),
		TotalRegistrosInventario: len(inventarios),
	}

	for _, proveedor := range proveedores {
		if proveedor.EstaActivo() {
			reporte.ProveedoresActivos++
		} else {
			reporte.ProveedoresInactivos++
		}
	}

	for _, orden := range ordenes {
		if orden.EstaConfirmada() {
			reporte.OrdenesConfirmadas++
		}

		if orden.EstaCancelada() {
			reporte.OrdenesCanceladas++
		}

		reporte.TotalCompras += orden.Total()
	}

	for _, transporte := range transportes {
		if transporte.EstaActivo() {
			reporte.TransportesActivos++
		} else {
			reporte.TransportesInactivos++
		}
	}

	for _, importacion := range importaciones {
		switch importacion.Estado() {
		case models.ImportacionEnTransito:
			reporte.ImportacionesEnTransito++

		case models.ImportacionEnAduana:
			reporte.ImportacionesEnAduana++

		case models.ImportacionLlegadaBodega:
			reporte.ImportacionesEnBodega++
		}
	}

	for _, inventario := range inventarios {
		reporte.TotalUnidadesInventario += inventario.Cantidad()
	}

	return reporte
}
