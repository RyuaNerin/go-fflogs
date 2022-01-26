package main

import (
	"context"
	"log"
	"os"

	"fflogs"

	"github.com/RyuaNerin/go-fflogs/structure"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	c, err := fflogs.NewClient(&fflogs.ClientOpt{
		ApiKey: os.Getenv("FFLOGS_V1_APIKEY"),
	})
	if err != nil {
		panic(err)
	}

	opt := fflogs.RankingEncounterOptions{
		EncounterID: 73,
		Metric:      &fflogs.MetricDps,
	}

	resp, err := c.RankingEncounter(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
	log.Println(resp.Rankings[0].GuildName)

	var respRaw structure.EncounterRankings
	err = c.Raw.RankingEncounter(context.Background(), &opt, &respRaw)
	if err != nil {
		panic(err)
	}
	log.Println(resp.Rankings[0].GuildName)
}
