package menu

import (
	"fmt"

	"sigi/models"
	"sigi/utils"
)

func (m *Menu) menuTransportes() {
	for {
		fmt.Println()
		fmt.Println("----- TRANSPORTES -----")
		fmt.Println("1. Registrar transporte")
		fmt.Println("2. Listar transportes")
		fmt.Println("3. Activar transporte")
		fmt.Println("4. Desactivar transporte")
		fmt.Println("5. Eliminar transporte")
		fmt.Println("0. Volver")

		opcion := utils.LeerTexto("Seleccione una opcion: ")

		switch opcion {
		case "1":
			m.registrarTransporte()
		case "2":
			m.listarTransportes()
		case "3":
			m.cambiarEstadoTransporte(true)
		case "4":
			m.cambiarEstadoTransporte(false)
		case "5":
			m.eliminarTransporte()
		case "0":
			return
		default:
			fmt.Println("Opcion no valida.")
		}
	}
}

func (m *Menu) registrarTransporte() {
	fmt.Println()
	fmt.Println("--- Registrar transporte ---")

	fmt.Println("Tipo de transporte:")
	fmt.Println("1. Maritimo")
	fmt.Println("2. Aereo")
	fmt.Println("3. Terrestre")

	opcionTipo := utils.LeerTexto("Seleccione el tipo: ")

	var tipo models.TipoTransporte

	switch opcionTipo {
	case "1":
		tipo = models.TransporteMaritimo
	case "2":
		tipo = models.TransporteAereo
	case "3":
		tipo = models.TransporteTerrestre
	default:
		fmt.Println("Tipo de transporte no valido.")
		return
	}

	empresa := utils.LeerTexto("Empresa: ")
	pais := utils.LeerTexto("Pais: ")
	contacto := utils.LeerTexto("Contacto: ")
	telefono := utils.LeerTexto("Telefono: ")
	correo := utils.LeerTexto("Correo: ")

	transporte, err := m.transporteService.Registrar(
		tipo,
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

	fmt.Println("Transporte registrado correctamente.")
	fmt.Println("Codigo generado:", transporte.Codigo())
}

func (m *Menu) listarTransportes() {
	fmt.Println()
	fmt.Println("--- Transportes ---")

	transportes := m.transporteService.ObtenerTodos()

	if len(transportes) == 0 {
		fmt.Println("No existen transportes registrados.")
		return
	}

	for _, transporte := range transportes {
		fmt.Printf(
			"Codigo: %s | Tipo: %s | Empresa: %s | Pais: %s | Estado: %s\n",
			transporte.Codigo(),
			transporte.Tipo(),
			transporte.Empresa(),
			transporte.Pais(),
			transporte.Estado(),
		)
	}
}

func (m *Menu) cambiarEstadoTransporte(activar bool) {
	codigo := utils.LeerTexto("Codigo del transporte: ")

	var err error

	if activar {
		err = m.transporteService.Activar(codigo)
	} else {
		err = m.transporteService.Desactivar(codigo)
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Estado actualizado correctamente.")
}

func (m *Menu) eliminarTransporte() {
	codigo := utils.LeerTexto("Codigo del transporte: ")

	if !utils.Confirmar("Desea eliminar este transporte") {
		fmt.Println("Operacion cancelada.")
		return
	}

	if err := m.transporteService.Eliminar(codigo); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transporte eliminado correctamente.")
}
