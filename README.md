# Vinilo Collection 🎵

Una aplicación web moderna para administrar y mostrar tu colección personal de vinilos, construida con Go, templ y Tailwind CSS.

## ✨ Características

- **Frontend Moderno**: Interfaz hermosa y minimalista con Tailwind CSS
- **Glassmorphism**: Efectos visuales avanzados en la vista de detalle
- **Búsqueda Inteligente**: Busca por artista, título o sello
- **Paginación**: Navegación eficiente por la colección
- **Modo Oscuro**: Soporte completo para tema oscuro
- **Responsive**: Diseño adaptativo para todos los dispositivos
- **SEO Optimizado**: Meta tags y estructura semántica
- **Base de Datos SQLite**: Almacenamiento local y portable

## 🛠️ Stack Tecnológico

- **Backend**: Go 1.24
- **Templates**: templ (componentes tipados)
- **CSS**: Tailwind CSS (CDN) con glassmorphism
- **Base de Datos**: SQLite con migraciones
- **Router**: Chi (ligero y rápido)
- **Migraciones**: Goose

## 🚀 Instalación

### Prerrequisitos

- Go 1.24 o superior

### 1. Clonar y configurar

```bash
git clone <tu-repo>
cd vinilo
```

### 2. Instalar dependencias

```bash
# Dependencias de Go
go mod tidy

# Dependencias de Node.js
npm install

# Herramientas de desarrollo
make install-deps
```

### 3. Configurar base de datos

```bash
# Aplicar migraciones
make migrate

# Poblar con datos de ejemplo
make seed
```

### 4. Compilar CSS

```bash
# Compilar Tailwind CSS
npm run build:css
```

### 5. Generar templates

```bash
# Generar templates templ
make generate-templates
```

## 🎯 Desarrollo

### Comandos principales

```bash
# Servidor de desarrollo
make dev

# Compilar binario
make build

# Aplicar migraciones
make migrate

# Poblar datos
make seed

# Limpiar archivos generados
make clean
```

### Desarrollo con hot reload

```bash
# Terminal 1: Servidor Go
make dev

# Terminal 2: Tailwind CSS en watch
npm run watch:css

# Terminal 3: Templates templ en watch
templ generate --watch
```

## 📁 Estructura del Proyecto

```
vinilo/
├── cmd/
│   ├── server/          # Servidor principal
│   └── seed/            # Comando para poblar datos
├── internal/
│   ├── database/        # Configuración de BD
│   ├── handlers/        # Handlers HTTP
│   ├── models/          # Modelos de datos
│   └── repository/      # Capa de acceso a datos
├── migrations/          # Migraciones SQL
├── web/
│   ├── input.css        # CSS de entrada para Tailwind
│   ├── static/          # Archivos estáticos
│   └── templates/       # Templates templ
├── data/                # Base de datos SQLite
├── Makefile             # Comandos de desarrollo
├── tailwind.config.js   # Configuración de Tailwind
└── package.json         # Dependencias de Node.js
```

## 🎨 Personalización

### Colores y temas

Los colores personalizados están definidos en `tailwind.config.js`:

```javascript
colors: {
  vinyl: {
    50: '#fef7f0',
    500: '#f2751f',
    600: '#e35d15',
    // ...
  }
}
```

### Glassmorphism

Los efectos glassmorphism están definidos en `web/input.css`:

```css
.glass-panel {
  @apply backdrop-blur-md bg-white/10 border border-white/20 rounded-xl shadow-xl;
}
```

## 📊 Modelo de Datos

### Campos del Record

- `id`: UUID único
- `titulo`: Título del álbum
- `artista`: Artista o banda
- `sello`: Sello discográfico
- `catalog_number`: Número de catálogo
- `anio`: Año de lanzamiento
- `formato`: Formato del vinilo (LP, EP, etc.)
- `generos`: Array de géneros musicales
- `estilos`: Array de estilos musicales
- `pais`: País de origen
- `tracklist`: Lista de canciones (JSON)
- `duracion_total`: Duración total
- `arte_url`: URL del artwork
- `condicion`: Estado del vinilo
- `notas`: Notas adicionales
- `created_at`: Fecha de creación
- `updated_at`: Fecha de actualización

## 🔧 Configuración Avanzada

### Variables de entorno

Copia `env.example` a `.env` y configura:

```bash
PORT=8080
DB_PATH=./data/vinilo.db
ENV=development
```

### Migraciones

Para crear una nueva migración:

```bash
make create-migration
# Ingresa el nombre de la migración
```

### Despliegue en Raspberry Pi

1. Compilar para ARM:
```bash
GOOS=linux GOARCH=arm GOARM=7 go build -o vinilo cmd/server/main.go
```

2. Transferir archivos:
```bash
scp vinilo pi@raspberrypi.local:/home/pi/
scp -r web/static pi@raspberrypi.local:/home/pi/
```

3. Ejecutar en Pi:
```bash
./vinilo
```

## 🎵 Agregar Nuevos Vinilos

### Método 1: Usando el seed

Edita `cmd/seed/main.go` y agrega nuevos records al array.

### Método 2: Directamente en la BD

```sql
INSERT INTO records (
  id, titulo, artista, sello, anio, formato, 
  generos, estilos, pais, tracklist, duracion_total, 
  arte_url, condicion, notas, created_at, updated_at
) VALUES (
  'uuid-here', 'Título', 'Artista', 'Sello', 2024, 'LP',
  '["Rock"]', '["Alternative"]', 'USA', '[]', '45:00',
  'https://example.com/artwork.jpg', 'Mint', 'Notas', 
  datetime('now'), datetime('now')
);
```

## 🐛 Troubleshooting

### Error: "templ: command not found"
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Error: "tailwindcss: command not found"
```bash
npm install -g tailwindcss
```

### Error: "goose: command not found"
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### CSS no se actualiza
```bash
# Regenerar CSS
npm run build:css

# O en modo watch
npm run watch:css
```

### Templates no se regeneran
```bash
# Regenerar templates
templ generate

# O en modo watch
templ generate --watch
```

## 📝 Licencia

MIT License - ver [LICENSE](LICENSE) para detalles.

## 🤝 Contribuir

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📞 Soporte

Si tienes problemas o preguntas:

1. Revisa la sección [Troubleshooting](#-troubleshooting)
2. Abre un issue en GitHub
3. Contacta al desarrollador

---

**¡Disfruta tu colección de vinilos! 🎵** 