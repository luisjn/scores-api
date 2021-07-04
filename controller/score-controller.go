package controller

import (
	"encoding/json"
	"net/http"

	"github.com/luisjn/scores-api/entity"
	"github.com/luisjn/scores-api/errors"
	"github.com/luisjn/scores-api/service"
)

var (
	scoreService service.ScoreService
)

type ScoreController interface {
	GetScores(res http.ResponseWriter, req *http.Request)
	AddScore(res http.ResponseWriter, req *http.Request)
	GetScoresByPlayer(res http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewScoreController(serv service.ScoreService) ScoreController {
	scoreService = serv
	return &controller{}
}

func (*controller) GetScores(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	scores, err := scoreService.GetAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the scores"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(scores)
}

func (*controller) AddScore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	var score entity.Score
	err := json.NewDecoder(req.Body).Decode(&score)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error unmarshalling the data"})
		return
	}

	errv := scoreService.Validate(&score)
	if errv != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: errv.Error()})
		return
	}

	errs := scoreService.Create(&score)
	if errs != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error saving the score"})
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func (*controller) GetScoresByPlayer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	player := req.URL.Query().Get("player")
	scores, err := scoreService.GetByPlayer(player)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting the scores"})
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(scores)
}
