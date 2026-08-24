package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"sigi/api"
	"sigi/app"
	"sigi/menu"
)

func main() {
	modo := flag.String("modo", "consola", "modo de ejecución: consola o api")
	flag.Parse()

	sistema := app.NuevoSistema()

	switch *modo {
	case "consola":
		sistemaMenu := menu.NuevoMenu(
			sistema.ProveedorService,
			sistema.OrdenService,
			sistema.TransporteService,
			sistema.ImportacionService,
			sistema.InventarioService,
			sistema.ReporteService,
		)
		sistemaMenu.Ejecutar()

	case "api":
		servidor := api.NuevoServidor(
			sistema.ProveedorService,
			sistema.OrdenService,
			sistema.TransporteService,
			sistema.ImportacionService,
			sistema.InventarioService,
			sistema.ReporteService,
		)
		fmt.Println("API SIGI disponible en http://localhost:8080")
		log.Fatal(http.ListenAndServe("localhost:8080", servidor.Handler()))

	default:
		log.Fatalf("modo no válido: %q. Use consola o api", *modo)
	}
}
