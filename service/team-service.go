package service
/*
import(
	"errors"
	"simpleAPI/entity"
	"simpleAPI/repository"
)

var (
	repo = repository.NewFirestoreRepository()
	//This can be changed to use any type of DB
)

type LeagueService interface{
	Validate(league *entity.League) error
	Create(league *entity.League) (*entity.League, error)
	GetAll() ([]entity.League, error)
}

type service struct{}

func NewLeagueService() LeagueService {
	return &service{}
}

func (*service) Validate(league *entity.League) error{
	if(league != nil){
		err := errors.New("The league is empty")
		return err
	}
	if(league.Name == ""){
		err := errors.New("The league name can't be empty")
		return err
	}
	return nil
}

func (*service) Create(league *entity.League) (*entity.League, error){
	return repo.Save(league)
}

func (*service) GetAll() ([]entity.League, error){
	return repo.GetAll()
}
*/