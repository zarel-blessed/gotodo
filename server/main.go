package main

import (
	"context"
	"fmt"
	"log"
	connect "server/connection"
	"server/model/todo"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func main() {
	client := connect.ToMongoDB()

	collection = client.Database("testdb").Collection("books")

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/todos", createTodo)
	router.GET("/todos", getTodos)
	router.POST("/todos/:id", getTodoById)
	router.DELETE("/todos/:id", deleteTodoById)
	router.Run("localhost:3080")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect MongoDB client!")
		}
	}()
}

func createTodo(c *gin.Context) {
	var newTodo todo.Model

	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input!"})
		return
	}

	newTodo.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), newTodo)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create new todo!"})
		return
	}

	c.JSON(201, newTodo)
}

func getTodos(c *gin.Context) {
	var todos []todo.Model

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	curr, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Finding Time Exceed: " + fmt.Sprintf("%v", err)})
		return
	}
	defer curr.Close(ctx)

	for curr.Next(ctx) {
		var todo todo.Model

		if err := curr.Decode(&todo); err != nil {
			c.JSON(501, gin.H{"error": "Failed to decode todo!"})
			return
		}

		todos = append(todos, todo)
	}

	c.JSON(200, todos)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Id"})
		return
	}

	var todo todo.Model
	err = collection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&todo)
	if err != nil {
		c.JSON(404, gin.H{"error": "Todo not found!"})
		return
	}

	c.JSON(200, todo)
}

func deleteTodoById(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Id"})
		return
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(200, gin.H{"message": "Todo deleted successfully!"})
}
