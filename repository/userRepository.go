//	Repository implement required crud operations with Users collection. Use GetClient() before use any crud operation.
package repository

import (
	"context"
	"errors"
	pagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"testTaskBitmediaLabs/entity"
	"time"
)

const (
	DBUri          = "mongodb://localhost:2717"
	DBName         = "Users"
	CollectionName = "users"
)

const (
	nilMongoClientError = "error: mongo client can't be nil"
)

var repositoryClient *mongo.Client

//	Create single client instance for repository's crud operations. Use GetClient() before use any crud operation.
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

// Firstly important to check if mongo client have already existed
func checkMongoClient() error {
	if repositoryClient == nil {
		return errors.New(nilMongoClientError)
	}
	return nil
}

// Func uses collection.InsertMany() for insertion users to DB.
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

// Implementation of pagination data using gobeam/mongo-go-pagination. It requires number of documents to read and number of page.
func ReadUsersPagination(limit int64, page int64) (*[]entity.User, error) {
	err := checkMongoClient()
	if err != nil {
		return nil, err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)

	// Querying paginated data
	paginatedData, err := pagination.New(collection).Limit(limit).Page(page).Find()
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

// Function gets User document by ID, decodes it into instance of struct type User and returns it and error if it occurred.
func ReadUserByID(id string) (entity.User, error) {
	var result entity.User
	err := checkMongoClient()
	if err != nil {
		return entity.User{}, err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&result)
	return result, err
}

// CreateUser() returns ID of new User document
func CreateUser(doc interface{}) (interface{}, error) {
	err := checkMongoClient()
	if err != nil {
		return nil, err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	result, err := collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// One of Update operation's variant: to replace User document by ID
func ReplaceUserByID(objectID primitive.ObjectID, doc interface{}) error {
	err := checkMongoClient()
	if err != nil {
		return err
	}
	collection := repositoryClient.Database(DBName).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = collection.ReplaceOne(ctx, bson.M{"_id": objectID}, doc)
	return err
}
