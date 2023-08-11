package main

import (
	"github.com/bonNope/makvesTask/internal/app"
	"github.com/bonNope/makvesTask/internal/handler"
	"github.com/bonNope/makvesTask/internal/models"
	"github.com/bonNope/makvesTask/internal/repository"
	"github.com/bonNope/makvesTask/internal/service"
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

func main() {
	entities, err := parseCSV()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(entities)
	s := service.NewService(repo)
	h := handler.NewHandler(s)

	srv := new(app.Server)
	if err := srv.Run(h.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func parseCSV() (map[int]*models.Entity, error) {
	file, err := os.Open("ueba.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entities := make([]*models.Entity, 10000)
	entitiesMap := map[int]*models.Entity{}

	err = gocsv.UnmarshalFile(file, entities)
	if err != nil {
		return nil, err
	}

	for _, v := range entities {
		if v != nil {
			entitiesMap[v.ID] = v
		}
	}

	return entitiesMap, nil
}
