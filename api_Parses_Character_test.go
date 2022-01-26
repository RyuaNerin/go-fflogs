package fflogs

import (
	"context"
	"testing"
)

func TestParsesCharacter(t *testing.T) {
	c := newTestClient()

	opt := ParsesCharacterOptions{
		CharacterName: testCharacterName,
		ServerName:    testServerName,
		ServerRegion:  testServerRegion,
		Encounter:     &testEncounter,
		Zone:          &testZone,
	}

	_, err := c.ParsesCharacter(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
