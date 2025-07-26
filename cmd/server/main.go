package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/rodrwan/vinilo/internal/database"
	"github.com/rodrwan/vinilo/internal/handlers"
	"github.com/rodrwan/vinilo/internal/repository"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// Configuración
	port := getEnv("PORT", "8080")
	dbPath := getEnv("DB_PATH", "./data/vinilo.db")

	// Inicializar base de datos
	db, err := database.NewDB(dbPath)
	if err != nil {
		log.Fatalf("Error inicializando base de datos: %v", err)
	}
	defer db.Close()

	// Inicializar repositorio
	recordRepo := repository.NewRecordRepository(db)

	// Inicializar handlers
	landingHandler := handlers.LandingHandler()
	recordsHandler := handlers.NewRecordsHandler(recordRepo)
	adminHandler := handlers.NewAdminHandler(recordRepo)
	// Configurar router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Compress(5))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Servir archivos estáticos (solo para favicon y otros assets)
	fileServer := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Rutas
	r.Get("/", landingHandler)
	r.Get("/records", recordsHandler.ListHandler())
	r.Get("/records/{id}", recordsHandler.DetailHandler())

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Rutas para el admin
	r.Get("/admin", adminHandler.HomeHandler())
	r.Get("/admin/records", adminHandler.ListHandler())
	r.Get("/admin/records/new", adminHandler.NewRecordHandler())
	r.Post("/admin/records", adminHandler.CreateRecordHandler())
	r.Get("/admin/records/{id}", adminHandler.DetailHandler())

	// Configurar servidor
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Canal para señales de terminación
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor en goroutine
	go func() {
		log.Printf("🚀 Servidor iniciado en http://localhost:%s", port)
		log.Printf("📊 Health check: http://localhost:%s/health", port)
		log.Printf("🗄️ Base de datos: %s", dbPath)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error iniciando servidor: %v", err)
		}
	}()

	// Esperar señal de terminación
	<-done
	log.Println("🛑 Cerrando servidor...")

	// Shutdown graceful
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error cerrando servidor: %v", err)
	}

	log.Println("✅ Servidor cerrado correctamente")
}

// getEnv obtiene una variable de entorno o retorna un valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
