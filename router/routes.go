package router

import(
	"fmt"
	"net/http"
	"encoding/json"
	"simpleAPI/repository"
	"simpleAPI/entity"
)



var (
	repo repository.LeagueRepository = repository.NewLeagueRepository()
)

func GetLeagues(writer http.ResponseWriter, request *http.Request){
	/*
	leagues = []League{
		League{
			Id:"A3418348-7DB9-4BA3-A0B8-668B95E17E27", 
			Name: "Serie A", 
			Current_season_id: 32523,
		},
	}
	*/

	//fmt.Println("Endpoint Hit: All Leagues Endpoint")

	writer.Header().Set("Content-Type", "application/json")
	leagues, err := repo.FindAll()
	if(err != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "Error getting the leagues"}`))
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(leagues)
	
	//json.NewEncoder(writer).Encode(articles)
}

/*
*/
func AddLeague(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var league entity.League
	err := json.NewDecoder(request.Body).Decode(&league)
	if(err != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "Error adding the league"}`))
		return
	}
	
	repo.Save(&league)
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(league)
}
/*
*/

func HomePage(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Homepage Endpoint Hit")
}