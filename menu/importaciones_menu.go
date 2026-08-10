package menu

import (
	"fmt"

	"sigi/models"
	"sigi/utils"
)

func (m *Menu) menuImportaciones() {
	for {
		fmt.Println()
		fmt.Println("----- IMPORTACIONES -----")
		fmt.Println("1. Registrar importacion")
		fmt.Println("2. Actualizar tracking")
		fmt.Println("3. Listar importaciones")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.registrarImportacion()
		case "2":
			m.actualizarTracking()
		case "3":
			m.listarImportaciones()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) registrarImportacion() {
	fmt.Println()
	fmt.Println("--- Registrar importacion ---")

	codigoOrden := utils.LeerTexto("Codigo de la orden confirmada: ")
	codigoTransporte := utils.LeerTexto("Codigo del transporte: ")
	ciudadOrigen := utils.LeerTexto("Ciudad de origen: ")
	ciudadDestino := utils.LeerTexto("Ciudad de destino: ")

	orden, err := m.ordenService.Buscar(codigoOrden)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	transporte, err := m.transporteService.Buscar(codigoTransporte)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	importacion, err := m.importacionService.Registrar(
		orden,
		transporte,
		ciudadOrigen,
		ciudadDestino,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Importacion registrada correctamente.")
	fmt.Println("Codigo generado:", importacion.Codigo())
	fmt.Println("Estado:", importacion.Estado())
	fmt.Println(
		"Fecha estimada de llegada:",
		importacion.FechaEstimadaLlegada().Format("02/01/2006"),
	)
}

func (m *Menu) actualizarTracking() {
	fmt.Println()
	fmt.Println("--- Actualizar tracking ---")

	codigo := utils.LeerTexto("Codigo de la importacion: ")

	fmt.Println("Nuevo estado:")
	fmt.Println("1. En transito")
	fmt.Println("2. En aduana")
	fmt.Println("3. Llego a bodega")

	opcion := utils.LeerTexto("Seleccione el estado: ")

	var estado models.EstadoImportacion

	switch opcion {
	case "1":
		estado = models.ImportacionEnTransito
	case "2":
		estado = models.ImportacionEnAduana
	case "3":
		estado = models.ImportacionLlegadaBodega
	default:
		fmt.Println("Estado no valido.")
		return
	}

	ubicacion := ""

	if estado == models.ImportacionLlegadaBodega {
		ubicacion = utils.LeerTexto("Ubicacion de bodega: ")
	}

	if err := m.importacionService.ActualizarTracking(
		codigo,
		estado,
		ubicacion,
	); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Tracking actualizado correctamente.")

	if estado == models.ImportacionLlegadaBodega {
		fmt.Println("La mercancia fue procesada para inventario.")
	}
}

func (m *Menu) listarImportaciones() {
	fmt.Println()
	fmt.Println("--- Importaciones ---")

	importaciones := m.importacionService.ObtenerTodos()

	if len(importaciones) == 0 {
		fmt.Println("No existen importaciones registradas.")
		return
	}

	for _, importacion := range importaciones {
		fmt.Printf(
			"Codigo: %s | Orden: %s | Transporte: %s | Origen: %s | Destino: %s | Estado: %s\n",
			importacion.Codigo(),
			importacion.Orden().Codigo(),
			importacion.Transporte().Codigo(),
			importacion.CiudadOrigen(),
			importacion.CiudadDestino(),
			importacion.Estado(),
		)
	}
}
