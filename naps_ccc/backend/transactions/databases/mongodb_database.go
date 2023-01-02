package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	user     = "pablo"
	password = "12345"
	database = "ccc"
	host     = "cluster0.rzcky.mongodb.net"
	// port     = 1234 // como es atlas no hay difinicion de puerto
)

/*
var (
	//trs     = "desarrollos3c"
	//password = "1n9d3s4rr0ll0s"
	database = "fac_mongo"
	//host     = "172.30.0.141"
	//port     = "27017"
)
*/
func GetCollection(collection string) *mongo.Collection {
	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, database, port)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", user, password, host, database)
	//uri := "mongodb://desarrollos3c:1n9d3s4rr0ll0s@172.30.0.141:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://desarrollos3c:1n9d3s4rr0ll0s@172.30.0.141:27017"))

	logErrors(err)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	logErrors(err)

	return client.Database(database).Collection(collection)
}

func logErrors(err error) {
	if err != nil {
		log.Fatal("ERROR!!! ", err)
	}
}
