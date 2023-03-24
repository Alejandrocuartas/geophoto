package collections

import (
	"os"

	"github.com/Alejandrocuartas/geophoto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func PhotoCollection() *mongo.Collection {
	myDatabase := database.Client.Database(os.Getenv("MONGO_DB"))
	PhotoCol := myDatabase.Collection("photo")
	return PhotoCol
}
