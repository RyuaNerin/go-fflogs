package fflogs

import (
	"context"

	"github.com/RyuaNerin/go-fflogs/structure"
)

// Obtains all parses for a character in the zone across all specs. Every parse is included and not just rankings.
func (c *Client) ParsesCharacter(context context.Context, opt *ParsesCharacterOptions) ([]structure.CharacterRanking, error) {
	var resp []structure.CharacterRanking
	err := c.Raw.ParsesCharacter(context, opt, &resp)
	return resp, err
}
