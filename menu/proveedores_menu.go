package menu

import (
	"fmt"

	"sigi/utils"
)

func (m *Menu) menuProveedores() {
	for {
		fmt.Println()
		fmt.Println("----- PROVEEDORES -----")
		fmt.Println("1. Registrar proveedor")
		fmt.Println("2. Listar proveedores")
		fmt.Println("3. Activar proveedor")
		fmt.Println("4. Desactivar proveedor")
		fmt.Println("5. Eliminar proveedor")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.registrarProveedor()
		case "2":
			m.listarProveedores()
		case "3":
			m.cambiarEstadoProveedor(true)
		case "4":
			m.cambiarEstadoProveedor(false)
		case "5":
			m.eliminarProveedor()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) registrarProveedor() {
	fmt.Println()
	fmt.Println("--- Registrar proveedor ---")

	empresa := utils.LeerTexto("Empresa: ")
	pais := utils.LeerTexto("Pais: ")
	contacto := utils.LeerTexto("Contacto: ")
	telefono := utils.LeerTexto("Telefono: ")
	correo := utils.LeerTexto("Correo: ")

	proveedor, err := m.proveedorService.Registrar(
		empresa,
		pais,
		contacto,
		telefono,
		correo,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Proveedor registrado correctamente.")
	fmt.Println("Codigo generado:", proveedor.Codigo())
}

func (m *Menu) listarProveedores() {
	fmt.Println()
	fmt.Println("--- Proveedores ---")

	proveedores := m.proveedorService.ObtenerTodos()

	if len(proveedores) == 0 {
		fmt.Println("No existen proveedores registrados.")
		return
	}

	for _, proveedor := range proveedores {
		fmt.Printf(
			"Codigo: %s | Empresa: %s | Pais: %s | Estado: %s\n",
			proveedor.Codigo(),
			proveedor.Empresa(),
			proveedor.Pais(),
			proveedor.Estado(),
		)
	}
}

func (m *Menu) cambiarEstadoProveedor(activar bool) {
	codigo := utils.LeerTexto("Codigo del proveedor: ")

	var err error

	if activar {
		err = m.proveedorService.Activar(codigo)
	} else {
		err = m.proveedorService.Desactivar(codigo)
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Estado actualizado correctamente.")
}

func (m *Menu) eliminarProveedor() {
	codigo := utils.LeerTexto("Codigo del proveedor: ")

	if !utils.Confirmar("Desea eliminar este proveedor") {
		fmt.Println("Operacion cancelada.")
		return
	}

	if err := m.proveedorService.Eliminar(codigo); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Proveedor eliminado correctamente.")
}
