package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Record representa un vinilo en la colección
type Record struct {
	ID            string         `json:"id" db:"id"`
	Titulo        string         `json:"titulo" db:"titulo"`
	Artista       string         `json:"artista" db:"artista"`
	Sello         sql.NullString `json:"sello" db:"sello"`
	CatalogNumber sql.NullString `json:"catalog_number" db:"catalog_number"`
	Anio          sql.NullInt32  `json:"anio" db:"anio"`
	Formato       sql.NullString `json:"formato" db:"formato"`
	Generos       sql.NullString `json:"generos" db:"generos"`
	Estilos       sql.NullString `json:"estilos" db:"estilos"`
	Pais          sql.NullString `json:"pais" db:"pais"`
	Tracklist     sql.NullString `json:"tracklist" db:"tracklist"`
	DuracionTotal sql.NullString `json:"duracion_total" db:"duracion_total"`
	ArteURL       sql.NullString `json:"arte_url" db:"arte_url"`
	Condicion     sql.NullString `json:"condicion" db:"condicion"`
	Notas         sql.NullString `json:"notas" db:"notas"`
	CreatedAt     time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" db:"updated_at"`
}

// RecordCreate representa los datos para crear un nuevo record
type RecordCreate struct {
	Titulo        string   `json:"titulo" validate:"required"`
	Artista       string   `json:"artista" validate:"required"`
	Sello         string   `json:"sello"`
	CatalogNumber string   `json:"catalog_number"`
	Anio          int      `json:"anio"`
	Formato       string   `json:"formato"`
	Generos       []string `json:"generos"`
	Estilos       []string `json:"estilos"`
	Pais          string   `json:"pais"`
	Tracklist     []Track  `json:"tracklist"`
	DuracionTotal string   `json:"duracion_total"`
	ArteURL       string   `json:"arte_url"`
	Condicion     string   `json:"condicion"`
	Notas         string   `json:"notas"`
}

// Track representa una canción en el tracklist
type Track struct {
	Numero   int    `json:"numero"`
	Titulo   string `json:"titulo"`
	Duracion string `json:"duracion"`
}

// RecordUpdate representa los datos para actualizar un record
type RecordUpdate struct {
	Titulo        *string  `json:"titulo"`
	Artista       *string  `json:"artista"`
	Sello         *string  `json:"sello"`
	CatalogNumber *string  `json:"catalog_number"`
	Anio          *int     `json:"anio"`
	Formato       *string  `json:"formato"`
	Generos       []string `json:"generos"`
	Estilos       []string `json:"estilos"`
	Pais          *string  `json:"pais"`
	Tracklist     []Track  `json:"tracklist"`
	DuracionTotal *string  `json:"duracion_total"`
	ArteURL       *string  `json:"arte_url"`
	Condicion     *string  `json:"condicion"`
	Notas         *string  `json:"notas"`
}

// NewRecord crea un nuevo record con ID generado
func NewRecord() *Record {
	return &Record{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// GetGenerosAsSlice convierte el string JSON de géneros a slice
func (r *Record) GetGenerosAsSlice() []string {
	if !r.Generos.Valid {
		return []string{}
	}

	var generos []string
	if err := json.Unmarshal([]byte(r.Generos.String), &generos); err != nil {
		return []string{}
	}
	return generos
}

// GetEstilosAsSlice convierte el string JSON de estilos a slice
func (r *Record) GetEstilosAsSlice() []string {
	if !r.Estilos.Valid {
		return []string{}
	}

	var estilos []string
	if err := json.Unmarshal([]byte(r.Estilos.String), &estilos); err != nil {
		return []string{}
	}
	return estilos
}

// GetTracklistAsSlice convierte el string JSON de tracklist a slice
func (r *Record) GetTracklistAsSlice() []Track {
	if !r.Tracklist.Valid {
		return []Track{}
	}

	var tracklist []Track
	if err := json.Unmarshal([]byte(r.Tracklist.String), &tracklist); err != nil {
		return []Track{}
	}
	return tracklist
}

// SetGeneros convierte el slice de géneros a JSON string
func (r *Record) SetGeneros(generos []string) {
	if len(generos) == 0 {
		r.Generos = sql.NullString{Valid: false}
		return
	}

	data, err := json.Marshal(generos)
	if err != nil {
		r.Generos = sql.NullString{Valid: false}
		return
	}

	r.Generos = sql.NullString{
		String: string(data),
		Valid:  true,
	}
}

// SetEstilos convierte el slice de estilos a JSON string
func (r *Record) SetEstilos(estilos []string) {
	if len(estilos) == 0 {
		r.Estilos = sql.NullString{Valid: false}
		return
	}

	data, err := json.Marshal(estilos)
	if err != nil {
		r.Estilos = sql.NullString{Valid: false}
		return
	}

	r.Estilos = sql.NullString{
		String: string(data),
		Valid:  true,
	}
}

// SetTracklist convierte el slice de tracks a JSON string
func (r *Record) SetTracklist(tracklist []Track) {
	if len(tracklist) == 0 {
		r.Tracklist = sql.NullString{Valid: false}
		return
	}

	data, err := json.Marshal(tracklist)
	if err != nil {
		r.Tracklist = sql.NullString{Valid: false}
		return
	}

	r.Tracklist = sql.NullString{
		String: string(data),
		Valid:  true,
	}
}

// GetDisplayTitle retorna el título para mostrar
func (r *Record) GetDisplayTitle() string {
	if r.Titulo != "" {
		return r.Titulo
	}
	return "Sin título"
}

// GetDisplayArtist retorna el artista para mostrar
func (r *Record) GetDisplayArtist() string {
	if r.Artista != "" {
		return r.Artista
	}
	return "Artista desconocido"
}

// GetYear retorna el año como string
func (r *Record) GetYear() string {
	if r.Anio.Valid {
		return string(rune(r.Anio.Int32))
	}
	return ""
}

// GetCondition retorna la condición o "No especificada"
func (r *Record) GetCondition() string {
	if r.Condicion.Valid {
		return r.Condicion.String
	}
	return "No especificada"
}

// GetArtworkURL retorna la URL del artwork o una imagen por defecto
func (r *Record) GetArtworkURL() string {
	if r.ArteURL.Valid && r.ArteURL.String != "" {
		return r.ArteURL.String
	}
	return "/static/images/default-vinyl.jpg"
}
