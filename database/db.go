package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Database struct {
	client *mongo.Client
}

func Connect() (*Database, error) {
	// carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Erro ao carregar as variáveis de ambiente: %v", err)
	}

	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &Database{
		client: client,
	}, nil
}

func (db *Database) Close() {
	err := db.client.Disconnect(context.TODO())

	if err != nil {
		return
	}

	fmt.Println("Disconnected from MongoDB!")
}

func (db *Database) Collection(name string) *mongo.Collection {
	return db.client.Database("rastros_da_mata_db").Collection(name)
}
