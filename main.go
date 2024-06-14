package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func main() {
	r := gin.Default();

	// Middleware function
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start)
		log.Printf("Request processed in %s", duration)
	})

	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"PONG",
		})
	})

	r.Run(":8080")
}


// r.Use(...) registers a global middleware function. Middleware in Gin is a way to process a request before it reaches the route handler and/or after the route handler has processed it.


// c.Next(): Calls the next middleware in the chain or, if this is the last middleware, the route handler. This is where the actual request processing happens.




// Process 

// Request Arrival

// A request to http://localhost:8080/ping arrives at the server.

// Middleware Execution

// The middleware function is executed:
// start := time.Now() records the start time.
// c.Next() passes control to the next handler in the chain, which in this case is the route handler for /ping.


// Route Handler Execution

// The route handler responds with {"message": "pong"} and finishes processing.



// ost-Processing Middleware Execution

// After the route handler has finished, the control returns to the middleware:
// duration := time.Since(start) calculates the total time taken to process the request.
// log.Printf("Request processed in %s", duration) logs the processing time.


// Response Sent

// The server sends the response back to the client.


// Purpose of the Middleware
// The middleware in this example serves to log the time taken to process each request. This is useful for monitoring the performance of your application and identifying slow requests.

// By using c.Next(), the middleware ensures that the request is processed by the subsequent handlers and then performs additional actions (logging the duration) after the request has been processed. This is a common pattern in web frameworks for handling tasks such as logging, authentication, and error handling.