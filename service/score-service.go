package service

import (
	"errors"
	"time"

	"github.com/luisjn/scores-api/entity"
	"github.com/luisjn/scores-api/repository"
)

var (
	repo repository.ScoreRepository
)

type ScoreService interface {
	Validate(score *entity.Score) error
	Create(score *entity.Score) error
	GetAll() ([]entity.Score, error)
	GetByPlayer(player string) ([]entity.Score, error)
}

type service struct{}

func NewScoreService(rep repository.ScoreRepository) ScoreService {
	repo = rep
	return &service{}
}

func (*service) Validate(score *entity.Score) error {
	if score == nil {
		err := errors.New("The score is empty")
		return err
	}
	if score.Player == "" {
		err := errors.New("The score player is empty")
		return err
	}
	return nil
}

func (*service) Create(score *entity.Score) error {
	score.CreatedAt = time.Now().UTC().Unix()
	return repo.Save(score)
}

func (*service) GetAll() ([]entity.Score, error) {
	return repo.FindAll()
}

func (*service) GetByPlayer(player string) ([]entity.Score, error) {
	return repo.FindByPlayer(player)
}
