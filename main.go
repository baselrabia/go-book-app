package main

import (
	"github.com/baselrabia/go-book-app/models"
	"github.com/baselrabia/go-book-app/routes"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// Middleware
	
	e.Use(middleware.Logger())
	/**
	This line adds a logger middleware to your Echo application. 
	The middleware.Logger() middleware logs details about incoming HTTP
	requests and outgoing responses. It typically logs information 
	such as the HTTP method, the requested URL, 
	the response status code, and the time taken to process the request.
	The logs are useful for debugging and monitoring 
	your application's behavior. They provide a record of 
	what's happening in your application and can help you identify 
	and troubleshoot issues.
	*/
	e.Use(middleware.Recover())
	/**
	This line adds a recover middleware to your Echo application.
	The middleware.Recover() middleware is a safety net 
	for your application. It recovers from panics 
	(unexpected errors that cause your application to crash) 
	that occur in your route handlers. If a panic happens,
	the recover middleware will catch it, prevent the application 
	from crashing, and respond with a 500 Internal Server Error 
	status code. This middleware is crucial for ensuring that
	your application remains robust and continues to function, 
	even if there are unexpected errors.
	*/


	// Database setup
	db := models.InitDB("books.db")
	models.Migrate(db)

	//routes
	// Routes
	api := e.Group("/api")
	routes.SetupRoutes(api)

	e.Logger.Fatal(e.Start(":8080"))
}
