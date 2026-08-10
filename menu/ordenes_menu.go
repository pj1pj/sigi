package menu

import (
	"fmt"

	"sigi/models"
	"sigi/utils"
)

func (m *Menu) menuOrdenes() {
	for {
		fmt.Println()
		fmt.Println("----- ORDENES DE COMPRA -----")
		fmt.Println("1. Crear orden")
		fmt.Println("2. Agregar producto")
		fmt.Println("3. Confirmar orden")
		fmt.Println("4. Cancelar orden")
		fmt.Println("5. Listar ordenes")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.crearOrden()
		case "2":
			m.agregarProductoOrden()
		case "3":
			m.confirmarOrden()
		case "4":
			m.cancelarOrden()
		case "5":
			m.listarOrdenes()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) crearOrden() {
	fmt.Println()
	fmt.Println("--- Crear orden de compra ---")

	codigoProveedor := utils.LeerTexto("Codigo del proveedor: ")

	proveedor, err := m.proveedorService.Buscar(codigoProveedor)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	orden, err := m.ordenService.Crear(proveedor)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Orden creada correctamente.")
	fmt.Println("Codigo generado:", orden.Codigo())
	fmt.Println("Estado:", orden.Estado())
}

func (m *Menu) agregarProductoOrden() {
	fmt.Println()
	fmt.Println("--- Agregar producto a orden ---")

	codigoOrden := utils.LeerTexto("Codigo de la orden: ")
	nombre := utils.LeerTexto("Nombre del producto: ")

	cantidad, err := utils.LeerEntero("Cantidad: ")
	if err != nil {
		fmt.Println("Error: la cantidad debe ser un numero entero.")
		return
	}

	precio, err := utils.LeerDecimal("Precio unitario: ")
	if err != nil {
		fmt.Println("Error: el precio debe ser un numero valido.")
		return
	}

	producto := models.NuevoProducto(nombre, cantidad, precio)

	if err := m.ordenService.AgregarProducto(codigoOrden, producto); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Producto agregado correctamente.")
}

func (m *Menu) confirmarOrden() {
	fmt.Println()
	fmt.Println("--- Confirmar orden ---")

	codigo := utils.LeerTexto("Codigo de la orden: ")

	if !utils.Confirmar("Desea confirmar esta orden") {
		fmt.Println("Operacion cancelada.")
		return
	}

	if err := m.ordenService.Confirmar(codigo); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Orden confirmada correctamente.")
}

func (m *Menu) cancelarOrden() {
	fmt.Println()
	fmt.Println("--- Cancelar orden ---")

	codigo := utils.LeerTexto("Codigo de la orden: ")

	if !utils.Confirmar("Desea cancelar esta orden") {
		fmt.Println("Operacion cancelada.")
		return
	}

	if err := m.ordenService.Cancelar(codigo); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Orden cancelada correctamente.")
}

func (m *Menu) listarOrdenes() {
	fmt.Println()
	fmt.Println("--- Ordenes de compra ---")

	ordenes := m.ordenService.ObtenerTodos()

	if len(ordenes) == 0 {
		fmt.Println("No existen ordenes registradas.")
		return
	}

	for _, orden := range ordenes {
		fmt.Printf(
			"Codigo: %s | Proveedor: %s | Productos: %d | Total: %.2f | Estado: %s\n",
			orden.Codigo(),
			orden.Proveedor().Empresa(),
			len(orden.Productos()),
			orden.Total(),
			orden.Estado(),
		)
	}
}
