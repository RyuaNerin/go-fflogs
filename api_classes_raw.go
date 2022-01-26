package fflogs

import (
	"context"
)

// Gets an array of Class objects. Each Class corresponds to a class in thegame.
func (c *Raw) Classes(context context.Context, resp interface{}) error {
	return c.call(
		context,
		"/classes",
		nil,
		resp,
	)
}
