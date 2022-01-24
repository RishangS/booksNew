// routes.go

package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)
	router.GET("/book/view/:book_id", getBook)
	router.GET("/book/add", showBookCreationPage)
	router.POST("/book/addnew", createBook)
	router.POST("/book/search", search)
	router.GET("/book/delete/:book_id", delete)

}
