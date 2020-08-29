package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/jeluchez/twittergo/middlew"
	"github.com/jeluchez/twittergo/routers"
)

// Manejadores pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	// handlers
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// dar permisos desde cualquier dispositivo remoto desde cualquier router(ruta)
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
