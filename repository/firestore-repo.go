package repository

import (
	"context"
	"log"

	"github.com/luisjn/scores-api/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

//NewFirestoreRepository creates a new repo (constructor)
func NewFirestoreRepository() ScoreRepository {
	return &repo{}
}

const (
	projectId      string = "golang-api-a091c"
	collectionName string = "scores"
)

func (*repo) Save(score *entity.Score) error {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"id":        score.Id,
		"points":    score.Points,
		"player":    score.Player,
		"createdAt": score.CreatedAt,
	})
	if err != nil {
		log.Fatalf("Failed adding a new score: %v", err)
		return err
	}

	return nil
}

func (*repo) FindAll() ([]entity.Score, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var scores []entity.Score
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of scores: %v", err)
			return nil, err
		}
		score := entity.Score{
			Id:        doc.Data()["id"].(int64),
			Points:    doc.Data()["points"].(int64),
			Player:    doc.Data()["player"].(string),
			CreatedAt: doc.Data()["createdAt"].(int64),
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (*repo) FindByPlayer(player string) ([]entity.Score, error) {
	var scores []entity.Score
	return scores, nil
}
