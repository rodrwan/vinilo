.PHONY: dev build migrate clean install-deps generate-templates

# Variables
BINARY_NAME=vinilo
BUILD_DIR=build
MIGRATIONS_DIR=migrations
DB_PATH=./data/vinilo.db

# Instalar dependencias
install-deps:
	go mod download
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Generar templates templ
generate-templates:
	templ generate

# Desarrollo con hot reload
dev: generate-templates
	@echo "🚀 Iniciando servidor de desarrollo..."
	@echo "📝 Templates generados"
	@echo "🌐 Servidor en http://localhost:8080"
	@echo ""
	@echo "Comandos útiles:"
	@echo "  make generate-templates  - Regenerar templates"
	@echo "  make migrate            - Aplicar migraciones"
	@echo ""
	go run cmd/server/main.go

# Build del binario
build: generate-templates
	@echo "🔨 Compilando binario..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/server/main.go
	@echo "✅ Binario generado en $(BUILD_DIR)/$(BINARY_NAME)"

# Aplicar migraciones
migrate:
	@echo "🗄️ Aplicando migraciones..."
	mkdir -p data
	goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_PATH) up
	@echo "✅ Migraciones aplicadas"

# Crear migración
create-migration:
	@read -p "Nombre de la migración: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql

# Limpiar archivos generados
clean:
	rm -rf $(BUILD_DIR)
	rm -f templ_*.go

# Seed inicial de datos
seed:
	@echo "🌱 Aplicando seed inicial..."
	go run cmd/seed/main.go

# Setup completo del proyecto
setup: install-deps migrate seed
	@echo "✅ Proyecto configurado completamente"

# Ejecutar tests
test:
	go test ./...

# Lint
lint:
	golangci-lint run

# Help
help:
	@echo "Comandos disponibles:"
	@echo "  make setup              - Configuración completa del proyecto"
	@echo "  make dev               - Servidor de desarrollo"
	@echo "  make build             - Compilar binario"
	@echo "  make migrate           - Aplicar migraciones"
	@echo "  make create-migration  - Crear nueva migración"
	@echo "  make seed              - Aplicar datos iniciales"
	@echo "  make clean             - Limpiar archivos generados"
	@echo "  make test              - Ejecutar tests"
	@echo "  make lint              - Ejecutar linter" 