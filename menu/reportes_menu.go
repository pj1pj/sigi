package menu

import (
	"fmt"

	"sigi/utils"
)

func (m *Menu) menuReportes() {
	for {
		fmt.Println()
		fmt.Println("----- REPORTES -----")
		fmt.Println("1. Reporte general")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.reporteGeneral()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) reporteGeneral() {
	fmt.Println()
	fmt.Println("========== REPORTE GENERAL SIGI ==========")

	reporte := m.reporteService.General()

	fmt.Println()
	fmt.Println("PROVEEDORES")
	fmt.Printf("Total: %d\n", reporte.TotalProveedores)
	fmt.Printf("Activos: %d\n", reporte.ProveedoresActivos)
	fmt.Printf("Inactivos: %d\n", reporte.ProveedoresInactivos)

	fmt.Println()
	fmt.Println("ORDENES DE COMPRA")
	fmt.Printf("Total: %d\n", reporte.TotalOrdenes)
	fmt.Printf("Confirmadas: %d\n", reporte.OrdenesConfirmadas)
	fmt.Printf("Canceladas: %d\n", reporte.OrdenesCanceladas)
	fmt.Printf("Total de compras: $%.2f\n", reporte.TotalCompras)

	fmt.Println()
	fmt.Println("TRANSPORTES")
	fmt.Printf("Total: %d\n", reporte.TotalTransportes)
	fmt.Printf("Activos: %d\n", reporte.TransportesActivos)
	fmt.Printf("Inactivos: %d\n", reporte.TransportesInactivos)

	fmt.Println()
	fmt.Println("IMPORTACIONES")
	fmt.Printf("Total: %d\n", reporte.TotalImportaciones)
	fmt.Printf("En preparación: %d\n", reporte.ImportacionesEnPreparacion)
	fmt.Printf("En transito: %d\n", reporte.ImportacionesEnTransito)
	fmt.Printf("En aduana: %d\n", reporte.ImportacionesEnAduana)
	fmt.Printf("En bodega: %d\n", reporte.ImportacionesEnBodega)

	fmt.Println()
	fmt.Println("INVENTARIO")
	fmt.Printf("Registros: %d\n", reporte.TotalRegistrosInventario)
	fmt.Printf("Unidades disponibles: %d\n", reporte.TotalUnidadesInventario)

	fmt.Println()
	fmt.Println("==========================================")

	utils.LeerTexto("Presione ENTER para continuar...")
}
