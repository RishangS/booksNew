<!-- Books App -->

Install Go
`https://go.dev/doc/install`


Place the project in the src path of your Go folder
    `go/src`

Install MongoDb
    `https://docs.mongodb.com/manual/installation/`

Run MongoDb locally and copy the mongo URL to `app.env` file


Port Number can be modified in `app.env` file

Run bookapp using
    `go run. `

In any browser open the local host with the same port as the application
    `http://localhost:8000/` 

Home page shows all the books present in db
To add new book click `Add Book` and add a new book by entering title and author name
after success response click on `Home` where the book will be visible and a delete option is also visible
Click on `delete` to delete the particular book
