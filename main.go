package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// r := httprouter.New()
	// uc := controllers.NewUserController(getSession())
	// r.GET("/user/:id", uc.GetUser)
	// r.POST("/user", uc.CreateUser)
	// // r.DELETE("/user/:id", uc.DeleteUser)
	// http.ListenAndServe("localhost:9000", r)

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "olá mundo\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func getSession() *mgo.Session {
// 	s, err := mgo.Dial("mongodb://localhost:27017")
// 	if err != nil {
// 		panic(err)
// 	}
// 	return s
// }
