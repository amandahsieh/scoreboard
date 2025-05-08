package scoreboard

import (
	"context"
	"log"
	"scoreboard/internal/scoreboard/db"
)

type ScoreboardService struct {
	Queries *db.Queries
}

func NewScoreboardService(queries *db.Queries) *ScoreboardService {
	return &ScoreboardService{Queries: queries}
}

func (s *ScoreboardService) ListAllScoreboards() ([]db.Scoreboard, error) {
	ctx := context.Background()
	scoreboards, err := s.Queries.GetAllScoreboards(ctx)
	if err != nil {
		log.Println("Error fetching scoreboards:", err)
		return nil, err
	}
	return scoreboards, nil
}

//func (s *ScoreboardService) CreateScoreboard(name string) (db.Scoreboard, error) {
//	ctx := context.Background()
//	scoreboard, err := s.Queries.CreateScoreboard(ctx, name)
//	if err != nil {
//		log.Println("Error creating scoreboard:", err)
//		return db.Scoreboard{}, err
//	}
//	return scoreboard, nil
//}
//
//func (s *ScoreboardService) GetScoreboardByID(id int32) (db.Scoreboard, error) {
//	ctx := context.Background()
//	scoreboard, err := s.Queries.GetScoreboardByID(ctx, id)
//	if err != nil {
//		log.Println("Error fetching scoreboard by ID:", err)
//		return db.Scoreboard{}, err
//	}
//	return scoreboard, nil
//}
//
//func (s *ScoreboardService) UpdateScoreboard(id int32, name string) (db.Scoreboard, error) {
//	ctx := context.Background()
//	params := queries.UpdateScoreboardParams{
//		ID:   id,
//		Name: name,
//	}
//	scoreboard, err := s.Queries.UpdateScoreboard(ctx, params)
//	if err != nil {
//		log.Println("Error updating scoreboard:", err)
//		return queries.Scoreboard{}, err
//	}
//	return scoreboard, nil
//}
//
//func (s *ScoreboardService) DeleteScoreboard(id int32) error {
//	ctx := context.Background()
//	err := s.Queries.DeleteScoreboard(ctx, id)
//	if err != nil {
//		log.Println("Error deleting scoreboard:", err)
//		return err
//	}
//	return nil
//}
