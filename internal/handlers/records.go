package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/rodrwan/vinilo/internal/models"
	"github.com/rodrwan/vinilo/internal/repository"
	"github.com/rodrwan/vinilo/web/templates"
)

// RecordsHandler maneja las operaciones de records
// Proporciona endpoints públicos para visualizar la colección de discos
// de vinilo con funcionalidades de listado, búsqueda y detalle.
type RecordsHandler struct {
	repo *repository.RecordRepository
}

// NewRecordsHandler crea un nuevo handler de records
// Parámetros:
//   - repo: Repositorio de records para operaciones de base de datos
//
// Retorna: Una instancia configurada de RecordsHandler
func NewRecordsHandler(repo *repository.RecordRepository) *RecordsHandler {
	return &RecordsHandler{repo: repo}
}

// ListHandler maneja el listado de records con búsqueda y paginación
//
// Endpoint: GET /records
//
// Funcionalidad:
// - Muestra una lista paginada de todos los records disponibles
// - Soporta búsqueda por texto en título y artista
// - Implementa paginación con 12 registros por página (optimizado para vista pública)
// - Permite navegación entre páginas
//
// Parámetros de Query:
//   - page: Número de página (opcional, default: 1)
//   - search: Término de búsqueda (opcional)
//
// Comportamiento:
// - Si se proporciona search, filtra records que contengan el término
// - Si no hay search, muestra todos los records
// - Calcula automáticamente el total de registros para paginación
// - Renderiza la vista RecordsList con datos paginados
// - Utiliza un límite de 12 registros por página (vs 20 en admin)
//
// Respuestas:
//   - 200: Lista de records renderizada correctamente
//   - 500: Error interno del servidor al obtener datos
//
// Vista: templates.RecordsList
//
// Diferencias con AdminHandler.ListHandler:
//   - Límite de 12 registros por página (vs 20 en admin)
//   - Vista pública sin funcionalidades administrativas
//   - Mismo template pero con contexto público
func (h *RecordsHandler) ListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener parámetros de query
		pageStr := r.URL.Query().Get("page")
		searchTerm := strings.TrimSpace(r.URL.Query().Get("search"))

		// Parsear página
		page := 1
		if pageStr != "" {
			if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
				page = p
			}
		}

		// Calcular offset
		limit := 12
		offset := (page - 1) * limit

		var records []*models.Record
		var total int
		var err error

		// Buscar o obtener todos
		if searchTerm != "" {
			records, err = h.repo.Search(searchTerm, limit, offset)
			if err != nil {
				http.Error(w, "Error buscando records", http.StatusInternalServerError)
				return
			}
			total, err = h.repo.CountSearch(searchTerm)
		} else {
			records, err = h.repo.GetAll(limit, offset)
			if err != nil {
				http.Error(w, "Error obteniendo records", http.StatusInternalServerError)
				return
			}
			total, err = h.repo.Count()
		}

		if err != nil {
			http.Error(w, "Error contando records", http.StatusInternalServerError)
			return
		}

		// Renderizar la página
		component := templates.RecordsList(records, total, page, searchTerm)
		templ.Handler(component).ServeHTTP(w, r)
	}
}

// DetailHandler maneja la vista de detalle de un record
//
// Endpoint: GET /records/{id}
//
// Funcionalidad:
// - Muestra información detallada de un record específico
// - Permite ver todos los campos del record seleccionado
// - Proporciona vista pública de los detalles del disco
//
// Parámetros de URL:
//   - id: Identificador único del record (requerido)
//
// Comportamiento:
// - Extrae el ID del record de la URL path
// - Valida que el ID esté presente en la URL
// - Busca el record en la base de datos por ID
// - Renderiza la vista de detalle con información completa
// - Proporciona vista pública sin funcionalidades de edición
//
// Respuestas:
//   - 200: Detalle del record renderizado correctamente
//   - 400: ID de record no proporcionado en la URL
//   - 404: Record no encontrado en la base de datos
//   - 500: Error interno del servidor
//
// Vista: templates.RecordDetail
//
// Diferencias con AdminHandler.DetailHandler:
//   - Vista pública sin funcionalidades administrativas
//   - Mismo template pero con contexto público
//   - Sin opciones de edición o eliminación
func (h *RecordsHandler) DetailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraer ID de la URL
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 3 {
			http.Error(w, "ID de record requerido", http.StatusBadRequest)
			return
		}

		recordID := pathParts[2]

		// Obtener record
		record, err := h.repo.GetByID(recordID)
		if err != nil {
			http.Error(w, "Record no encontrado", http.StatusNotFound)
			return
		}

		// Renderizar la página de detalle
		component := templates.RecordDetail(record)
		templ.Handler(component).ServeHTTP(w, r)
	}
}
