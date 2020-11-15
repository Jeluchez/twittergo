package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jeluchez/twittergo/bd"
	"github.com/jeluchez/twittergo/models"
)

/*AltaRelacion realiza el registro de la relacion entre usuarios */
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	type JSONRes struct {
		Status bool
		Msg    string
	}
	response := JSONRes{true, "AltaRelacion"}

	res, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
