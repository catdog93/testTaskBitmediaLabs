package repository

import (
	"context"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"time"

	//"log"
	"testTaskBitmediaLabs/entity"
)

const (
	DBUri          = "mongodb://localhost:2717"
	DBName         = "Users"
	CollectionName = "users"
)

var clientRepository *mongo.Client

// create one client for repository's crud operations
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(DBUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	clientRepository = client
	return client
}

// When we have JSON with data that are too large it's better use mongoimport.
func InsertUsers(docs []interface{}) error {
	collection := clientRepository.Database(DBName).Collection(CollectionName)

	//unique index for Email field
	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)

	_, err = collection.InsertMany(context.TODO(), docs)
	return err
}

func ReadUsersPagination(limit int64, page int64) (*[]entity.User, error) {
	filter := bson.M{}
	collection := clientRepository.Database(DBName).Collection(CollectionName)

	// Querying paginated data
	paginatedData, err := mongopagination.New(collection).Limit(limit).Page(page).Filter(filter).Find()
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for _, raw := range paginatedData.Data {
		var user *entity.User
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			users = append(users, *user)
		}
	}
	return &users, nil
}

func ReadUserByID(id string) (entity.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := clientRepository.Database(DBName).Collection(CollectionName)

	// Here's user decoded document
	var result entity.User

	objectID, err := primitive.ObjectIDFromHex(id)
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	return result, err
}

func CreateUser(doc interface{}) (interface{}, error) {
	collection := clientRepository.Database(DBName).Collection(CollectionName)

	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func ReplaceUser(objectID primitive.ObjectID, doc interface{}) error {
	collection := clientRepository.Database(DBName).Collection(CollectionName)

	_, err := collection.ReplaceOne(context.TODO(), bson.M{"_id": objectID}, doc)
	return err
}
