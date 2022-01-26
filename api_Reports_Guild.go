package fflogs

import (
	"context"

	"github.com/RyuaNerin/go-fflogs/structure"
)

// Gets an object that contains a total count and an array of EncounterRanking objects and a total number of rankings for that encounter. Each EncounterRanking corresponds to a single character or guild/team.
func (c *Client) ReportsGuild(context context.Context, opt *ReportsGuildOptions) ([]structure.Report, error) {
	var resp []structure.Report
	err := c.Raw.ReportsGuild(context, opt, &resp)
	return resp, err
}
