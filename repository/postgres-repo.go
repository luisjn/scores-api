package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/luisjn/scores-api/entity"
)

type postgresRepo struct{}

func NewPostgresRepository() ScoreRepository {
	return &postgresRepo{}
}

func (*postgresRepo) Save(score *entity.Score) error {
	ctx := context.Background()
	conn, err := connectDb(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return err
	}
	defer conn.Close(ctx)

	sql := "INSERT INTO scores (points, player, created_at) VALUES ($1, $2, $3)"
	if _, err := conn.Exec(ctx, sql, score.Points, score.Player, score.CreatedAt); err != nil {
		log.Fatalf("Failed adding a new score: %v", err)
		return err
	}
	return nil
}

func (*postgresRepo) FindByPlayer(player string) ([]entity.Score, error) {
	ctx := context.Background()
	conn, err := connectDb(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}
	defer conn.Close(ctx)

	var scores []entity.Score
	if rows, err := conn.Query(ctx, "SELECT * FROM scores WHERE player = $1 ORDER BY points DESC, created_at ASC", player); err != nil {
		log.Fatalf("Unable to select due to: %v", err)
		return nil, err
	} else {
		defer rows.Close()

		var tmp entity.Score

		for rows.Next() {
			rows.Scan(&tmp.Id, &tmp.Points, &tmp.Player, &tmp.CreatedAt)
			scores = append(scores, tmp)
		}
		if rows.Err() != nil {
			log.Fatalf("Error while reading scores table: %v", err)
			return nil, err
		}
	}
	return scores, nil
}

func (*postgresRepo) FindAll() ([]entity.Score, error) {
	ctx := context.Background()
	conn, err := connectDb(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}
	defer conn.Close(ctx)

	var scores []entity.Score
	if rows, err := conn.Query(ctx, "SELECT * FROM scores"); err != nil {
		log.Fatalf("Unable to select due to: %v", err)
		return nil, err
	} else {
		defer rows.Close()

		var tmp entity.Score

		for rows.Next() {
			rows.Scan(&tmp.Id, &tmp.Points, &tmp.Player, &tmp.CreatedAt)
			scores = append(scores, tmp)
		}
		if rows.Err() != nil {
			log.Fatalf("Error while reading scores table: %v", err)
			return nil, err
		}
	}

	return scores, nil
}

func connectDb(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("COCKROACHDB_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
