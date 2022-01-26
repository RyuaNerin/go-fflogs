package fflogs

import (
	"context"
)

// Gets an array of Zone objects. Each zone corresponds to a raid/dungeon instance in the game and has its own set of encounters.
func (c *Raw) Zones(context context.Context, resp interface{}) error {
	return c.call(
		context,
		"/zones",
		nil,
		resp,
	)
}
