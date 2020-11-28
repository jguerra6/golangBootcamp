package main

import(
	
	
	"simpleAPI/controller"
	"simpleAPI/router"
	"simpleAPI/service"
	"simpleAPI/repository"
)

var(
	leagueRepository repository.LeagueRepository = repository.NewFirestoreRepository()
	leagueService service.LeagueService = service.NewLeagueService(leagueRepository)
	leagueController controller.LeagueController = controller.NewLeagueController(leagueService)
	//leagueController = controller.NewLeagueController(leagueService)
	httpRouter = router.NewMuxRouter()
)

func handleRequests(){
	//Define the port
	const port string = ":8081"

	
	//Handle the routes

	httpRouter.GET("/", leagueController.HomePage)
	httpRouter.GET("/leagues", leagueController.GetLeagues)
	httpRouter.POST("/leagues", leagueController.AddLeague)

	//Create the router to handle all the paths and methods
	httpRouter.SERVE(port)

	
}

func main(){
	handleRequests()
}

