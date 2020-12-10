package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sakiib/GoStuff/Packages/myPackages"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(home)
}

func main() {
	fmt.Println("Packages!")
	fmt.Println(myPackages.Min(10, 20), myPackages.Max(10, 20))
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	http.ListenAndServe(":8080", router)
}