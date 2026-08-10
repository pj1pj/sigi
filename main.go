package main

import (
	"sigi/menu"
	"sigi/repository"
	"sigi/services"
	"sigi/utils"
)

func main() {
	proveedorRepository := repository.NuevoProveedorRepository()
	ordenRepository := repository.NuevaOrdenRepository()
	transporteRepository := repository.NuevoTransporteRepository()
	importacionRepository := repository.NuevaImportacionRepository()
	inventarioRepository := repository.NuevoInventarioRepository()

	generadorCodigo := utils.NuevoGeneradorCodigo()

	proveedorService := services.NuevoProveedorService(
		proveedorRepository,
		generadorCodigo,
	)

	ordenService := services.NuevaOrdenService(
		ordenRepository,
		generadorCodigo,
	)

	transporteService := services.NuevoTransporteService(
		transporteRepository,
		generadorCodigo,
	)

	inventarioService := services.NuevoInventarioService(
		inventarioRepository,
	)

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

	sistemaMenu := menu.NuevoMenu(
		proveedorService,
		ordenService,
		transporteService,
		importacionService,
		inventarioService,
		reporteService,
	)

	sistemaMenu.Ejecutar()
}
