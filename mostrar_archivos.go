package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Configurar flag para el directorio
	var directorio string

	//d para especificar el directorio
	//. para usar el directorio actual
	flag.StringVar(&directorio, "d", ".", "Directorio a listar")

	//Procesa los argumentos y asigna los valores a las variables correspondientes, en este caso
	//a la variable directorio
	flag.Parse()

	// Leer contenido del directorio
	files, err := os.ReadDir(directorio)
	if err != nil {
		log.Fatalf("Error al leer el directorio: %v", err)
	}

	// Obtener la ruta absoluta
	absPath, err := filepath.Abs(directorio)
	if err != nil {
		log.Fatalf("Error obteniendo ruta absoluta: %v", err)
	}

	fmt.Printf("Archivos en %s:\n", absPath)

	// Listar solo los archivos (excluyendo directorios)
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			log.Printf("Error obteniendo info de %s: %v", file.Name(), err)
			continue
		}

		// Mostrar el nombre solo si es un archivo
		if !info.IsDir() {
			fmt.Println("- " + file.Name())
		}
	}
}
