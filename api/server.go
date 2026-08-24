// Package api expone las funcionalidades de SIGI mediante HTTP y JSON.
package api

import (
	"net/http"
	"strings"
	"sync"

	"sigi/models"
	"sigi/services"
)

// Servidor enlaza las rutas HTTP con los servicios de negocio existentes.
type Servidor struct {
	proveedorService   *services.ProveedorService
	ordenService       *services.OrdenService
	transporteService  *services.TransporteService
	importacionService *services.ImportacionService
	inventarioService  *services.InventarioService
	reporteService     *services.ReporteService
	mux                *http.ServeMux
	mu                 sync.RWMutex
}

func NuevoServidor(
	proveedorService *services.ProveedorService,
	ordenService *services.OrdenService,
	transporteService *services.TransporteService,
	importacionService *services.ImportacionService,
	inventarioService *services.InventarioService,
	reporteService *services.ReporteService,
) *Servidor {
	servidor := &Servidor{
		proveedorService: proveedorService, ordenService: ordenService,
		transporteService: transporteService, importacionService: importacionService,
		inventarioService: inventarioService, reporteService: reporteService,
		mux: http.NewServeMux(),
	}
	servidor.registrarRutas()
	return servidor
}

// Handler devuelve el manejador que debe recibir net/http.
func (s *Servidor) Handler() http.Handler {
	// Los modelos se comparten entre solicitudes y sus métodos mutan punteros.
	// Este bloqueo de aplicación evita carreras entre lecturas y operaciones de escritura.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			s.mu.RLock()
			defer s.mu.RUnlock()
		} else {
			s.mu.Lock()
			defer s.mu.Unlock()
		}
		s.mux.ServeHTTP(w, r)
	})
}

func (s *Servidor) registrarRutas() {
	s.mux.HandleFunc("/api/v1/proveedores", s.manejarProveedores)
	s.mux.HandleFunc("/api/v1/proveedores/", s.manejarProveedorPorCodigo)
	s.mux.HandleFunc("/api/v1/transportes", s.manejarTransportes)
	s.mux.HandleFunc("/api/v1/ordenes", s.manejarOrdenes)
	s.mux.HandleFunc("/api/v1/ordenes/", s.manejarOrdenPorCodigo)
	s.mux.HandleFunc("/api/v1/importaciones", s.manejarImportaciones)
	s.mux.HandleFunc("/api/v1/importaciones/", s.manejarImportacionPorCodigo)
	s.mux.HandleFunc("/api/v1/inventario", s.manejarInventario)
	s.mux.HandleFunc("/api/v1/reportes/general", s.manejarReporteGeneral)
	s.mux.HandleFunc("/api/", s.manejarRutaNoEncontrada)
}

func (s *Servidor) manejarProveedores(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var solicitud proveedorRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}

		proveedor, err := s.proveedorService.Registrar(
			solicitud.Empresa, solicitud.Pais, solicitud.Contacto, solicitud.Telefono, solicitud.Correo,
		)
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		responderJSON(w, http.StatusCreated, nuevoProveedorResponse(proveedor))

	case http.MethodGet:
		proveedores := s.proveedorService.ObtenerTodos()
		respuesta := make([]proveedorResponse, 0, len(proveedores))
		for _, proveedor := range proveedores {
			respuesta = append(respuesta, nuevoProveedorResponse(proveedor))
		}
		responderJSON(w, http.StatusOK, respuesta)

	default:
		responderMetodoNoPermitido(w, "GET, POST")
	}
}

func (s *Servidor) manejarProveedorPorCodigo(w http.ResponseWriter, r *http.Request) {
	partes := partesRuta(r.URL.Path, "/api/v1/proveedores/")
	if len(partes) == 1 && r.Method == http.MethodGet {
		proveedor, err := s.proveedorService.Buscar(partes[0])
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		responderJSON(w, http.StatusOK, nuevoProveedorResponse(proveedor))
		return
	}
	if len(partes) == 1 {
		responderMetodoNoPermitido(w, "GET")
		return
	}

	if len(partes) == 2 && partes[1] == "estado" {
		if r.Method != http.MethodPatch {
			responderMetodoNoPermitido(w, "PATCH")
			return
		}

		var solicitud estadoProveedorRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}
		if solicitud.Activo == nil {
			responderError(w, http.StatusBadRequest, "el campo activo es obligatorio")
			return
		}

		var err error
		if *solicitud.Activo {
			err = s.proveedorService.Activar(partes[0])
		} else {
			err = s.proveedorService.Desactivar(partes[0])
		}
		if err != nil {
			responderErrorServicio(w, err)
			return
		}

		proveedor, err := s.proveedorService.Buscar(partes[0])
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		responderJSON(w, http.StatusOK, nuevoProveedorResponse(proveedor))
		return
	}

	responderError(w, http.StatusNotFound, "ruta no encontrada")
}

func (s *Servidor) manejarTransportes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var solicitud transporteRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}

		transporte, err := s.transporteService.Registrar(
			models.TipoTransporte(solicitud.Tipo), solicitud.Empresa, solicitud.Pais,
			solicitud.Contacto, solicitud.Telefono, solicitud.Correo,
		)
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		responderJSON(w, http.StatusCreated, nuevoTransporteResponse(transporte))

	case http.MethodGet:
		transportes := s.transporteService.ObtenerTodos()
		respuesta := make([]transporteResponse, 0, len(transportes))
		for _, transporte := range transportes {
			respuesta = append(respuesta, nuevoTransporteResponse(transporte))
		}
		responderJSON(w, http.StatusOK, respuesta)

	default:
		responderMetodoNoPermitido(w, "GET, POST")
	}
}

