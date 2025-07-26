package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/rodrwan/vinilo/internal/database"
	"github.com/rodrwan/vinilo/internal/models"
	"github.com/rodrwan/vinilo/internal/repository"
)

func main() {
	log.Println("🌱 Iniciando seed de datos...")

	// Conectar a la base de datos
	db, err := database.NewDB("./data/vinilo.db")
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}
	defer db.Close()

	// Crear repositorio
	repo := repository.NewRecordRepository(db)

	// Datos de ejemplo
	records := []*models.Record{
		{
			ID:            "1",
			Titulo:        "The Dark Side of the Moon",
			Artista:       "Pink Floyd",
			Sello:         sql.NullString{String: "Harvest", Valid: true},
			CatalogNumber: sql.NullString{String: "SHVL 804", Valid: true},
			Anio:          sql.NullInt32{Int32: 1973, Valid: true},
			Formato:       sql.NullString{String: "LP", Valid: true},
			Pais:          sql.NullString{String: "UK", Valid: true},
			Condicion:     sql.NullString{String: "Mint", Valid: true},
			DuracionTotal: sql.NullString{String: "42:49", Valid: true},
			ArteURL:       sql.NullString{String: "https://upload.wikimedia.org/wikipedia/en/3/3b/Dark_Side_of_the_Moon.png", Valid: true},
			Notas:         sql.NullString{String: "Uno de los álbumes más influyentes de la historia del rock", Valid: true},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "2",
			Titulo:        "Abbey Road",
			Artista:       "The Beatles",
			Sello:         sql.NullString{String: "Apple", Valid: true},
			CatalogNumber: sql.NullString{String: "PCS 7088", Valid: true},
			Anio:          sql.NullInt32{Int32: 1969, Valid: true},
			Formato:       sql.NullString{String: "LP", Valid: true},
			Pais:          sql.NullString{String: "UK", Valid: true},
			Condicion:     sql.NullString{String: "Near Mint", Valid: true},
			DuracionTotal: sql.NullString{String: "47:23", Valid: true},
			ArteURL:       sql.NullString{String: "https://upload.wikimedia.org/wikipedia/en/4/42/Beatles_-_Abbey_Road.jpg", Valid: true},
			Notas:         sql.NullString{String: "El último álbum grabado por los Beatles", Valid: true},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "3",
			Titulo:        "Kind of Blue",
			Artista:       "Miles Davis",
			Sello:         sql.NullString{String: "Columbia", Valid: true},
			CatalogNumber: sql.NullString{String: "CL 1355", Valid: true},
			Anio:          sql.NullInt32{Int32: 1959, Valid: true},
			Formato:       sql.NullString{String: "LP", Valid: true},
			Pais:          sql.NullString{String: "USA", Valid: true},
			Condicion:     sql.NullString{String: "Very Good", Valid: true},
			DuracionTotal: sql.NullString{String: "45:44", Valid: true},
			ArteURL:       sql.NullString{String: "https://upload.wikimedia.org/wikipedia/en/9/9c/MilesDavisKindofBlue.jpg", Valid: true},
			Notas:         sql.NullString{String: "Considerado el mejor álbum de jazz de todos los tiempos", Valid: true},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "4",
			Titulo:        "Led Zeppelin IV",
			Artista:       "Led Zeppelin",
			Sello:         sql.NullString{String: "Atlantic", Valid: true},
			CatalogNumber: sql.NullString{String: "SD 7208", Valid: true},
			Anio:          sql.NullInt32{Int32: 1971, Valid: true},
			Formato:       sql.NullString{String: "LP", Valid: true},
			Pais:          sql.NullString{String: "USA", Valid: true},
			Condicion:     sql.NullString{String: "Mint", Valid: true},
			DuracionTotal: sql.NullString{String: "42:34", Valid: true},
			ArteURL:       sql.NullString{String: "https://upload.wikimedia.org/wikipedia/en/2/26/Led_Zeppelin_-_Led_Zeppelin_IV.jpg", Valid: true},
			Notas:         sql.NullString{String: "Incluye 'Stairway to Heaven', una de las canciones más famosas del rock", Valid: true},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            "5",
			Titulo:        "Thriller",
			Artista:       "Michael Jackson",
			Sello:         sql.NullString{String: "Epic", Valid: true},
			CatalogNumber: sql.NullString{String: "FE 38112", Valid: true},
			Anio:          sql.NullInt32{Int32: 1982, Valid: true},
			Formato:       sql.NullString{String: "LP", Valid: true},
			Pais:          sql.NullString{String: "USA", Valid: true},
			Condicion:     sql.NullString{String: "Near Mint", Valid: true},
			DuracionTotal: sql.NullString{String: "42:19", Valid: true},
			ArteURL:       sql.NullString{String: "https://en.wikipedia.org/wiki/Thriller_(album)#/media/File:Michael_Jackson_-_Thriller.png", Valid: true},
			Notas:         sql.NullString{String: "El álbum más vendido de todos los tiempos", Valid: true},
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	// Configurar géneros y estilos para cada record
	records[0].SetGeneros([]string{"Rock", "Progressive Rock"})
	records[0].SetEstilos([]string{"Psychedelic Rock", "Art Rock"})
	records[0].SetTracklist([]models.Track{
		{Numero: 1, Titulo: "Speak to Me", Duracion: "1:30"},
		{Numero: 2, Titulo: "Breathe (In the Air)", Duracion: "2:43"},
		{Numero: 3, Titulo: "On the Run", Duracion: "3:36"},
		{Numero: 4, Titulo: "Time", Duracion: "6:53"},
		{Numero: 5, Titulo: "The Great Gig in the Sky", Duracion: "4:36"},
		{Numero: 6, Titulo: "Money", Duracion: "6:23"},
		{Numero: 7, Titulo: "Us and Them", Duracion: "7:49"},
		{Numero: 8, Titulo: "Any Colour You Like", Duracion: "3:26"},
		{Numero: 9, Titulo: "Brain Damage", Duracion: "3:49"},
		{Numero: 10, Titulo: "Eclipse", Duracion: "2:03"},
	})

	records[1].SetGeneros([]string{"Rock", "Pop"})
	records[1].SetEstilos([]string{"Pop Rock", "Psychedelic Rock"})
	records[1].SetTracklist([]models.Track{
		{Numero: 1, Titulo: "Come Together", Duracion: "4:20"},
		{Numero: 2, Titulo: "Something", Duracion: "3:03"},
		{Numero: 3, Titulo: "Maxwell's Silver Hammer", Duracion: "3:27"},
		{Numero: 4, Titulo: "Oh! Darling", Duracion: "3:26"},
		{Numero: 5, Titulo: "Octopus's Garden", Duracion: "2:51"},
		{Numero: 6, Titulo: "I Want You (She's So Heavy)", Duracion: "7:47"},
		{Numero: 7, Titulo: "Here Comes the Sun", Duracion: "3:05"},
		{Numero: 8, Titulo: "Because", Duracion: "2:45"},
		{Numero: 9, Titulo: "You Never Give Me Your Money", Duracion: "4:02"},
		{Numero: 10, Titulo: "Sun King", Duracion: "2:26"},
		{Numero: 11, Titulo: "Mean Mr. Mustard", Duracion: "1:06"},
		{Numero: 12, Titulo: "Polythene Pam", Duracion: "1:12"},
		{Numero: 13, Titulo: "She Came In Through the Bathroom Window", Duracion: "1:57"},
		{Numero: 14, Titulo: "Golden Slumbers", Duracion: "1:31"},
		{Numero: 15, Titulo: "Carry That Weight", Duracion: "1:36"},
		{Numero: 16, Titulo: "The End", Duracion: "2:05"},
		{Numero: 17, Titulo: "Her Majesty", Duracion: "0:23"},
	})

	records[2].SetGeneros([]string{"Jazz"})
	records[2].SetEstilos([]string{"Modal Jazz", "Cool Jazz"})
	records[2].SetTracklist([]models.Track{
		{Numero: 1, Titulo: "So What", Duracion: "9:22"},
		{Numero: 2, Titulo: "Freddie Freeloader", Duracion: "9:46"},
		{Numero: 3, Titulo: "Blue in Green", Duracion: "5:37"},
		{Numero: 4, Titulo: "All Blues", Duracion: "11:33"},
		{Numero: 5, Titulo: "Flamenco Sketches", Duracion: "9:26"},
	})

	records[3].SetGeneros([]string{"Rock", "Hard Rock"})
	records[3].SetEstilos([]string{"Blues Rock", "Heavy Metal"})
	records[3].SetTracklist([]models.Track{
		{Numero: 1, Titulo: "Black Dog", Duracion: "4:54"},
		{Numero: 2, Titulo: "Rock and Roll", Duracion: "3:40"},
		{Numero: 3, Titulo: "The Battle of Evermore", Duracion: "5:52"},
		{Numero: 4, Titulo: "Stairway to Heaven", Duracion: "8:02"},
		{Numero: 5, Titulo: "Misty Mountain Hop", Duracion: "4:38"},
		{Numero: 6, Titulo: "Four Sticks", Duracion: "4:44"},
		{Numero: 7, Titulo: "Going to California", Duracion: "3:31"},
		{Numero: 8, Titulo: "When the Levee Breaks", Duracion: "7:07"},
	})

	records[4].SetGeneros([]string{"Pop", "R&B"})
	records[4].SetEstilos([]string{"Disco", "Funk"})
	records[4].SetTracklist([]models.Track{
		{Numero: 1, Titulo: "Wanna Be Startin' Somethin'", Duracion: "6:03"},
		{Numero: 2, Titulo: "Baby Be Mine", Duracion: "4:20"},
		{Numero: 3, Titulo: "The Girl Is Mine", Duracion: "3:42"},
		{Numero: 4, Titulo: "Thriller", Duracion: "5:57"},
		{Numero: 5, Titulo: "Beat It", Duracion: "4:18"},
		{Numero: 6, Titulo: "Billie Jean", Duracion: "4:54"},
		{Numero: 7, Titulo: "Human Nature", Duracion: "4:06"},
		{Numero: 8, Titulo: "P.Y.T. (Pretty Young Thing)", Duracion: "3:59"},
		{Numero: 9, Titulo: "The Lady in My Life", Duracion: "4:59"},
	})

	// Insertar records
	for _, record := range records {
		if err := repo.Create(record); err != nil {
			log.Printf("Error creando record %s: %v", record.Titulo, err)
		} else {
			log.Printf("✅ Creado: %s - %s", record.Artista, record.Titulo)
		}
	}

	log.Println("🎉 Seed completado exitosamente")
}
