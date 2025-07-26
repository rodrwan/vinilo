package repository

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/rodrwan/vinilo/internal/database"
	"github.com/rodrwan/vinilo/internal/models"
)

// RecordRepository maneja las operaciones de base de datos para records
type RecordRepository struct {
	db *database.DB
}

// NewRecordRepository crea un nuevo repositorio de records
func NewRecordRepository(db *database.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

// Create crea un nuevo record en la base de datos
func (r *RecordRepository) Create(record *models.Record) error {
	query := `
		INSERT INTO records (
			id, titulo, artista, sello, catalog_number, anio, formato,
			generos, estilos, pais, tracklist, duracion_total, arte_url,
			condicion, notas, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query,
		record.ID,
		record.Titulo,
		record.Artista,
		record.Sello,
		record.CatalogNumber,
		record.Anio,
		record.Formato,
		record.Generos,
		record.Estilos,
		record.Pais,
		record.Tracklist,
		record.DuracionTotal,
		record.ArteURL,
		record.Condicion,
		record.Notas,
		record.CreatedAt,
		record.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creando record: %w", err)
	}

	log.Printf("✅ Record creado: %s - %s", record.Artista, record.Titulo)
	return nil
}

// GetByID obtiene un record por su ID
func (r *RecordRepository) GetByID(id string) (*models.Record, error) {
	query := `SELECT * FROM records WHERE id = ?`

	var record models.Record
	err := r.db.QueryRow(query, id).Scan(
		&record.ID,
		&record.Titulo,
		&record.Artista,
		&record.Sello,
		&record.CatalogNumber,
		&record.Anio,
		&record.Formato,
		&record.Generos,
		&record.Estilos,
		&record.Pais,
		&record.Tracklist,
		&record.DuracionTotal,
		&record.ArteURL,
		&record.Condicion,
		&record.Notas,
		&record.CreatedAt,
		&record.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record no encontrado: %s", id)
		}
		return nil, fmt.Errorf("error obteniendo record: %w", err)
	}

	return &record, nil
}

// GetAll obtiene todos los records con paginación
func (r *RecordRepository) GetAll(limit, offset int) ([]*models.Record, error) {
	query := `
		SELECT * FROM records 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo records: %w", err)
	}
	defer rows.Close()

	var records []*models.Record
	for rows.Next() {
		var record models.Record
		err := rows.Scan(
			&record.ID,
			&record.Titulo,
			&record.Artista,
			&record.Sello,
			&record.CatalogNumber,
			&record.Anio,
			&record.Formato,
			&record.Generos,
			&record.Estilos,
			&record.Pais,
			&record.Tracklist,
			&record.DuracionTotal,
			&record.ArteURL,
			&record.Condicion,
			&record.Notas,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando record: %w", err)
		}
		records = append(records, &record)
	}

	return records, nil
}

// Search busca records por término
func (r *RecordRepository) Search(term string, limit, offset int) ([]*models.Record, error) {
	query := `
		SELECT * FROM records 
		WHERE titulo LIKE ? OR artista LIKE ? OR sello LIKE ?
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`

	searchTerm := "%" + strings.ToLower(term) + "%"
	rows, err := r.db.Query(query, searchTerm, searchTerm, searchTerm, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error buscando records: %w", err)
	}
	defer rows.Close()

	var records []*models.Record
	for rows.Next() {
		var record models.Record
		err := rows.Scan(
			&record.ID,
			&record.Titulo,
			&record.Artista,
			&record.Sello,
			&record.CatalogNumber,
			&record.Anio,
			&record.Formato,
			&record.Generos,
			&record.Estilos,
			&record.Pais,
			&record.Tracklist,
			&record.DuracionTotal,
			&record.ArteURL,
			&record.Condicion,
			&record.Notas,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando record: %w", err)
		}
		records = append(records, &record)
	}

	return records, nil
}

// Count obtiene el total de records
func (r *RecordRepository) Count() (int, error) {
	query := `SELECT COUNT(*) FROM records`

	var count int
	err := r.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error contando records: %w", err)
	}

	return count, nil
}

// CountSearch obtiene el total de records que coinciden con la búsqueda
func (r *RecordRepository) CountSearch(term string) (int, error) {
	query := `
		SELECT COUNT(*) FROM records 
		WHERE titulo LIKE ? OR artista LIKE ? OR sello LIKE ?
	`

	searchTerm := "%" + strings.ToLower(term) + "%"
	var count int
	err := r.db.QueryRow(query, searchTerm, searchTerm, searchTerm).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error contando búsqueda: %w", err)
	}

	return count, nil
}

// Update actualiza un record existente
func (r *RecordRepository) Update(record *models.Record) error {
	query := `
		UPDATE records SET 
			titulo = ?, artista = ?, sello = ?, catalog_number = ?, 
			anio = ?, formato = ?, generos = ?, estilos = ?, pais = ?,
			tracklist = ?, duracion_total = ?, arte_url = ?, 
			condicion = ?, notas = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query,
		record.Titulo,
		record.Artista,
		record.Sello,
		record.CatalogNumber,
		record.Anio,
		record.Formato,
		record.Generos,
		record.Estilos,
		record.Pais,
		record.Tracklist,
		record.DuracionTotal,
		record.ArteURL,
		record.Condicion,
		record.Notas,
		record.UpdatedAt,
		record.ID,
	)

	if err != nil {
		return fmt.Errorf("error actualizando record: %w", err)
	}

	log.Printf("✅ Record actualizado: %s - %s", record.Artista, record.Titulo)
	return nil
}

// Delete elimina un record por su ID
func (r *RecordRepository) Delete(id string) error {
	query := `DELETE FROM records WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error eliminando record: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error obteniendo filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("record no encontrado: %s", id)
	}

	log.Printf("✅ Record eliminado: %s", id)
	return nil
}

// GetByArtist obtiene records por artista
func (r *RecordRepository) GetByArtist(artist string, limit, offset int) ([]*models.Record, error) {
	query := `
		SELECT * FROM records 
		WHERE artista LIKE ? 
		ORDER BY anio DESC, titulo ASC 
		LIMIT ? OFFSET ?
	`

	searchTerm := "%" + strings.ToLower(artist) + "%"
	rows, err := r.db.Query(query, searchTerm, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo records por artista: %w", err)
	}
	defer rows.Close()

	var records []*models.Record
	for rows.Next() {
		var record models.Record
		err := rows.Scan(
			&record.ID,
			&record.Titulo,
			&record.Artista,
			&record.Sello,
			&record.CatalogNumber,
			&record.Anio,
			&record.Formato,
			&record.Generos,
			&record.Estilos,
			&record.Pais,
			&record.Tracklist,
			&record.DuracionTotal,
			&record.ArteURL,
			&record.Condicion,
			&record.Notas,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando record: %w", err)
		}
		records = append(records, &record)
	}

	return records, nil
}
