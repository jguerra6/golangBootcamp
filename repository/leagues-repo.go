package repository

import(
	"context"
	"log"
	"simpleAPI/entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type LeagueRepository interface {
	Save(league *entity.League) (*entity.League, error)
	FindAll() ([]entity.League, error)
}

type repo struct{}

func NewLeagueRepository() LeagueRepository{
	return &repo{}
}

func createClient(ctx context.Context) *firestore.Client {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
			log.Fatalf("Failed to create client: %v", err)
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
	client, err := firestore.NewClient(ctx, projectId)
	if(err != nil){
		log.Fatal("Failed to create a Firestore Client: ", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"country": league.Country,
		"name": league.Name,
		"current_season_id": league.Current_season_id,
	})

	if(err != nil){
		log.Fatal("Failed adding league: ", err)
		return nil, err
	}

	return league, nil
}

func (*repo) FindAll() ([]entity.League, error){
	ctx := context.Background()
	/*
	client, err := firestore.NewClient(ctx, projectId)
	if(err != nil){
		log.Fatal("Failed to create a Firestore Client: ", err)
		return nil, err
	}*/
	client := createClient(ctx)
	defer client.Close()
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
	return leagues, nil

}