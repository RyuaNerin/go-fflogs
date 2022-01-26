package fflogs

import (
	"context"
	"testing"
)

func TestRankingCharacter(t *testing.T) {
	c := newTestClient()

	opt := RankingCharacterOptions{
		CharacterName: testCharacterName,
		ServerName:    testServerName,
		ServerRegion:  testServerRegion,
		Zone:          &testZone,
		Encounter:     &testEncounter,
	}

	_, err := c.RankingCharacter(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
