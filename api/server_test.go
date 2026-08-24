package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"sigi/api"
	"sigi/app"
)

func TestFlujoHTTPCompleto(t *testing.T) {
	sistema := app.NuevoSistema()
	servidor := api.NuevoServidor(
		sistema.ProveedorService,
		sistema.OrdenService,
		sistema.TransporteService,
		sistema.ImportacionService,
		sistema.InventarioService,
		sistema.ReporteService,
	).Handler()

	respuesta := solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", `{
        "empresa":"ACME Importaciones", "pais":"China", "contacto":"Li Wei",
        "telefono":"+86 1234567", "correo":"li@acme.example"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	proveedor := cuerpoJSON(t, respuesta)
	debeSer(t, proveedor, "codigo", "PRV-0001")

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/proveedores/PRV-0001", "")
	debeTenerEstado(t, respuesta, http.StatusOK)
	debeSer(t, cuerpoJSON(t, respuesta), "empresa", "ACME Importaciones")

	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/proveedores/PRV-0001/estado", `{"activo":false}`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	debeSer(t, cuerpoJSON(t, respuesta), "estado", "Inactivo")

	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/proveedores/PRV-0001/estado", `{"activo":true}`)
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/proveedores", "")
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/transportes", `{
        "tipo":"Marítimo", "empresa":"Ocean Cargo", "pais":"Ecuador", "contacto":"Ana Paz",
        "telefono":"+593 2222222", "correo":"ana@ocean.example"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	debeSer(t, cuerpoJSON(t, respuesta), "codigo", "TRN-0001")

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/transportes", "")
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/ordenes", `{"codigo_proveedor":"PRV-0001"}`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	debeSer(t, cuerpoJSON(t, respuesta), "codigo", "ORD-0001")

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/ordenes/ORD-0001/productos", `{
        "nombre":"Computadora portátil", "cantidad":3, "precio_unitario":850.50
    }`)
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/ordenes/ORD-0001/confirmacion", `{"confirmada":true}`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	debeSer(t, cuerpoJSON(t, respuesta), "estado", "Confirmada")

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/ordenes", "")
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/importaciones", `{
        "codigo_orden":"ORD-0001", "codigo_transporte":"TRN-0001",
        "ciudad_origen":"Shenzhen", "ciudad_destino":"Guayaquil"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	debeSer(t, cuerpoJSON(t, respuesta), "estado", "En preparación")
	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/reportes/general", "")
	debeTenerEstado(t, respuesta, http.StatusOK)
	if cuerpoJSON(t, respuesta)["importaciones_en_preparacion"] != float64(1) {
		t.Fatal("el reporte no contabilizó la importación en preparación")
	}

	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/importaciones/IMP-0001/tracking", `{"estado":"En tránsito"}`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/importaciones/IMP-0001/tracking", `{"estado":"En aduana"}`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/importaciones/IMP-0001/tracking", `{
        "estado":"Llegó a bodega", "ubicacion":"Bodega A-01"
    }`)
	debeTenerEstado(t, respuesta, http.StatusOK)

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/inventario", "")
	debeTenerEstado(t, respuesta, http.StatusOK)
	var inventarios []map[string]any
	if err := json.Unmarshal(respuesta.Body.Bytes(), &inventarios); err != nil {
		t.Fatalf("respuesta de inventario no es JSON válido: %v", err)
	}
	if len(inventarios) != 1 {
		t.Fatalf("se esperó un registro de inventario, se obtuvieron %d", len(inventarios))
	}

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/reportes/general", "")
	debeTenerEstado(t, respuesta, http.StatusOK)
	reporte := cuerpoJSON(t, respuesta)
	if reporte["total_proveedores"] != float64(1) || reporte["total_unidades_inventario"] != float64(3) {
		t.Fatalf("reporte inesperado: %#v", reporte)
	}
}

func TestErroresJSONUniformes(t *testing.T) {
	sistema := app.NuevoSistema()
	servidor := api.NuevoServidor(
		sistema.ProveedorService,
		sistema.OrdenService,
		sistema.TransporteService,
		sistema.ImportacionService,
		sistema.InventarioService,
		sistema.ReporteService,
	).Handler()

	respuesta := solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", `{"empresa":""}`)
	debeTenerEstado(t, respuesta, http.StatusBadRequest)
	if respuesta.Header().Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("Content-Type inesperado: %q", respuesta.Header().Get("Content-Type"))
	}
	if cuerpoJSON(t, respuesta)["error"] == "" {
		t.Fatal("la respuesta de error no contiene el campo error")
	}

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/proveedores/PRV-9999", "")
	debeTenerEstado(t, respuesta, http.StatusNotFound)

	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/no-existe", "")
	debeTenerEstado(t, respuesta, http.StatusNotFound)
	if cuerpoJSON(t, respuesta)["error"] != "ruta no encontrada" {
		t.Fatal("la ruta inexistente no devolvió el error JSON esperado")
	}

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", "{invalid")
	debeTenerEstado(t, respuesta, http.StatusBadRequest)
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", "")
	debeTenerEstado(t, respuesta, http.StatusBadRequest)
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", `{"empresa":"X","campo_extra":true}`)
	debeTenerEstado(t, respuesta, http.StatusBadRequest)

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores/PRV-9999", "{}")
	debeTenerEstado(t, respuesta, http.StatusMethodNotAllowed)
	respuesta = solicitud(t, servidor, http.MethodDelete, "/api/v1/transportes", "")
	debeTenerEstado(t, respuesta, http.StatusMethodNotAllowed)
}

func TestReglasDeImportacionYValidaciones(t *testing.T) {
	sistema := app.NuevoSistema()
	servidor := api.NuevoServidor(
		sistema.ProveedorService,
		sistema.OrdenService,
		sistema.TransporteService,
		sistema.ImportacionService,
		sistema.InventarioService,
		sistema.ReporteService,
	).Handler()

	respuesta := solicitud(t, servidor, http.MethodPost, "/api/v1/proveedores", `{
        "empresa":"Proveedor de prueba", "pais":"Ecuador", "contacto":"Ana",
        "telefono":"+593 2222222", "correo":"ana@example.com"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/transportes", `{
        "tipo":"Terrestre", "empresa":"Transporte de prueba", "pais":"Ecuador",
        "contacto":"Luis", "telefono":"+593 2333333", "correo":"luis@example.com"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)

	// Una orden pendiente no puede iniciar una importación.
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/ordenes", `{"codigo_proveedor":"PRV-0001"}`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/importaciones", `{
        "codigo_orden":"ORD-0001", "codigo_transporte":"TRN-0001",
        "ciudad_origen":"Quito", "ciudad_destino":"Guayaquil"
    }`)
	debeTenerEstado(t, respuesta, http.StatusConflict)

	// Se confirma una orden y se desactiva el transporte para validar esa regla.
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/ordenes/ORD-0001/productos", `{
        "nombre":"Carga", "cantidad":1, "precio_unitario":10
    }`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/ordenes/ORD-0001/confirmacion", `{"confirmada":true}`)
	debeTenerEstado(t, respuesta, http.StatusOK)
	if err := sistema.TransporteService.Desactivar("TRN-0001"); err != nil {
		t.Fatalf("no se pudo preparar el transporte inactivo: %v", err)
	}
	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/importaciones", `{
        "codigo_orden":"ORD-0001", "codigo_transporte":"TRN-0001",
        "ciudad_origen":"Quito", "ciudad_destino":"Guayaquil"
    }`)
	debeTenerEstado(t, respuesta, http.StatusConflict)
	if err := sistema.TransporteService.Activar("TRN-0001"); err != nil {
		t.Fatalf("no se pudo reactivar el transporte: %v", err)
	}

	respuesta = solicitud(t, servidor, http.MethodPost, "/api/v1/importaciones", `{
        "codigo_orden":"ORD-0001", "codigo_transporte":"TRN-0001",
        "ciudad_origen":"Quito", "ciudad_destino":"Guayaquil"
    }`)
	debeTenerEstado(t, respuesta, http.StatusCreated)
	// No se permite saltar directamente de preparación a aduana.
	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/importaciones/IMP-0001/tracking", `{"estado":"En aduana"}`)
	debeTenerEstado(t, respuesta, http.StatusConflict)
	respuesta = solicitud(t, servidor, http.MethodGet, "/api/v1/ordenes/ORD-9999", "")
	debeTenerEstado(t, respuesta, http.StatusNotFound)
	respuesta = solicitud(t, servidor, http.MethodPatch, "/api/v1/importaciones/IMP-9999/tracking", `{"estado":"En tránsito"}`)
	debeTenerEstado(t, respuesta, http.StatusNotFound)
}

func solicitud(t *testing.T, handler http.Handler, metodo, ruta, cuerpo string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(metodo, ruta, bytes.NewBufferString(cuerpo))
	if cuerpo != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	respuesta := httptest.NewRecorder()
	handler.ServeHTTP(respuesta, req)
	return respuesta
}

func debeTenerEstado(t *testing.T, respuesta *httptest.ResponseRecorder, esperado int) {
	t.Helper()
	if respuesta.Code != esperado {
		datos, _ := io.ReadAll(respuesta.Body)
		t.Fatalf("estado HTTP: se obtuvo %d, se esperaba %d. Cuerpo: %s", respuesta.Code, esperado, datos)
	}
}

func cuerpoJSON(t *testing.T, respuesta *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	var cuerpo map[string]any
	if err := json.Unmarshal(respuesta.Body.Bytes(), &cuerpo); err != nil {
		t.Fatalf("respuesta no es un objeto JSON válido: %v. Cuerpo: %s", err, respuesta.Body.String())
	}
	return cuerpo
}

func debeSer(t *testing.T, cuerpo map[string]any, campo, esperado string) {
	t.Helper()
	if cuerpo[campo] != esperado {
		t.Fatalf("campo %s: se obtuvo %#v, se esperaba %q", campo, cuerpo[campo], esperado)
	}
}
