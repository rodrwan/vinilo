package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/rodrwan/vinilo/internal/models"
	"github.com/rodrwan/vinilo/internal/repository"
	"github.com/rodrwan/vinilo/web/templates"
)

// AdminHandler maneja las operaciones administrativas de records
// Proporciona endpoints para gestionar la colección de discos de vinilo
// desde una interfaz administrativa con funcionalidades CRUD completas.
type AdminHandler struct {
	repo *repository.RecordRepository
}

// NewAdminHandler crea un nuevo handler administrativo
// Parámetros:
//   - repo: Repositorio de records para operaciones de base de datos
//
// Retorna: Una instancia configurada de AdminHandler
func NewAdminHandler(repo *repository.RecordRepository) *AdminHandler {
	return &AdminHandler{repo: repo}
}

// ListHandler maneja el listado administrativo de records
//
// Endpoint: GET /admin/records
//
// Funcionalidad:
// - Muestra una lista paginada de todos los records en la base de datos
// - Soporta búsqueda por texto en título y artista
// - Implementa paginación con 20 registros por página
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
//
// Respuestas:
//   - 200: Lista de records renderizada correctamente
//   - 500: Error interno del servidor al obtener datos
//
// Vista: templates.RecordsList
func (h *AdminHandler) ListHandler() http.HandlerFunc {
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
		limit := 20
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

		// Renderizar la página administrativa
		component := templates.RecordsList(records, total, page, searchTerm)
		templ.Handler(component).ServeHTTP(w, r)
	}
}

// DetailHandler maneja la vista administrativa de detalle de un record
//
// Endpoint: GET /admin/records/{id}
//
// Funcionalidad:
// - Muestra información detallada de un record específico
// - Permite ver todos los campos del record seleccionado
// - Proporciona contexto para futuras operaciones de edición
//
// Parámetros de URL:
//   - id: Identificador único del record (requerido)
//
// Comportamiento:
// - Extrae el ID del record de la URL path
// - Valida que el ID esté presente en la URL
// - Busca el record en la base de datos por ID
// - Renderiza la vista de detalle con información completa
//
// Respuestas:
//   - 200: Detalle del record renderizado correctamente
//   - 400: ID de record no proporcionado en la URL
//   - 404: Record no encontrado en la base de datos
//   - 500: Error interno del servidor
//
// Vista: templates.RecordDetail
func (h *AdminHandler) DetailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extraer ID de la URL
		pathParts := strings.Split(r.URL.Path, "/")
		if len(pathParts) < 4 {
			http.Error(w, "ID de record requerido", http.StatusBadRequest)
			return
		}

		recordID := pathParts[3]

		// Obtener record
		record, err := h.repo.GetByID(recordID)
		if err != nil {
			http.Error(w, "Record no encontrado", http.StatusNotFound)
			return
		}

		// Renderizar la página de detalle administrativa
		component := templates.RecordDetail(record)
		templ.Handler(component).ServeHTTP(w, r)
	}
}

// HomeHandler maneja la página principal del admin
//
// Endpoint: GET /admin
//
// Funcionalidad:
// - Dashboard principal del panel administrativo
// - Muestra estadísticas generales de la colección
// - Lista los records más recientes agregados
// - Proporciona una vista general del estado del sistema
//
// Comportamiento:
// - Obtiene el conteo total de records en la base de datos
// - Recupera los 5 records más recientes (ordenados por fecha de creación)
// - Renderiza el dashboard con estadísticas y lista de recientes
//
// Respuestas:
//   - 200: Dashboard administrativo renderizado correctamente
//   - 500: Error interno del servidor al obtener estadísticas
//
// Vista: templates.AdminDashboard
//
// Datos mostrados:
//   - Total de records en la colección
//   - Lista de 5 records más recientes
func (h *AdminHandler) HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener estadísticas
		totalRecords, err := h.repo.Count()
		if err != nil {
			http.Error(w, "Error obteniendo estadísticas", http.StatusInternalServerError)
			return
		}

		// Obtener últimas 5 canciones registradas
		recentRecords, err := h.repo.GetAll(5, 0)
		if err != nil {
			http.Error(w, "Error obteniendo records recientes", http.StatusInternalServerError)
			return
		}

		// Renderizar dashboard
		component := templates.AdminDashboard(totalRecords, recentRecords)
		templ.Handler(component).ServeHTTP(w, r)
	}
}

