-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS records (
    id TEXT PRIMARY KEY,
    titulo TEXT NOT NULL,
    artista TEXT NOT NULL,
    sello TEXT,
    catalog_number TEXT,
    anio INTEGER,
    formato TEXT,
    generos TEXT, -- JSON array como string
    estilos TEXT, -- JSON array como string
    pais TEXT,
    tracklist TEXT, -- JSON object como string
    duracion_total TEXT,
    arte_url TEXT,
    condicion TEXT,
    notas TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Índices para búsquedas eficientes
CREATE INDEX IF NOT EXISTS idx_records_artista ON records(artista);
CREATE INDEX IF NOT EXISTS idx_records_anio ON records(anio);
CREATE INDEX IF NOT EXISTS idx_records_sello ON records(sello);
CREATE INDEX IF NOT EXISTS idx_records_created_at ON records(created_at);

-- Trigger para actualizar updated_at automáticamente
CREATE TRIGGER IF NOT EXISTS update_records_updated_at 
    AFTER UPDATE ON records
    FOR EACH ROW
    BEGIN
        UPDATE records SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
    END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS records;
-- +goose StatementEnd 