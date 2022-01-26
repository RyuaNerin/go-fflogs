package fflogs

import (
	"context"

	"github.com/RyuaNerin/go-fflogs/structure"
)

// Gets an object that contains a totl count and an array of EncounterRanking objects and a total number of rankings for that encounter. Each EncounterRanking corresponds to a single character or guild/team.
func (c *Client) RankingEncounter(context context.Context, opt *RankingEncounterOptions) (*structure.EncounterRankings, error) {
	var resp structure.EncounterRankings
	err := c.Raw.RankingEncounter(context, opt, &resp)
	return &resp, err
}
