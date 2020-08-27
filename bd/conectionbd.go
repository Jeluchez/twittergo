package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoC  objeto de la coneccion
var MongoC = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:Twitter123@tweetergo.gkl0y.mongodb.net/<dbname>?retryWrites=true&w=majority")

// // context
// en go no existen lo que son variables globales. Pero context nos provee un espacio en memoria para guardar informacion del contexto siempre y cuando se esté ejecutando

//ConnectBD esta funcion hace la coneccion a la db
func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	//  si err es distinto de nil significa que hubo un error en la BD
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("success conection with the BD")

	return client
}

//CheckConnection  verificacion de que si se esté haciendo ping
func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	//  si err es distinto de nil significa que hubo un error
	if err != nil {
		return 0
	}
	return 1
}
