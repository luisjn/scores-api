package repository

import "github.com/luisjn/scores-api/entity"

type ScoreRepository interface {
	Save(score *entity.Score) error
	FindAll() ([]entity.Score, error)
	FindByPlayer(player string) ([]entity.Score, error)
}
