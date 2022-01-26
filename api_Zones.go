package fflogs

import (
	"context"

	"fflogs/structure"
)

// Gets an array of Zone objects. Each zone corresponds to a raid/dungeon instance in the game and has its own set of encounters.
func (c *Client) Zones(context context.Context) ([]structure.Zone, error) {
	var resp []structure.Zone
	err := c.Raw.Zones(context, &resp)
	return resp, err
}
