package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// DB representa la conexión a la base de datos
type DB struct {
	*sql.DB
}

// NewDB crea una nueva conexión a la base de datos SQLite
func NewDB(dbPath string) (*DB, error) {
	// Crear directorio si no existe
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("error creando directorio de BD: %w", err)
	}

	// Abrir conexión a SQLite
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo BD: %w", err)
	}

	// Verificar conexión
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando a BD: %w", err)
	}

	// Configurar conexión
	db.SetMaxOpenConns(1) // SQLite solo permite una conexión de escritura
	db.SetMaxIdleConns(1)

	log.Printf("✅ Base de datos conectada: %s", dbPath)
	return &DB{db}, nil
}

// Close cierra la conexión a la base de datos
func (db *DB) Close() error {
	return db.DB.Close()
}

// Ping verifica la conexión a la base de datos
func (db *DB) Ping() error {
	return db.DB.Ping()
}
