package main

import (
	"fmt"
	"net/http"
	"os"
)

func manejadorHola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>¡Hola Mundo desde mi propio servidor en Go! 🚀</h1>")
}

func manejadorPrueba(w http.ResponseWriter, r *http.Request) {
	contenido, err := os.ReadFile("prueba.txt")
	if err != nil {
		http.Error(w, "No se pudo leer el archivo de prueba", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "<h2>📄 Contenido de tu archivo de prueba:</h2>")
	fmt.Fprintf(w, "<pre>%s</pre>", string(contenido))
}

func main() {
	http.HandleFunc("/", manejadorHola)
	http.HandleFunc("/prueba", manejadorPrueba)

	fmt.Println("🌐 Servidor web iniciado en http://localhost:8080")
	fmt.Println("Presiona Ctrl + C en la terminal para apagarlo.")
	http.ListenAndServe(":8080", nil)
}
