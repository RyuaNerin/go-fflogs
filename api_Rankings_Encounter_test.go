package fflogs

import (
	"context"
	"testing"
)

func TestRankingEncounter(t *testing.T) {
	c := newTestClient()

	opt := RankingEncounterOptions{
		EncounterID: testEncounter,
	}

	_, err := c.RankingEncounter(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
