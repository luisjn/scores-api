package main

import (
	"os"

	"github.com/luisjn/scores-api/controller"
	router "github.com/luisjn/scores-api/http"
	"github.com/luisjn/scores-api/repository"
	"github.com/luisjn/scores-api/service"
)

var (
	httpRouter router.Router = router.NewChiRouter()
	// scoreRepository repository.ScoreRepository = repository.NewFirestoreRepository()
	scoreRepository repository.ScoreRepository = repository.NewPostgresRepository()
	scoreService    service.ScoreService       = service.NewScoreService(scoreRepository)
	scoreController controller.ScoreController = controller.NewScoreController(scoreService)
)

func main() {
	// httpRouter.GET("/scores", scoreController.GetScores)
	httpRouter.GET("/scores", scoreController.GetScoresByPlayer)
	httpRouter.POST("/scores", scoreController.AddScore)

	httpRouter.SERVE(os.Getenv("PORT"))
}
