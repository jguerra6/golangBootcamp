package main

import(
	
	"log"
	"net/http"
	"fmt"
	"simpleAPI/controller"
	"simpleAPI/router"
	
)

var(
	leagueController controller.LeagueController = controller.NewLeagueController()
	httpRouter = router.NewMuxRouter()
)

func handleRequests(){
	//Define the port
	const port string = ":8081"

	//Create the mux router to handle all the paths and methods
	//myRouter := mux.NewRouter().StrictSlash(true)
	httpRouter.GET("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprintln(writer, "Up and running...")
	})
	httpRouter.GET("/leagues", leagueController.GetLeagues)
	httpRouter.POST("/leagues", leagueController.AddLeague)

	httpRouter.SERVE(port)

	log.Println("Listening to port ", port)
	//log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}

