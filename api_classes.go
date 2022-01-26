package fflogs

import (
	"context"

	"fflogs/structure"
)

// Gets an array of Class objects. Each Class corresponds to a class in thegame.
func (c *Client) Classes(context context.Context) ([]structure.Class, error) {
	var resp []structure.Class
	err := c.Raw.Classes(context, &resp)
	return resp, err
}
