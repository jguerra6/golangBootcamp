package main

import(
	
	"log"
	"net/http"
	
	"github.com/gorilla/mux"

	"simpleAPI/router"
)


func handleRequests(){
	//Define the port
	const port string = ":8081"

	//Create the mux router to handle all the paths and methods
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", router.HomePage)
	myRouter.HandleFunc("/leagues", router.GetLeagues).Methods("GET")
	myRouter.HandleFunc("/leagues", router.AddLeague).Methods("POST")
	log.Println("Listening to port ", port)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}