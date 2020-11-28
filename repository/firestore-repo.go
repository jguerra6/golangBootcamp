package repository

import(
	"context"
	"log"
	"simpleAPI/entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)



type repo struct{}

//Create a new Firestore instance
func NewFirestoreRepository() LeagueRepository{
	return &repo{}
}

func createClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
			log.Fatalf("Failed to create Firestore Client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

const (
	projectId string = "bootcamp-d79dd"
	collectionName string = "leagues"
)

func (*repo) Save(league *entity.League) (*entity.League, error){
	ctx := context.Background()
	
	client := createClient(ctx)

	

	_, _, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"country": league.Country,
		"name": league.Name,
		"current_season_id": league.Current_season_id,
	})

	if(err != nil){
		log.Fatal("Failed adding league: ", err)
		return nil, err
	}

	defer client.Close()
	return league, nil
}

func (*repo) GetAll() ([]entity.League, error){
	ctx := context.Background()
	
	client := createClient(ctx)
	
	var leagues []entity.League
	iter := client.Collection(collectionName).Documents(ctx)
	//log.Println(*iterator);
	for{
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		
		if(err != nil){
			log.Fatal("Failed get all the leagues: ", err)
			return nil, err
		}

		league := entity.League{
			Country: doc.Data()["country"].(string),
			Name: doc.Data()["name"].(string),
			Current_season_id: doc.Data()["current_season_id"].(int64),
		}

		leagues = append(leagues, league)
	}
	
	defer client.Close()
	return leagues, nil

}