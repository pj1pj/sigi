// Package app construye las dependencias compartidas por las interfaces de SIGI.
package app

import (
	"sigi/repository"
	"sigi/services"
	"sigi/utils"
)

// Sistema agrupa los servicios de negocio que usan tanto la consola como la API.
type Sistema struct {
	ProveedorService   *services.ProveedorService
	OrdenService       *services.OrdenService
	TransporteService  *services.TransporteService
	ImportacionService *services.ImportacionService
	InventarioService  *services.InventarioService
	ReporteService     *services.ReporteService
}

// NuevoSistema crea una instancia independiente de SIGI con repositorios en memoria.
func NuevoSistema() *Sistema {
	proveedorRepository := repository.NuevoProveedorRepository()
	ordenRepository := repository.NuevaOrdenRepository()
	transporteRepository := repository.NuevoTransporteRepository()
	importacionRepository := repository.NuevaImportacionRepository()
	inventarioRepository := repository.NuevoInventarioRepository()
	generadorCodigo := utils.NuevoGeneradorCodigo()

	proveedorService := services.NuevoProveedorService(proveedorRepository, generadorCodigo)
	ordenService := services.NuevaOrdenService(ordenRepository, generadorCodigo)
	transporteService := services.NuevoTransporteService(transporteRepository, generadorCodigo)
	inventarioService := services.NuevoInventarioService(inventarioRepository)
	importacionService := services.NuevaImportacionService(
		importacionRepository,
		inventarioService,
		generadorCodigo,
	)
	reporteService := services.NuevoReporteService(
		proveedorRepository,
		ordenRepository,
		transporteRepository,
		importacionRepository,
		inventarioRepository,
	)

	return &Sistema{
		ProveedorService:   proveedorService,
		OrdenService:       ordenService,
		TransporteService:  transporteService,
		ImportacionService: importacionService,
		InventarioService:  inventarioService,
		ReporteService:     reporteService,
	}
}
