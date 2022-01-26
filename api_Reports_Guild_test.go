package fflogs

import (
	"context"
	"testing"
)

func TestReportsGuild(t *testing.T) {
	c := newTestClient()

	opt := ReportsGuildOptions{
		GuildName:    testGuildName,
		ServerName:   testServerName,
		ServerRegion: testServerRegion,
	}

	_, err := c.ReportsGuild(context.Background(), &opt)
	if err != nil {
		panic(err)
	}
}
