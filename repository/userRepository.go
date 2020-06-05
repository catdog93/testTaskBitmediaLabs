package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	//"log"
	"testTaskBitmediaLabs/entity"
)

const (
	DBUri          = "mongodb://localhost:2717"
	DBName         = "Users"
	CollectionName = "users"
)

// When we have files with data that are too large it's better use mongoimport.
func InsertUsers(docs []interface{}) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(DBUri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	collection := client.Database(DBName).Collection(CollectionName)

	//unique index for Email field
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)

	_, err = collection.InsertMany(context.TODO(), docs)
	return err
}

func ReadUsers(limit uint64) ([]entity.User, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(DBUri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database(DBName).Collection(CollectionName)

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))

	// Here's an array in which you can store the decoded documents
	var results []entity.User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}
	// Close the cursor once finished
	defer cur.Close(context.TODO())
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem entity.User
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}
	return results, nil
}

func ReadUser(id string) (*entity.User, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(DBUri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database(DBName).Collection(CollectionName)

	// Here's user decoded document
	var result entity.User

	objectID, err := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)
	return &result, err
}

func CreateUser(doc interface{}) (interface{}, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(DBUri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database(DBName).Collection(CollectionName)

	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func ReplaceUser(objId primitive.ObjectID, doc interface{}) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(DBUri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	collection := client.Database(DBName).Collection(CollectionName)

	_, err = collection.ReplaceOne(context.TODO(), bson.M{"_id": objId}, doc)
	return err
}
