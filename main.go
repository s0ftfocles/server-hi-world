package main

import (
	"fmt"
	"net/http"
)

func manejadorHola(w http.ResponseWriter, r *http.Request) {
	// Esto es lo que el servidor le va a responder al navegador web
	fmt.Fprintln(w, "<h1>¡Hola Mundo desde mi propio servidor en Go! 🚀</h1>")
	fmt.Fprintln(w, "<p>Felicidades, este servidor está corriendo de forma nativa en tu Kali Linux.</p>")
}

func main() {
	// Le decimos al servidor que escuche en la ruta principal "/" y use nuestra función
	http.HandleFunc("/", manejadorHola)

	fmt.Println("🌐 Servidor web iniciado en http://localhost:8080")
	fmt.Println("Presiona Ctrl + C en la terminal para apagarlo.")

	// Arrancamos el servidor en el puerto 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error al arrancar el servidor:", err)
	}
}
