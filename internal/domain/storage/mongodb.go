package storage

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoDB struct {
	collection *mongo.Collection
}

func NewMongoDB(hostName, databaseName, collectionName, userName, password string, timeout time.Duration) (*MongoDB, error) {
	database, err := connect(hostName, databaseName, userName, password, timeout)
	if err != nil {
		return &MongoDB{}, err
	}

	return &MongoDB{
		collection: database.Collection(collectionName),
	}, nil
}

func (db *MongoDB) SaveTodo(ctx context.Context, name, description string) (Todo, error) {
	if _, err := db.GetTodoByName(ctx, name); err != ErrNotFound {
		return Todo{}, ErrAlreadyInList
	}

	todo := Todo{
		ID: createID(),
		Name: name,
		Description: description,
	}

	if _, err := db.collection.InsertOne(ctx, todo); err != nil {
		return Todo{}, fmt.Errorf("storage.SaveTodo got error on insert: %v", err)
	}

	return todo, nil
}

func (db *MongoDB) GetTodoList(ctx context.Context) ([]Todo, error) {
	todos := []Todo{}

	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return todos, fmt.Errorf("storage.GetTodoList failed to find a collection cursor: %v", err)
	}

	for cursor.Next(ctx) {
		var todo Todo
		if err = cursor.Decode(&todo); err != nil {
			return todos, fmt.Errorf("storage.GetTodoList: cursor failed to decode next todo in collection: %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (db *MongoDB) GetTodoByName(ctx context.Context, name string) (Todo, error) {
	log.Printf("storage.GetTodoByName starting to search to todo with name %v", name)
	var todo Todo

	if err := db.collection.FindOne(ctx, bson.M{"name": name}).Decode(&todo); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Todo{}, ErrNotFound
		}
		return Todo{}, fmt.Errorf("storage.GetTodoByName got unexpected error on FindOne: %v", err)
	}
	log.Printf("storage.GetTodoByName found todo: %+v", todo)

	return todo, nil
}

func (db *MongoDB) EditTodo(ctx context.Context, id string, todo Todo) (Todo, error) {
	todoUpdate := bson.M{
		"$set": bson.M{
			"name":        todo.Name,
			"description": todo.Description,
		},
	}

	if _, err := db.collection.UpdateOne(ctx, bson.M{"id": id}, todoUpdate); err != nil {
		return Todo{}, fmt.Errorf("storage.EditTodo got error from UpdateOne: %v", err)
	}

	todo.ID = id

	return todo, nil
}

func (db *MongoDB) DeleteTodo(ctx context.Context, id string) error {
	if _, err := db.collection.DeleteOne(ctx, bson.M{"id": id}); err != nil {
		return fmt.Errorf("storage.DeleteTodo got error from DeleteOne: %v", err)
	}

	return nil
}

func (db *MongoDB) ClearTodoList(ctx context.Context) error {
	database := db.collection.Database()
	collectionName := db.collection.Name()

	if err := db.collection.Drop(ctx); err != nil {
		return fmt.Errorf("storage.ClearTodoList got error from collection.Drop: %v", err)
	}

	if err := database.CreateCollection(ctx, collectionName); err != nil {
		return fmt.Errorf("storage.ClearTodoList got error from database.CreateCollection: %v", err)
	}

	return nil
}

func connect(hostName, databaseName, userName, password string, timeout time.Duration) (*mongo.Database, error) {
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", userName, password, hostName, databaseName)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return &mongo.Database{}, fmt.Errorf("storage.connect failed to create mongo client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return &mongo.Database{}, fmt.Errorf("storage.connect failed to connect to mongo client: %v", err)
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		return &mongo.Database{}, fmt.Errorf("storage.connect failed to ping mongo client: %v", err)
	}

	return client.Database(databaseName), nil
}