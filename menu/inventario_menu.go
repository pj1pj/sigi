package menu

import (
	"fmt"

	"sigi/utils"
)

func (m *Menu) menuInventario() {
	for {
		fmt.Println()
		fmt.Println("----- INVENTARIO -----")
		fmt.Println("1. Consultar inventario")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.listarInventario()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) listarInventario() {
	fmt.Println()
	fmt.Println("--- Inventario ---")

	inventarios := m.inventarioService.ObtenerTodos()

	if len(inventarios) == 0 {
		fmt.Println("No existen registros de inventario.")
		return
	}

	for indice, inventario := range inventarios {
		fmt.Printf(
			"%d. Producto: %s | Cantidad: %d | Proveedor: %s | Importacion: %s | Ubicacion: %s\n",
			indice+1,
			inventario.Producto().Nombre(),
			inventario.Cantidad(),
			inventario.Proveedor().Empresa(),
			inventario.Importacion().Codigo(),
			inventario.Ubicacion(),
		)
	}
}