// CreateRecordHandler maneja la creación de nuevos records
//
// Endpoint: POST /admin/records/new
//
// Funcionalidad:
// - Procesa el formulario de creación de nuevos records
// - Valida los datos requeridos del formulario
// - Crea un nuevo record en la base de datos
// - Redirige al usuario después de la creación exitosa
//
// Método: POST
//
// Parámetros del Formulario:
//   - titulo: Título del disco (requerido)
//   - artista: Nombre del artista (requerido)
//   - anio: Año de lanzamiento (opcional)
//   - generos: Géneros musicales (opcional)
//
// Validaciones:
// - Título y artista son campos obligatorios
// - El año debe ser un número válido si se proporciona
// - Los géneros se almacenan como texto libre
//
// Comportamiento:
// - Parsea el formulario enviado
// - Valida que los campos requeridos estén presentes
// - Convierte el año a entero si se proporciona
// - Crea un nuevo record con los datos proporcionados
// - Guarda el record en la base de datos
// - Redirige al usuario a la lista de records
//
// Respuestas:
//   - 302: Redirección a /admin/records después de creación exitosa
//   - 400: Datos del formulario inválidos o campos requeridos faltantes
//   - 405: Método HTTP no permitido (solo POST)
//   - 500: Error interno del servidor al crear el record
//
// Redirección: /admin/records
func (h *AdminHandler) CreateRecordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Parsear formulario
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parseando formulario", http.StatusBadRequest)
			return
		}

		// Obtener datos del formulario
		titulo := r.FormValue("titulo")
		artista := r.FormValue("artista")
		anioStr := r.FormValue("anio")
		generos := r.FormValue("generos")

		if titulo == "" || artista == "" {
			http.Error(w, "Título y artista son requeridos", http.StatusBadRequest)
			return
		}

		// Parsear año
		var anio sql.NullInt32
		if anioStr != "" {
			if y, err := strconv.Atoi(anioStr); err == nil {
				anio = sql.NullInt32{Int32: int32(y), Valid: true}
			}
		}

		// Crear nuevo record
		record := models.NewRecord()
		record.Titulo = titulo
		record.Artista = artista
		record.Anio = anio
		if generos != "" {
			record.Generos = sql.NullString{String: generos, Valid: true}
		}

		// Guardar en base de datos
		if err := h.repo.Create(record); err != nil {
			http.Error(w, "Error creando record", http.StatusInternalServerError)
			return
		}

		// Redirigir a la lista de records
		http.Redirect(w, r, "/admin/records", http.StatusSeeOther)
	}
}

// NewRecordHandler maneja la vista del formulario para crear un nuevo record
//
// Endpoint: GET /admin/records/new
//
// Funcionalidad:
// - Muestra el formulario para crear un nuevo record
// - Proporciona una interfaz de usuario para ingresar datos del disco
// - Permite al administrador agregar nuevos records a la colección
//
// Comportamiento:
// - Renderiza el formulario de creación de records
// - El formulario incluye campos para título, artista, año y géneros
// - Al enviar el formulario, se procesa en CreateRecordHandler
//
// Respuestas:
//   - 200: Formulario de creación renderizado correctamente
//
// Vista: templates.NewRecordForm
//
// Campos del formulario:
//   - Título (texto, requerido)
//   - Artista (texto, requerido)
//   - Año (número, opcional)
//   - Géneros (texto, opcional)
func (h *AdminHandler) NewRecordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component := templates.NewRecordForm()
		templ.Handler(component).ServeHTTP(w, r)
	}
}
