package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sigi/utils"
)

type errorResponse struct {
	Error string `json:"error"`
}

func responderJSON(w http.ResponseWriter, estado int, cuerpo any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(estado)
	if err := json.NewEncoder(w).Encode(cuerpo); err != nil {
		return
	}
}

func responderError(w http.ResponseWriter, estado int, mensaje string) {
	responderJSON(w, estado, errorResponse{Error: mensaje})
}

func decodificarJSON(w http.ResponseWriter, r *http.Request, destino any) bool {
	decodificador := json.NewDecoder(r.Body)
	decodificador.DisallowUnknownFields()
	if err := decodificador.Decode(destino); err != nil {
		responderError(w, http.StatusBadRequest, fmt.Sprintf("JSON inválido: %v", err))
		return false
	}

	if err := decodificador.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		responderError(w, http.StatusBadRequest, "el cuerpo debe contener un único objeto JSON")
		return false
	}

	return true
}

func responderErrorServicio(w http.ResponseWriter, err error) {
	mensaje := err.Error()
	mensajeNormalizado := strings.ToLower(mensaje)

	switch {
	case errors.Is(err, utils.ErrRegistroNoEncontrado):
		responderError(w, http.StatusNotFound, mensaje)
	case errors.Is(err, utils.ErrOperacionNoPermitida),
		strings.Contains(mensajeNormalizado, "no se puede"),
		strings.Contains(mensajeNormalizado, "transicion"),
		strings.Contains(mensajeNormalizado, "no está disponible"):
		responderError(w, http.StatusConflict, mensaje)
	default:
		responderError(w, http.StatusBadRequest, mensaje)
	}
}

func responderMetodoNoPermitido(w http.ResponseWriter, permitidos string) {
	w.Header().Set("Allow", permitidos)
	responderError(w, http.StatusMethodNotAllowed, "método HTTP no permitido")
}
