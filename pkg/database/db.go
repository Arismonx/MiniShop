package database

import (
	"context"
	"log"
	"time"

	"github.com/Arismonx/MiniShop/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func DBConnect(pctx context.Context, cfg *config.Config) *mongo.Client {
	//Cancel connect database after 5 sec.
	ctx, cancel := context.WithTimeout(pctx, 5*time.Second)
	defer cancel()

	//Connecting Database
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Db.Url))

	if err != nil {
		log.Fatal("Error: Connect to database error: ", err)
	}

	//Pinging
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error: Pinging to database error: ", err)
	}

	return client
}
