package handlers

import (
	"net/http"

	"github.com/rodrwan/vinilo/web/templates"
)

// LandingHandler maneja la página principal
//
// Endpoint: GET /
//
// Funcionalidad:
// - Página de inicio pública del sitio web
// - Presenta la aplicación de gestión de colección de vinilos
// - Punto de entrada principal para usuarios visitantes
//
// Comportamiento:
// - Renderiza la página de landing con información general
// - Proporciona navegación hacia las funcionalidades principales
// - Sirve como página de bienvenida para nuevos usuarios
//
// Respuestas:
//   - 200: Página de landing renderizada correctamente
//
// Vista: templates.Landing
//
// Propósito:
// - Introducir la aplicación a los usuarios
// - Proporcionar acceso a la navegación principal
// - Mostrar información general sobre la colección de vinilos
func LandingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Renderizar la página de landing
		templates.Landing().Render(r.Context(), w)
	}
}
