//	Use GetClient() before use any crud operation
package repository

import (
	"context"
	"errors"
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

const (
	nilContextError = "error: context can't be nil"
)

var repositoryClient *mongo.Client

//	Create single client and context instances for repository's crud operations
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(DBUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	repositoryClient = client
	return client
}

func checkMongoClient() error {
	if repositoryClient == nil {
		return errors.New(nilContextError)
	}
	return nil
}

//	Use GetClient() before use crud operation. When we have JSON with data that are too large it's better use mongoimport
func InsertUsers(docs []interface{}) error {
	err := checkMongoClient()
	if err != nil {
		return err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	//unique index for Email field
	ctx := context.TODO()
	_, err = collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bsonx.Doc{{"email", bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		},
	)

	_, err = collection.InsertMany(ctx, docs)
	return err
}

//	Use GetClient() before use crud operation
func ReadUsersPagination(limit int64, page int64) (*[]entity.User, error) {
	err := checkMongoClient()
	if err != nil {
		return nil, err
	}
	filter := bson.M{}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)

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

//	Use GetClient() before use crud operation
func ReadUserByID(id string) (entity.User, error) {
	// Here's user decoded document
	var result entity.User
	err := checkMongoClient()
	if err != nil {
		return entity.User{}, err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	return result, err
}

//	Use GetClient() before use crud operation
func CreateUser(doc interface{}) (interface{}, error) {
	err := checkMongoClient()
	if err != nil {
		return nil, err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

//	Use GetClient() before use crud operation
func ReplaceUser(objectID primitive.ObjectID, doc interface{}) error {
	err := checkMongoClient()
	if err != nil {
		return err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	_, err = collection.ReplaceOne(ctx, bson.M{"_id": objectID}, doc)
	return err
}
