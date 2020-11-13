package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jeluchez/twittergo/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un email de parámetro y chequea si ya está en la BD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	fmt.Println("aqui llega")

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	fmt.Println(resultado)

	ID := resultado.ID.Hex()
	if err != nil {

		return resultado, false, ID
	}
	return resultado, true, ID
}
