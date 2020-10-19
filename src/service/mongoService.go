package service

import (
	"bundle_magnet/src/models/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*
	MongoDB ye bağlanmak için kullanılır.
 */
func ConnectMongoDB(srv string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(srv)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return client, nil
}

/*
	Geçerli koleksiyonu getirmek için
 */
func GetCollection(client mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)

	return collection
}

/*
	Parametre geçilen koleksiyonda arama yapmak için
 */
func FindData(collection mongo.Collection, filter bson.D) ([]*entity.Channel, error) {
	var results []*entity.Channel

	findOptions := options.Find()

	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem entity.Channel
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)

			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)

		return nil, err
	}

	cur.Close(context.TODO())

	return results, nil
}