package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings" // Importación para manejo de cadenas
)

func main() {
	// Flag para decidir si se muestran solo archivos de imagen
	var soloImagenes bool
	flag.BoolVar(&soloImagenes, "i", false, "Mostrar solo archivos de imagen (.jpg, .jpeg, .png)")

	// Configurar flag para el directorio
	var directorio string
	flag.StringVar(&directorio, "d", ".", "Directorio a listar")

	flag.Parse()

	// Leer el directorio
	files, err := leerDirectorio(directorio)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	absPath, err := obtenerRutaAbsoluta(directorio)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Archivos en %s:\n", absPath)
	if soloImagenes {
		// Mostrar solo archivos de imagen
		mostrarImagenes(files)
	} else {
		// Mostrar todos los archivos
		mostrarTodos(files)
	}
}

// leerDirectorio lee el contenido de un directorio y retorna su lista de entradas
func leerDirectorio(directorio string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(directorio)
	if err != nil {
		return nil, fmt.Errorf("error al leer el directorio: %w", err)
	}
	return files, nil
}

// mostrarImagenes recibe una lista de entradas y muestra solo aquellas que correspondan a archivos de imagen
func mostrarImagenes(files []os.DirEntry) {
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			log.Printf("Error obteniendo info de %s: %v", file.Name(), err)
			continue
		}
		if !info.IsDir() {
			lowerName := strings.ToLower(file.Name())
			if strings.HasSuffix(lowerName, ".jpg") ||
				strings.HasSuffix(lowerName, ".jpeg") ||
				strings.HasSuffix(lowerName, ".png") {
				fmt.Println("- " + file.Name())
			}
		}
	}
}

// mostrarTodos muestra el nombre de todas las entradas encontradas en el directorio
func mostrarTodos(files []os.DirEntry) {
	for _, file := range files {
		fmt.Println("- " + file.Name())
	}
}

// obtenerRutaAbsoluta retorna la ruta absoluta del directorio dado
func obtenerRutaAbsoluta(directorio string) (string, error) {
	absPath, err := filepath.Abs(directorio)
	if err != nil {
		return "", fmt.Errorf("error obteniendo ruta absoluta: %w", err)
	}
	return absPath, nil
}
