// handlers.book.go

package main

import (
	"booksnew/database"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = database.Db().Database("goTest").Collection("books")

func showIndexPage(c *gin.Context) {

	// books := getAllBooks()

	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		c.AbortWithStatus(http.StatusBadRequest)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO())
	fmt.Println("results")
	fmt.Println(results)

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": results,
		},
	)

}

func getBook(c *gin.Context) {
	var result book
	fmt.Println("param id")
	searchID := c.Param("book_id")
	searchID_int, _ := strconv.Atoi(searchID)
	fmt.Printf("%v", searchID_int)

	err := userCollection.FindOne(context.TODO(), bson.M{"id": searchID_int}).Decode(&result)
	if err != nil {

		c.AbortWithError(http.StatusNotFound, err)
	}
	fmt.Println("result")

	fmt.Println(result.ID)
	fmt.Println(result.Author)
	fmt.Println(result.Title)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"book.html",
		// Pass the data that the page uses
		gin.H{
			"title":   result.Title,
			"payload": result,
		},
	)
}

func showBookCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Add New Book"}, "add-book.html")

}

func createBook(c *gin.Context) {

	var newbook book
	newbook.Title = c.PostForm("title")
	newbook.Author = c.PostForm("author")
	newbook.ID = getbookID()
	insertResult, err := userCollection.InsertOne(context.TODO(), newbook)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	fmt.Println("Inserted a single document: ", insertResult)
	render(c, gin.H{
		"title":   "Submission Successful",
		"payload": newbook}, "submission-successful.html")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}

func search(c *gin.Context) {
	// Check if the book ID is valid
	bookAuthor := c.PostForm("author")
	// Check if the book exists
	fmt.Println("Author Name ")
	fmt.Println(bookAuthor)
	var result book

	err := userCollection.FindOne(context.TODO(), bson.M{"author": bookAuthor}).Decode(&result)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("result")

	fmt.Println(result.ID)
	fmt.Println(result.Author)
	fmt.Println(result.Title)

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"book.html",
		// Pass the data that the page uses
		gin.H{
			"title":   result.Title,
			"payload": result,
		},
	)
}

func getbookID() int {

	var results []primitive.M                                   //slice for multiple documents
	cur, err := userCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		fmt.Println(err)

	}
	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO())
	fmt.Println("results")
	fmt.Println(results)
	ID := len(results) + 1
	fmt.Println("ID")
	fmt.Println(ID)
	return ID
}

func delete(c *gin.Context) {
	fmt.Println("param id")
	searchID := c.Param("book_id")
	searchID_int, _ := strconv.Atoi(searchID)
	fmt.Printf("%v", searchID_int)

	res, err := userCollection.DeleteOne(context.TODO(), bson.M{"id": searchID_int})
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println("result")

	fmt.Println(res)
	c.Redirect(http.StatusFound, "/")
}
