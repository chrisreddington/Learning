package main

import (
	"fmt"
	"net/http"

	"github.com/chrisreddington/learning/golang/mymodule/controllers"
)

func main() {

	fmt.Println("Registering controllers...")

	controllers.RegisterControllers()
	fmt.Println("Registered controllers!")
	fmt.Println("Starting server on localhost:3000")
	http.ListenAndServe(":3000", nil)
}

/*u1 := models.User{
	ID:        1,
	FirstName: "Chris",
	LastName:  "Reddington",
}
u2 := models.User{
	ID:        2,
	FirstName: "Chris",
	LastName:  "Bloggs",
}

	if u1.ID == u2.ID {
		println("Same user!")
	} else if u1.FirstName == u2.FirstName {
		print("Similar users...")
	} else {
		print("Different users!")
	}*/

/*func main() {
	u := models.User{
		ID:        2,
		FirstName: "Chris",
		LastName:  "Reddington",
	}

	fmt.Println(u)
	port := 3000
	_, err := startWebServer(port, 2)

	fmt.Println(err)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// do things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, nil
}*/