func (s *Servidor) manejarOrdenes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var solicitud ordenRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}

		proveedor, err := s.proveedorService.Buscar(solicitud.CodigoProveedor)
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		orden, err := s.ordenService.Crear(proveedor)
		if err != nil {
			responderErrorServicio(w, err)
			return
		}
		responderJSON(w, http.StatusCreated, nuevaOrdenResponse(orden))

	case http.MethodGet:
		ordenes := s.ordenService.ObtenerTodos()
		respuesta := make([]ordenResponse, 0, len(ordenes))
		for _, orden := range ordenes {
			respuesta = append(respuesta, nuevaOrdenResponse(orden))
		}
		responderJSON(w, http.StatusOK, respuesta)

	default:
		responderMetodoNoPermitido(w, "GET, POST")
	}
}

func (s *Servidor) manejarOrdenPorCodigo(w http.ResponseWriter, r *http.Request) {
	partes := partesRuta(r.URL.Path, "/api/v1/ordenes/")
	if len(partes) != 2 {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	codigo := partes[0]
	switch partes[1] {
	case "productos":
		if r.Method != http.MethodPost {
			responderMetodoNoPermitido(w, "POST")
			return
		}
		var solicitud productoRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}
		err := s.ordenService.AgregarProducto(
			codigo, models.NuevoProducto(solicitud.Nombre, solicitud.Cantidad, solicitud.PrecioUnitario),
		)
		if err != nil {
			responderErrorServicio(w, err)
			return
		}

	case "confirmacion":
		if r.Method != http.MethodPatch {
			responderMetodoNoPermitido(w, "PATCH")
			return
		}
		var solicitud confirmacionOrdenRequest
		if !decodificarJSON(w, r, &solicitud) {
			return
		}
		if solicitud.Confirmada == nil || !*solicitud.Confirmada {
			responderError(w, http.StatusBadRequest, "confirmada debe ser true")
			return
		}
		if err := s.ordenService.Confirmar(codigo); err != nil {
			responderErrorServicio(w, err)
			return
		}

	default:
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}

	orden, err := s.ordenService.Buscar(codigo)
	if err != nil {
		responderErrorServicio(w, err)
		return
	}
	responderJSON(w, http.StatusOK, nuevaOrdenResponse(orden))
}

func (s *Servidor) manejarImportaciones(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responderMetodoNoPermitido(w, "POST")
		return
	}

	var solicitud importacionRequest
	if !decodificarJSON(w, r, &solicitud) {
		return
	}

	orden, err := s.ordenService.Buscar(solicitud.CodigoOrden)
	if err != nil {
		responderErrorServicio(w, err)
		return
	}
	transporte, err := s.transporteService.Buscar(solicitud.CodigoTransporte)
	if err != nil {
		responderErrorServicio(w, err)
		return
	}
	importacion, err := s.importacionService.Registrar(
		orden, transporte, solicitud.CiudadOrigen, solicitud.CiudadDestino,
	)
	if err != nil {
		responderErrorServicio(w, err)
		return
	}
	responderJSON(w, http.StatusCreated, nuevaImportacionResponse(importacion))
}

func (s *Servidor) manejarImportacionPorCodigo(w http.ResponseWriter, r *http.Request) {
	partes := partesRuta(r.URL.Path, "/api/v1/importaciones/")
	if len(partes) != 2 || partes[1] != "tracking" {
		responderError(w, http.StatusNotFound, "ruta no encontrada")
		return
	}
	if r.Method != http.MethodPatch {
		responderMetodoNoPermitido(w, "PATCH")
		return
	}

	var solicitud trackingRequest
	if !decodificarJSON(w, r, &solicitud) {
		return
	}
	if err := s.importacionService.ActualizarTracking(
		partes[0], models.EstadoImportacion(solicitud.Estado), solicitud.Ubicacion,
	); err != nil {
		responderErrorServicio(w, err)
		return
	}

	importacion, err := s.importacionService.Buscar(partes[0])
	if err != nil {
		responderErrorServicio(w, err)
		return
	}
	responderJSON(w, http.StatusOK, nuevaImportacionResponse(importacion))
}

func (s *Servidor) manejarInventario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderMetodoNoPermitido(w, "GET")
		return
	}

	inventarios := s.inventarioService.ObtenerTodos()
	respuesta := make([]inventarioResponse, 0, len(inventarios))
	for _, inventario := range inventarios {
		respuesta = append(respuesta, nuevoInventarioResponse(inventario))
	}
	responderJSON(w, http.StatusOK, respuesta)
}

func (s *Servidor) manejarReporteGeneral(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		responderMetodoNoPermitido(w, "GET")
		return
	}

	responderJSON(w, http.StatusOK, nuevoReporteGeneralResponse(s.reporteService.General()))
}

func (s *Servidor) manejarRutaNoEncontrada(w http.ResponseWriter, _ *http.Request) {
	responderError(w, http.StatusNotFound, "ruta no encontrada")
}

func partesRuta(ruta, prefijo string) []string {
	resto := strings.Trim(strings.TrimPrefix(ruta, prefijo), "/")
	if resto == "" {
		return nil
	}
	return strings.Split(resto, "/")
}
