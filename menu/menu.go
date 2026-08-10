package menu

import (
	"fmt"

	"sigi/services"
	"sigi/utils"
)

type Menu struct {
	proveedorService   *services.ProveedorService
	ordenService       *services.OrdenService
	transporteService  *services.TransporteService
	importacionService *services.ImportacionService
	inventarioService  *services.InventarioService
	reporteService     *services.ReporteService
}

func NuevoMenu(
	proveedorService *services.ProveedorService,
	ordenService *services.OrdenService,
	transporteService *services.TransporteService,
	importacionService *services.ImportacionService,
	inventarioService *services.InventarioService,
	reporteService *services.ReporteService,
) *Menu {
	return &Menu{
		proveedorService:   proveedorService,
		ordenService:       ordenService,
		transporteService:  transporteService,
		importacionService: importacionService,
		inventarioService:  inventarioService,
		reporteService:     reporteService,
	}
}

func (m *Menu) Ejecutar() {
	for {
		m.mostrarPrincipal()

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.menuProveedores()
		case "2":
			m.menuOrdenes()
		case "3":
			m.menuTransportes()
		case "4":
			m.menuImportaciones()
		case "5":
			m.menuInventario()
		case "6":
			m.menuReportes()
		case "0":
			fmt.Println("Saliendo de SIGI...")
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) mostrarPrincipal() {
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println("       SISTEMA SIGI")
	fmt.Println("====================================")
	fmt.Println("1. Gestion de proveedores")
	fmt.Println("2. Gestion de ordenes de compra")
	fmt.Println("3. Gestion de transportes")
	fmt.Println("4. Gestion de importaciones")
	fmt.Println("5. Gestion de inventario")
	fmt.Println("6. Reportes")
	fmt.Println("0. Salir")
	fmt.Println("====================================")
}
