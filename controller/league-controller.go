package controller

import(
	"fmt"
	"net/http"
	"encoding/json"
	
	"simpleAPI/entity"
	"simpleAPI/errors"
	"simpleAPI/service"
)

type controller struct{}

var (
	leagueService service.LeagueService = service.NewLeagueService()
)

type LeagueController interface{
	GetLeagues(writer http.ResponseWriter, request *http.Request)
	AddLeague(writer http.ResponseWriter, request *http.Request)
	HomePage(writer http.ResponseWriter, request *http.Request)
}

func NewLeagueController() LeagueController{
	return &controller{}
}

func (*controller) GetLeagues(writer http.ResponseWriter, request *http.Request){

	writer.Header().Set("Content-Type", "application/json")
	leagues, err := leagueService.GetAll()
	if(err != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errors.ServiceError{Message: "Error getting the leagues"})
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(leagues)
	
	
}


func (*controller) AddLeague(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var league entity.League
	err := json.NewDecoder(request.Body).Decode(&league)
	if(err != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errors.ServiceError{Message: "Error adding the league"})
		return
	}
	
	err1 := leagueService.Validate(&league)
	if(err1 != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := leagueService.Create(&league)

	if(err2 != nil){
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errors.ServiceError{Message: "Error saving the league"})
		return
	}

	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(result)
}


func (*controller) HomePage(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Homepage Endpoint Hit")
}