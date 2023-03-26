package collections

import (
	"os"

	"github.com/Alejandrocuartas/geophoto/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserCollection() *mongo.Collection {
	myDatabase := database.Client.Database(os.Getenv("MONGO_DB"))
	UserCol := myDatabase.Collection("user")
	return UserCol
}
